package main

import "github.com/remde/myfeed/cli"

type config struct {
	websites []string
	maxLinks int
	browser  string
}

//this is placeholder. getInitConfig will parse from a txt in the root
func getInitConfig() *config {
	initConfig := new(config)
	initConfig.websites = []string{"lobsters", "reddit", "hackernews"}
	initConfig.maxLinks = 5
	initConfig.browser = "google-chrome"
	return initConfig
}

func main() {
	initConfig := getInitConfig()
	cli.Start(initConfig)
}
