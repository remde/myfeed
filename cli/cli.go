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
	lobstersTable, lobstersArticles := initWebsites(config)
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
