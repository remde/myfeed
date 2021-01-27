package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config *config) {
	initWebsites(config)
	//reader := bufio.NewReader(os.Stdin)
}

func initWebsites(config *config) {
	for _, chosenWebsite := range config.Websites {
		if chosenWebsite == "lobsters" {
			website.InitLobsters(config.MaxLinks)
		} else if chosenWebsite == "reddit" {
			website.InitReddit(config.MaxLinks)
		} else {
			fmt.Println("This website is not yet supported: " + chosenWebsite)
		}
	}
}
