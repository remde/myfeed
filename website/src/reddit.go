package source

import "github.com/remde/myfeed/website"

type redditTitle struct {
	Title string `json:"selftext"`
	URL   string `json:"url"`
}

func (reddit redditTitle) getTitles(subreddit string) {
	arrayOfTitles := new([]redditTitle)
	website.SetTitlesToStruct(arrayOfTitles, getRedditURL(subreddit))
}

func getRedditURL(subreddit string) string {
	return "https://reddit.com/r/" + subreddit + "/.json"
}
