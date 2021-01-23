package parser

//Config is the structure for the user's default config
type Config struct {
	websites []string
	maxLinks int
	browser  string
}

//GetInitConfig parses the config.txt file in the root to obtain the user's default config
func GetInitConfig() *Config {
	initConfig := new(Config)
	initConfig.websites = []string{"reddit", "lobsters"}
	initConfig.maxLinks = 5
	initConfig.browser = "chrome"
	return initConfig
}
