package lobsters

import "github.com/remde/myfeed/website"

type lobsterTitle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (lobster lobsterTitle) getTitles() {
	arrayOfTitles := new([]lobsterTitle)
	website.SetTitlesToStruct(arrayOfTitles, "https://lobste.rs")
}
