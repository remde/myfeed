package cli

type config struct {
	Websites []string
	MaxLinks int
	Browser  string
}

//placeholder: currently hardcoded, should get and save to a "database"
//file (json?) as configuration depending on user input
func getInitConfig() *config {
	initConfig := new(config)
	initConfig.Websites = []string{"reddit", "lobsters"}
	initConfig.MaxLinks = 5
	initConfig.Browser = "chrome"
	return initConfig
}
