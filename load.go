package loadConfig

type Config struct {
	config *map[string]string
}

func NewConfig() *Config {
	var config = &Config{
		config: new(map[string]string),
	}
	return config
}

func (c *Config) loadConfig(filename string)  {
	c.config = readConfig(filename)
}


func (c *Config) GetConfig(conStr string) string {
	cf := *c.config
	return cf[conStr]
}
