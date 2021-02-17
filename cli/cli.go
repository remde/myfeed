package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config *config) {
	lobstersTable := initWebsites(config)
	for {
		articleURL, articleCommentsURL := getArticleURL(lobstersTable)
		linkToOpen := chooseArticleOrComment(articleURL, articleCommentsURL)
		openArticle(linkToOpen)
	}
}

//Placeholder: currently returning only lobstersTable
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

func openArticle(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
