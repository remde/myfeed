package cli

type config struct {
	Websites []string
	MaxLinks int
	Browser  string
}

//Placeholder: currently hardcoded
//GetInitConfig parses the config.txt file in the root to obtain the user's default config
func GetInitConfig() *config {
	initConfig := new(config)
	initConfig.Websites = []string{"reddit", "lobsters"}
	initConfig.MaxLinks = 5
	initConfig.Browser = "chrome"
	return initConfig
}
