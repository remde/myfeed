package main

import (
	"github.com/remde/myfeed/cli"
)

func main() {
	initConfig := cli.GetInitConfig()
	cli.Start(initConfig)
}
