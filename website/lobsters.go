package website

import "fmt"

//RawLobsterStruct is the structure gotten from lobste.rs website
type rawLobsterStruct struct {
	Title         string `json:"title"`
	URL           string `json:"url"`
	Score         int    `json:"score"`
	CommentCount  int    `json:"comment_count"`
	CommentsURL   string `json:"comments_url"`
	SubmitterUser struct {
		Username string `json:"username"`
	} `json:"submitter_user"`
}

//InitLobsters intializes Lobster structure and prints the titles to the CLI
func InitLobsters(maxLinks int) {
	lobstersURL := "https://lobste.rs"
	maxLinksFilter := "$filter=" + fmt.Sprint(maxLinks)
	rawLobsterStruct := new([]rawLobsterStruct)
	setTitlesToStruct(lobstersURL, maxLinksFilter, rawLobsterStruct)
}
