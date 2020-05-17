package downloader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ENDPOINT = "https://www.reddit.com/r/"
	LIMIT    = 25
)

func buildEndpoint(subr string, limit int) string {
	return fmt.Sprintf("%s%s%s%s%d", ENDPOINT, subr, ".json?", "limit=", limit)
}

func MakeRequestForReddit(subreddit string, limit int) {
	uriString := buildEndpoint(subreddit, limit)

	fmt.Println(uriString)
	resp, err := http.Get(uriString)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Something went wrong")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Something went wrong")
	}

	fmt.Printf(string(data))

}
