package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

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

func getImageLink(url string) string {
	return strings.Replace(url, "amp;", "", 1)
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

	for _, v := range need.Data.Children {
		wg.Add(1)
		f := getImageLink(v.Data.Preview.Images[0].Source.URL)
		go downloadImage(f, subreddit)
	}

	wg.Wait()
}

func downloadImage(url string, sub string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot download image")
	}
	defer resp.Body.Close()

	if exists := CheckIfDirExists("memes"); !exists {
		done := CreateDir("memes")
		if !done {
			fmt.Println("in create memes")
			log.Fatal("Failed to create folder")
		}
	}

	path := fmt.Sprintf("%s%s", "memes/", sub)

	if exists := CheckIfDirExists(path); !exists {
		done := CreateDir(path)
		if !done {
			fmt.Println("in folderpath")
			log.Fatal("Failed to create folder")
		}
	}

	homeDir, _ := os.UserHomeDir()

	fileName := fmt.Sprintf("%s%s%s%s%s", homeDir, "/Pictures/memes/", sub, "/", time.Now())

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot download image")
	}

	_, e := io.Copy(file, resp.Body)

	if e != nil {
		log.Fatal("ok")
	}
	fmt.Println("image written to : ", fileName)
	wg.Done()

	return true, nil
}

func CheckIfDirExists(folder string) bool {

	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s%s%s", homeDir, "/Pictures/", folder)

	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func CreateDir(folder string) bool {

	homeDir, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s%s%s", homeDir, "/Pictures/", folder)

	if err := os.Mkdir(path, 0755); err != nil {
		return false
	}
	return true
}
