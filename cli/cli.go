package cli

import (
	"fmt"

	"github.com/remde/myfeed/parser"
	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config *parser.Config) {
	initWebsites(config)
	//reader := bufio.NewReader(os.Stdin)
}

func initWebsites(config *parser.Config) {
	switch configWebsite := config.Websites[0]; configWebsite {
	case "lobsters":
		website.InitLobsters(config.MaxLinks)
	case "reddit":
		website.InitReddit(config.MaxLinks)
	default:
		fmt.Println("This website is not yet supported: " + configWebsite)
	}
}
