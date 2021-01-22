package website

//RawLobsterStruct is the structure gotten from lobste.rs website
type RawLobsterStruct struct {
	Title         string `json:"title"`
	URL           string `json:"url"`
	Score         int    `json:"score"`
	CommentCount  int    `json:"comment_count"`
	CommentsURL   string `json:"comments_url"`
	SubmitterUser struct {
		Username string `json:"username"`
	} `json:"submitter_user"`
}

type lobsterStruct struct {
	Title        string
	URL          string
	Score        int
	CommentCount int
	CommentsURL  string
	SubmiterUser string
}

//GetLobstersTitles stores titles and URLs to target struct
func GetLobstersTitles(arrayOfTitles *[]RawLobsterStruct) {
	lobstersURL := "https://lobste.rs"
	SetTitlesToStruct(arrayOfTitles, lobstersURL)
}
