package goinworld

type Config struct {
	EngineHost string
	EngineSSL  bool
	StudioHost string
	StudioSSL  bool
}

func NewConfig() *Config {
	return &Config{
		EngineHost: "api-engine.inworld.ai:443",
		EngineSSL:  true,
		StudioHost: "api-studio.inworld.ai:443",
		StudioSSL:  true,
	}
}

//todo load config // viper or something
