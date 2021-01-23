package main

import (
	"github.com/remde/myfeed/cli"
	"github.com/remde/myfeed/parser"
)

func main() {
	initConfig := parser.GetInitConfig()
	cli.Start(initConfig)
}
