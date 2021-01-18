package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LobsterTitles struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func main() {
	var lobsterTitles []LobsterTitles

	req, err := http.NewRequest("GET", "https://lobste.rs", nil)
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

	json.Unmarshal(bytes, &lobsterTitles)

	fmt.Println(lobsterTitles)
}
