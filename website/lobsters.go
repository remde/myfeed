package website

//LobsterTitle is the structure for the lobste.rs website
type LobsterTitle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

//GetLobstersTitles stores titles and URLs to target struct
func GetLobstersTitles(arrayOfTitles *[]LobsterTitle) {
	lobstersURL := "https://lobste.rs"
	SetTitlesToStruct(arrayOfTitles, lobstersURL)
}
