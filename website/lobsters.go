package website

import "fmt"

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

type rawLobsterStructArr []rawLobsterStruct

//InitLobsters intializes Lobster structure and prints the titles to the CLI
func InitLobsters(maxLinks int) []WebsiteStruct {
	lobstersURL := "https://lobste.rs"
	rawLobsterStruct := new(rawLobsterStructArr)
	setTitlesToStruct(lobstersURL, rawLobsterStruct)
	rawLobsterStruct.printTitlesToScreen(maxLinks)
	lobsterStruct := rawLobsterStruct.makeWebsiteStruct(maxLinks)
	return lobsterStruct
}

func (rawLobsterStruct rawLobsterStructArr) printTitlesToScreen(maxLinks int) {
	for i := 0; i <= maxLinks; i++ {
		fmt.Printf("%d: %s\n", i+1, rawLobsterStruct[i].Title)
	}
}

func (rawLobsterStruct rawLobsterStructArr) makeWebsiteStruct(maxLinks int) []WebsiteStruct {
	lobsterStruct := []WebsiteStruct{}
	for i := 0; i <= maxLinks; i++ {
		article := WebsiteStruct{}
		article.Index = i + 1
		article.URL = rawLobsterStruct[i].URL
		article.CommentsURL = rawLobsterStruct[i].CommentsURL
		lobsterStruct = append(lobsterStruct, article)
	}
	return lobsterStruct
}
