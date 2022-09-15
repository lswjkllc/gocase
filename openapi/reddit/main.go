package main

import (
	"fmt"
	"gocase/openapi/reddit/greddit"
	"log"

	"github.com/jzelinskie/geddit"
)

func main() {
	// gredditMain()
	gedditMain()
}

func gredditMain() {
	// New oauth session
	o, err := greddit.NewOAuthSession(
		"7MLq2lH8xtpvziCQqO1dow",
		"AKhIKZ0KJ1qSkye9aMjXQqvbJnwlww",
		"TyANClient/0.1 by mustom",
		"",
	)
	if err != nil {
		log.Fatal(err)
	}
	// Login using our personal reddit account.
	err = o.LoginAuth("lswjkllc", "spreddit@me.com")
	if err != nil {
		log.Fatal(err)
	}

	// Build options
	opts := greddit.SearchOptions{
		After:      "",
		Query:      "apple watch",
		Limit:      5,
		RestrictSr: true,
		Sort:       greddit.HotSubmissions,
		Time:       "day"}
	// Fetch posts from r/videos, sorted by Hot.
	posts, err := o.SearchSubmissions("ios", opts)
	if err != nil {
		log.Fatal(err)
	}
	// Print the title and URL
	fmt.Println("Total:", len(posts))
	fullIds := make(map[string]struct{})
	for _, p := range posts {
		if _, ok := fullIds[p.FullID]; ok {
			log.Fatal("Duplicate id:", p.FullID)
			continue
		}
		fmt.Println(p.FullID, "--->", p.Title, "--->", p.URL)
	}
}

func gedditMain() {
	o, err := geddit.NewOAuthSession(
		"rmyI6J5Q9dBf2S0I9YzwsQ",
		"idC3nZbWrac2eE9ST5dFCo5e3IXpSw",
		"ChangeMeClient/0.1 by diniu85",
		"",
	)
	if err != nil {
		log.Fatal(err)
	}
	// Login using our personal reddit account.
	err = o.LoginAuth("diniu85", "niudi@123")
	if err != nil {
		log.Fatal(err)
	}

	submission := SearchSubmissions(o)
	SearchComments(o, submission)

}

func SearchComments(o *geddit.OAuthSession, submission *geddit.Submission) {
	fmt.Println("submission:", submission)
	// We can pass options to the query if desired (blank for now).
	opts := geddit.ListingOptions{
		Limit: 3,
	}

	posts, err := o.Comments(submission, geddit.HotSubmissions, opts)
	if err != nil {
		panic(err)
	}

	for i, p := range posts {
		fmt.Println(i, "===> ", p)
	}
}

func SearchSubmissions(o *geddit.OAuthSession) *geddit.Submission {
	// We can pass options to the query if desired (blank for now).
	opts := geddit.ListingOptions{
		Limit:   1,
		Article: "apple watch",
	}

	// Fetch posts from r/videos, sorted by Hot.
	posts, err := o.SubredditSubmissions("videos", geddit.HotSubmissions, opts)
	if err != nil {
		log.Fatal(err)
	}

	return posts[0]
}
