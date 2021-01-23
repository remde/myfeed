package parser

//Config is the structure for the user's default config
type Config struct {
	Websites []string
	MaxLinks int
	Browser  string
}

//GetInitConfig parses the config.txt file in the root to obtain the user's default config
func GetInitConfig() *Config {
	initConfig := new(Config)
	initConfig.Websites = []string{"reddit", "lobsters"}
	initConfig.MaxLinks = 5
	initConfig.Browser = "chrome"
	return initConfig
}
