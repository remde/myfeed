package website

//LobsterTitles is how lobste.rs titles and URLs are stored
type LobsterTitles struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

//RedditTitles is how reddit.com titles and URLS are stored
type RedditTitles struct {
	Title string
	URL   string
}
