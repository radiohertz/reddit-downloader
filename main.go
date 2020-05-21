package main

import (
	"log"
	"os"
	"time"

	"github.com/axiiomatic/reddit-downloader/cli"
)

func main() {

	start := time.Now()

	li := cli.Cli{Commands: []cli.Command{}}

	li.Init(os.Args)

	log.Println("took ", time.Since(start))

}
