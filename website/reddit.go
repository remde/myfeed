package website

//RawRedditStruct is the structure gotten from Reddit.com website
type RawRedditStruct struct {
	Data struct {
		Children []struct {
			Data struct {
				Title       string  `json:"title"`
				UpvoteRatio float64 `json:"upvote_ratio"`
				NumComments int     `json:"num_comments"`
				URL         string  `json:"url"`
			} `json:"data,omitempty"`
		}
	}
}

//GetRedditTitles stores titles and URLs to target struct
func GetRedditTitles(arrayOfTitles *[]RawRedditStruct) {
	redditURL := "https://reddit.com/r/cscareerquestions.json"
	SetTitlesToStruct(arrayOfTitles, redditURL)
}
