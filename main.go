package main

import (
	"log"
	"os"
	"time"

	"github.com/axiiomatic/reddit-downloader/cli"
)

func main() {

	time.Now().Nanosecond()

	start := time.Now()

	c1 := cli.Command{Name: "subr", Value: "string"}

	li := cli.Cli{Commands: []cli.Command{c1}}

	li.Init(os.Args)

	log.Println("took ", time.Since(start))

}
