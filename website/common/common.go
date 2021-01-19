package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetTitles(targetStruct, websiteUrl) {
	req, err := http.NewRequest("GET", websiteUrl, nil)
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
