package greddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/google/go-querystring/query"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type transport struct {
	http.RoundTripper
	useragent string
}

// Any request headers can be modified here.
func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", t.useragent)
	return t.RoundTripper.RoundTrip(req)
}

// OAuthSession represents an OAuth session with reddit.com --
// all authenticated API calls are methods bound to this type.
type OAuthSession struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	OAuthConfig  *oauth2.Config
	//TokenExpiry  time.Time
	UserAgent string
	ctx       context.Context
	throttle  *rate.RateLimiter
}

// NewOAuthSession creates a new session for those who want to log into a
// reddit account via OAuth.
func NewOAuthSession(clientID, clientSecret, useragent, redirectURL string) (*OAuthSession, error) {
	o := &OAuthSession{}

	if len(useragent) > 0 {
		o.UserAgent = useragent
	} else {
		o.UserAgent = "TyANClient/0.1 by mustom"
	}

	// Set OAuth config
	o.OAuthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
		RedirectURL: redirectURL,
	}
	// Inject our custom HTTP client so that a user-defined UA can
	// be passed during any authentication requests.
	c := &http.Client{}
	c.Transport = &transport{http.DefaultTransport, o.UserAgent}
	o.ctx = context.WithValue(context.Background(), oauth2.HTTPClient, c)
	return o, nil
}

// Throttle sets the interval of each HTTP request.
// Disable by setting interval to 0. Disabled by default.
// Throttling is applied to invidual OAuthSession types.
func (o *OAuthSession) Throttle(interval time.Duration) {
	if interval == 0 {
		o.throttle = nil
		return
	}
	o.throttle = rate.New(1, interval)
}

// LoginAuth creates the required HTTP client with a new token.
func (o *OAuthSession) LoginAuth(username, password string) error {
	// Fetch OAuth token.
	t, err := o.OAuthConfig.PasswordCredentialsToken(o.ctx, username, password)
	if err != nil {
		return err
	}
	if !t.Valid() {
		msg := "Invalid OAuth token"
		if t != nil {
			if extra := t.Extra("error"); extra != nil {
				msg = fmt.Sprintf("%s: %s", msg, extra)
			}
		}
		return errors.New(msg)
	}
	o.Client = o.OAuthConfig.Client(o.ctx, t)
	return nil
}

// AuthCodeURL creates and returns an auth URL which contains an auth code.
func (o *OAuthSession) AuthCodeURL(state string, scopes []string) string {
	o.OAuthConfig.Scopes = scopes
	return o.OAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

// CodeAuth creates and sets a token using an authentication code returned from AuthCodeURL.
func (o *OAuthSession) CodeAuth(code string) error {
	t, err := o.OAuthConfig.Exchange(o.ctx, code)
	if err != nil {
		return err
	}
	o.Client = o.OAuthConfig.Client(o.ctx, t)
	return nil
}

// NeedsCaptcha check whether CAPTCHAs are needed for the Submit function.
func (o *OAuthSession) NeedsCaptcha() (bool, error) {
	var b bool
	err := o.getBody("https://oauth.reddit.com/api/needs_captcha", &b)
	if err != nil {
		return false, err
	}
	return b, nil
}

// Get request
func (o *OAuthSession) getBody(link string, d interface{}) error {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}

	if o.Client == nil {
		return errors.New("OAuth Session lacks HTTP client! Use func (o OAuthSession) LoginAuth() to make one.")
	}

	// Throttle request
	if o.throttle != nil {
		o.throttle.Wait()
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Check response status
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}

	return nil
}

// Post request
func (o *OAuthSession) postBody(link string, form url.Values, d interface{}) error {
	req, err := http.NewRequest("POST", link, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	// This is needed to avoid rate limits
	//req.Header.Set("User-Agent", o.UserAgent)

	// POST form provided
	req.PostForm = form

	if o.Client == nil {
		return errors.New("OAuth Session lacks HTTP client! Use func (o OAuthSession) LoginAuth() to make one.")
	}

	// Throttle request
	if o.throttle != nil {
		o.throttle.Wait()
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// The caller may want JSON decoded, or this could just be an update/delete request.
	if d != nil {
		err = json.Unmarshal(body, d)
		if err != nil {
			return err
		}
	}

	return nil
}

// SubredditSubmissions returns the submissions on the given subreddit using OAuth.
func (o *OAuthSession) SubredditSubmissions(subreddit string, sort PopularitySort, params ListingOptions) ([]*Submission, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	baseUrl := OauthBaseUrl

	// If subbreddit given, add to URL
	if subreddit != "" {
		baseUrl += "/r/" + subreddit
	}

	link := fmt.Sprintf(baseUrl+"/%s.json?%s", sort, v.Encode())

	r := new(Response)
	err = o.getBody(link, r)
	if err != nil {
		return nil, err
	}

	submissions := make([]*Submission, len(r.Data.Children))
	for i, child := range r.Data.Children {
		submissions[i] = child.Data
	}

	return submissions, nil
}

/*
after: fullname of a thing
before: fullname of a thing
category: a string no longer than 5 characters
count: a positive integer (default: 0)
include_facets: boolean value
limit: the maximum number of items desired (default: 25, maximum: 100)
q: a string no longer than 512 characters
restrict_sr: boolean value
show: (optional) the string all
sort: one of (relevance, hot, top, new, comments)
sr_detail: (optional) expand subreddits
t: one of (hour, day, week, month, year, all)
type: (optional) comma-delimited list of result types (sr, link, user)
*/
type SearchOptions struct {
	After         string `url:"after,omitempty"`
	Before        string `url:"before,omitempty"`
	Category      string `url:"category,omitempty"`
	Count         int    `url:"count,omitempty"`
	IncludeFacets bool   `url:"include_facets,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Query         string `url:"q,omitempty"`
	RestrictSr    bool   `url:"restrict_sr,omitempty"`
	Show          string `url:"show,omitempty"`
	Sort          string `url:"sort,omitempty"`
	SrDetail      bool   `url:"sr_detail,omitempty"`
	Time          string `url:"t,omitempty"`
	Type          string `url:"type,omitempty"`
}

func (o *OAuthSession) SearchSubmissions(subreddit string, params SearchOptions) ([]*Submission, error) {
	// Query parameters
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	// Base url
	// baseUrl := WwwBaseURL
	baseUrl := OauthBaseUrl
	// If subbreddit given, add to URL
	if subreddit != "" {
		baseUrl += "/r/" + subreddit
	}
	// Link
	link := fmt.Sprintf(baseUrl+"/search?%s", v.Encode())

	r := new(Response)
	err = o.getBody(link, r)
	if err != nil {
		return nil, err
	}

	submissions := make([]*Submission, len(r.Data.Children))
	for i, child := range r.Data.Children {
		submissions[i] = child.Data
	}

	return submissions, nil
}
