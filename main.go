package main

import (
	"os"

	"github.com/axiiomatic/reddit-downloader/cli"
)

func main() {

	c1 := cli.Command{Name: "subr", Value: "string"}

	li := cli.Cli{Commands: []cli.Command{c1}}

	li.Init(os.Args)

}
