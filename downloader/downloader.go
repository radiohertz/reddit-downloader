package downloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ENDPOINT = "https://www.reddit.com/r/"
	LIMIT    = 25
)

type Image struct {
	Source struct {
		URL string `json:"url"`
	} `json:"source"`
}

type Post struct {
	Data struct {
		Title   string `json:"title"`
		Preview struct {
			Images []Image `json:"images"`
		} `json:"preview"`
	} `json:"data"`
}

type Data struct {
	Kind string `json:"kind"`
	Data struct {
		Dist     int    `json:"dist"`
		Children []Post `json:"children"`
	} `json:"data"`
}

func buildEndpoint(subr string, limit int) string {
	return fmt.Sprintf("%s%s%s%s%d", ENDPOINT, subr, ".json?", "limit=", limit)
}

func MakeRequestForReddit(subreddit string, limit int) {
	uriString := buildEndpoint(subreddit, limit)

	client := &http.Client{}

	req, err := http.NewRequest("GET", uriString, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")

	resp, err := client.Do(req)

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

	var need Data

	err = json.Unmarshal(data, &need)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", need.Data.Children[0].Data.Preview.Images)

}
