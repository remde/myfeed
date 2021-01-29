package website

type rawRedditStruct struct {
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

//InitReddit intializes Reddit structure and prints the titles to the CLI
func InitReddit(maxLinks int) {
	redditURL := "https://reddit.com/r/cscareerquestions.json"
	maxLinksFilter := "" //placeholder
	redditURL = redditURL + maxLinksFilter
	rawRedditStruct := new([]rawRedditStruct)
	setTitlesToStruct(redditURL, rawRedditStruct)
}
