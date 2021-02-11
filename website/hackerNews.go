package website

import "fmt"

type rawHackerNewsStruct struct {
}

type rawHackerNewsStructArr []rawHackerNewsStruct

//InitHackerNews intializes Hacker News structure and prints the titles to the CLI
func InitHackerNews(maxLinks int) WebsiteTable {
	hackerNewsURL := "https://hackernews.com"
	rawHackerNewsStruct := new(rawHackerNewsStruct)
	setTitlesToStruct(hackerNewsURL, rawHackerNewsStruct)
	rawHackerNewsStruct.printTitlesToScreen(maxLinks)
	hackerNewsStructTable := rawHackerNewsStruct.makeWebsiteStruct(maxLinks)
	return hackerNewsStructTable
}

func (rawHackerNewsStruct rawHackerNewsStructArr) printTitlesToScreen(maxLinks int) {
	for i := 0; i <= maxLinks; i++ {
		fmt.Printf("%d: %s\n", i+1, rawHackerNewsStruct[i].Title)
	}
}

func (rawHackerNewsStruct rawHackerNewsStructArr) makeWebsiteStruct(maxLinks int) WebsiteTable {
	websiteTable := make(WebsiteTable, maxLinks)
	hackerNewsStruct := websiteStructArr{}
	for i := 0; i <= maxLinks; i++ {
		article := WebsiteStruct{}
		article.URL = rawHackerNewsStruct[i].URL
		article.CommentsURL = rawHackerNewsStruct[i].CommentsURL
		hackerNewsStruct = append(hackerNewsStruct, article)
		websiteTable[i+1] = hackerNewsStruct
	}
	return websiteTable
}
