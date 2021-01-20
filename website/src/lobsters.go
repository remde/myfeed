package lobsters

import "github.com/remde/myfeed/website"

type LobsterTitles struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func getTitles() {
	var arrayOfTitles []LobsterTitles
	website.SetTitlesToStruct(arrayOfTitles, "https://lobste.rs")
}
