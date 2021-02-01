package website

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//WebsiteStruct is the generic structure for storing titles and URLs
type WebsiteStruct struct {
	URL         string
	CommentsURL string
}

type websiteStructArr []WebsiteStruct

//WebsiteTable is the structure that holds the website structure arrays indexed
type WebsiteTable map[int]websiteStructArr

type website interface {
	printTitlesToScreen()
	makeWebsiteStruct() WebsiteTable
}

func setTitlesToStruct(websiteURL string, targetStruct interface{}) {
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
