package website

//RedditTitle is the structure for reddit.com website
type RedditTitle struct {
	Title string `json:"selftext"`
	URL   string `json:"url"`
}

//GetRedditTitles stores titles and URLs to target struct
func GetRedditTitles(arrayOfTitles *[]RedditTitle) {
	redditURL := "https://reddit.com/r/cscareerquestions.json"
	SetTitlesToStruct(arrayOfTitles, redditURL)
}
