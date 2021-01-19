package lobsters

import (
	"github.com/remde/myfeed/website/common"
)

type LobsterTitles struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func getTitles() {
	arrayOfTitles = new(LobsterTitles)
	common.GetTitles(arrayOfTitles)
}
