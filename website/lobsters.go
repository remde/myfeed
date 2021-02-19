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
func InitLobsters(maxLinks int) (WebsiteTable, string) {
	lobstersURL := "https://lobste.rs"
	rawLobsterStruct := new(rawLobsterStructArr)
	setTitlesToStruct(lobstersURL, rawLobsterStruct)
	articleTitles := rawLobsterStruct.printTitlesToScreen(maxLinks)
	lobsterStructTable := rawLobsterStruct.makeWebsiteStruct(maxLinks)
	return lobsterStructTable, articleTitles
}

func (rawLobsterStruct rawLobsterStructArr) printTitlesToScreen(maxLinks int) string {
	articleTitles := ""
	for i := 0; i <= maxLinks; i++ {
		articleTitles += fmt.Sprint(i+1) + ": " + rawLobsterStruct[i].Title + "\n"
	}
	return articleTitles
}

func (rawLobsterStruct rawLobsterStructArr) makeWebsiteStruct(maxLinks int) WebsiteTable {
	websiteTable := make(WebsiteTable, maxLinks)
	for i := 0; i <= maxLinks; i++ {
		lobsterStruct := websiteStructArr{}
		article := WebsiteStruct{rawLobsterStruct[i].URL, rawLobsterStruct[i].CommentsURL}
		lobsterStruct = append(lobsterStruct, article)
		websiteTable[i+1] = lobsterStruct
	}
	return websiteTable
}
