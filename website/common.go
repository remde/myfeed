package website

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//WebsiteStruct is the generic structure for storing titles and URLs
type WebsiteStruct struct {
	Title        string
	URL          string
	Score        int
	CommentCount int
	CommentsURL  string
	SubmiterUser string
}

//SetTitlesToStruct makes the GET request and places Titles to the target structure
func setTitlesToStruct(websiteURL string, maxLinks string, targetStruct interface{}) {
	req, err := http.NewRequest("GET", websiteURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(bytes, &targetStruct)
}
