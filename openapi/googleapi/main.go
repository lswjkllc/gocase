package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "portable power station", "Search term")
	maxResults = flag.Int64("max-results", 50, "Max YouTube results")
)

const developerKey = "AIzaSyB4HbZGMwP0lZSdgSp98c-_X17nVKRujPc"

// const developerKey = "AIzaSyBubAhjHpeEWbQgvvEgqeFOUHt9Gg7Zu38"

func getUSPacificDateFormat(t time.Time, deltaHour int64) string {
	timeLocation := time.FixedZone("US/Pacific", -7*60*60)
	uspTime := t.In(timeLocation).Add(time.Hour * time.Duration(deltaHour))
	uspTimeStr := uspTime.Format("2006-01-02")
	return uspTimeStr
}

func main() {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Call the search.list method to retrieve results matching the specified
	after := getUSPacificDateFormat(time.Now(), -48) + "T00:00:00Z" // "2006-01-02T15:04:05.000Z"
	before := getUSPacificDateFormat(time.Now(), -48) + "T23:59:59Z"
	fmt.Println("after:", after, "before:", before)
	items, err := SearchVideos(service, *query, after, before, *maxResults)
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Println("Total:", items.TotalResults, "Max:", items.MaxResults)
	for i, item := range items.Items {
		printItem(i, item)

		// Call the video.list method to retrieve the specified video
		video, err := GetVideoById(service, item.Pid)
		if err != nil {
			log.Fatalf("Error making YouTube API call: %v", err)
		}
		printVideo(video)
		fmt.Println()
	}
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printItem(i int, item SItem) {
	fmt.Println(i, item.Pid)
	// jstr, _ := json.Marshal(item)
	// fmt.Println(i, string(jstr))
}

func SearchVideos(service *youtube.Service, keyword string, after, before string, maxResults int64) (*SearchItems, error) {
	parts := []string{"snippet"}
	fields := googleapi.Field("nextPageToken, pageInfo, items(id(kind, videoId), snippet(channelTitle, channelId, title))")
	call := service.Search.List(parts).Fields(fields).Q(keyword).MaxResults(maxResults).PublishedAfter(after).PublishedBefore(before)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var respItems []*youtube.SearchResult = response.Items
	nextPageToken := response.NextPageToken
	for nextPageToken != "" {
		response, err = call.PageToken(nextPageToken).Do()
		if err != nil {
			fmt.Println("NextPageToken:", nextPageToken, "Error:", err)
			continue
		}
		respItems = append(respItems, response.Items...)
		nextPageToken = response.NextPageToken
	}
	searchItems := &SearchItems{Keyword: keyword, MaxResults: maxResults}

	items := make([]SItem, 0)
	for _, item := range respItems {
		pubAt := item.Snippet.PublishedAt
		fmt.Println(pubAt)
		sm := SItem{
			Pid: item.Id.VideoId, Kind: item.Id.Kind,
			PublishedAt: item.Snippet.PublishedAt, ChannelId: item.Snippet.ChannelId,
			Title: item.Snippet.Title, Description: item.Snippet.Description}
		items = append(items, sm)
	}
	searchItems.Items = items
	searchItems.TotalResults = int64(len(items))

	return searchItems, nil
}

func printVideo(video *Video) {
	jstr, _ := json.Marshal(video)
	fmt.Println(string(jstr))
}

func GetVideoById(service *youtube.Service, id string) (*Video, error) {
	parts := []string{"snippet", "contentDetails", "statistics", "topicDetails", "recordingDetails"}
	fields := googleapi.Field("items(id, snippet(publishedAt, channelId, title, description, categoryId, defaultLanguage), contentDetails/duration, statistics, topicDetails/topicCategories, recordingDetails/location)")
	call := service.Videos.List(parts).Fields(fields).Id(id)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	item := response.Items[0]
	video := &Video{
		Pid: item.Id, Title: item.Snippet.Title, Description: item.Snippet.Description,
		PublishedAt: item.Snippet.PublishedAt, ChannelId: item.Snippet.ChannelId,
		TopicDetails:   TopicDetail{TopicCategories: item.TopicDetails.TopicCategories},
		ContentDetails: ContentDetail{Duration: item.ContentDetails.Duration},
		Statistics: Statistic{
			ViewCount: item.Statistics.ViewCount, LikeCount: item.Statistics.LikeCount,
			DislikeCount: item.Statistics.DislikeCount, FavoriteCount: item.Statistics.FavoriteCount,
			CommentCount: item.Statistics.CommentCount},
		RecordingDetails: RecordingDetail{Location: item.RecordingDetails.Location},
	}

	return video, nil
}

type SearchItems struct {
	Keyword      string `json:"keyword"`
	TotalResults int64  `json:"totalResults"`
	MaxResults   int64  `json:"maxResults"`

	Items []SItem `json:"items"`
}

type SItem struct {
	Pid         string `json:"pid"`
	Kind        string `json:"kind"`
	PublishedAt string `json:"publishedAt"`
	ChannelId   string `json:"channelId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Video struct {
	/// 原始数据
	Pid              string          `json:"pid"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	PublishedAt      string          `json:"publishedAt"`
	ChannelId        string          `json:"channelId"`
	TopicDetails     TopicDetail     `json:"topicDetails"`
	ContentDetails   ContentDetail   `json:"contentDetails"`
	Statistics       Statistic       `json:"statistics"`
	RecordingDetails RecordingDetail `json:"recordingDetails"`

	Others Other `json:"other"`
}

type TopicDetail struct {
	// TopicIds        []string
	TopicCategories []string `json:"topicCategories"`
}

type ContentDetail struct {
	Duration string `json:"duration"`
}

type Statistic struct {
	ViewCount     uint64 `json:"viewCount"`
	LikeCount     uint64 `json:"likeCount"`
	DislikeCount  uint64 `json:"dislikeCount"`
	FavoriteCount uint64 `json:"favoriteCount"`
	CommentCount  uint64 `json:"commentCount"`
}

type RecordingDetail struct {
	Location *GeoPoint `json:"location"`
}

type GeoPoint = youtube.GeoPoint

// type Location struct {
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

type Other struct {
	Keyword      string `json:"keyword"`
	IsCleaned    bool   `json:"isCleaned"`
	DataId       int    `json:"dataId"`
	CapturedTime string `json:"capturedTime"`
	Mention      string `json:"mention"`
}
