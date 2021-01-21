package cli

import "fmt"

//Start initializes the CLI app
func Start(config interface{}) {
	printTitlesToScreen(config)
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println(config)
}

func printTitlesToScreen(config interface{}) {
	//TODO
}
