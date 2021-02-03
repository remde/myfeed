package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config *config) {
	lobstersTable := initWebsites(config)
	printChooseArticle()
	articleURL := getArticleURL(lobstersTable)
	openArticle(articleURL, config.Browser)
}

func initWebsites(config *config) website.WebsiteTable {
	lobstersTable := website.WebsiteTable{}
	for _, chosenWebsite := range config.Websites {
		if chosenWebsite == "lobsters" {
			lobstersTable = website.InitLobsters(config.MaxLinks)
		} else if chosenWebsite == "reddit" {
			website.InitReddit(config.MaxLinks)
		} else {
			fmt.Println("This website is not yet supported: " + chosenWebsite)
		}
	}
	return lobstersTable
}

func openArticle(url string, browser string) {
	//open article
}
