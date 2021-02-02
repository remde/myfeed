package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config *config) {
	lobstersTable := initWebsites(config)
	fmt.Println("Enter number: ")
	var index int
	_, err := fmt.Scanf("%d", &index)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(lobstersTable[index])
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
