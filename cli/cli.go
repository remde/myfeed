package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start() {
	initConfig := getInitConfig()
	lobstersTable, lobstersArticles := initWebsites(initConfig)
	for {
		fmt.Println(lobstersArticles)
		articleURL, articleCommentsURL := getArticleURL(lobstersTable)
		linkToOpen := chooseArticleOrComment(articleURL, articleCommentsURL)
		openArticle(linkToOpen)
		fmt.Println()
	}
}

//Placeholder: currently returning only lobstersTable
func initWebsites(config *config) (website.WebsiteTable, string) {
	lobstersTable := website.WebsiteTable{}
	lobstersArticles := ""
	for _, chosenWebsite := range config.Websites {
		if chosenWebsite == "lobsters" {
			lobstersTable, lobstersArticles := website.InitLobsters(config.MaxLinks)
			return lobstersTable, lobstersArticles
		} else if chosenWebsite == "reddit" {
			website.InitReddit(config.MaxLinks)
		} else {
			fmt.Println("This website is not yet supported: " + chosenWebsite)
		}
	}
	return lobstersTable, lobstersArticles
}
