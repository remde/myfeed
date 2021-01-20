package website

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//SetTitlesToStruct makes the GET request and places Titles to the target structure
func SetTitlesToStruct(targetStruct interface{}, websiteURL string) {
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
