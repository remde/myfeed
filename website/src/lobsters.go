package lobsters

import "github.com/remde/myfeed/website"

func getTitles() {
	arrayOfTitles := new([]website.LobsterTitles)
	website.SetTitlesToStruct(arrayOfTitles, "https://lobste.rs")
}
