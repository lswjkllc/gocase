package greddit

// PopularitySort represents the possible ways to sort submissions by popularity.
type PopularitySort string

const (
	DefaultPopularity        PopularitySort = ""
	HotSubmissions                          = "hot"
	NewSubmissions                          = "new"
	RisingSubmissions                       = "rising"
	TopSubmissions                          = "top"
	ControversialSubmissions                = "controversial"
)

type ListingOptions struct {
	Time    string `url:"t,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	After   string `url:"after,omitempty"`
	Before  string `url:"before,omitempty"`
	Count   int    `url:"count,omitempty"`
	Show    string `url:"show,omitempty"`
	Article string `url:"article,omitempty"`
}

// response
type Response struct {
	Kind string       `json:"kind,omitempty"`
	Data ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	After    string          `json:"after,omitempty"`
	Before   string          `json:"before,omitempty"`
	Children []*DataChildren `json:"children,omitempty"`
}

type DataChildren struct {
	Kind string      `json:"kind,omitempty"`
	Data *Submission `json:"data,omitempty"`
}

const (
	T2 string = "t2" // author => 提交者
	T3 string = "t3" // submission => 帖子
	T5 string = "t5" // subreddit => 贴吧
)
