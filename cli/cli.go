package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

//Start initializes the CLI app
func Start(config interface{}) {
	printTitlesToScreen(config)
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println(config)
}

func printTitlesToScreen(config interface{}) {
	lobster := new([]website.LobsterTitle)
	website.GetLobstersTitles(lobster)
	reddit := new([]website.RedditTitle)
	website.GetRedditTitles(reddit)
	fmt.Println(reddit)
}
