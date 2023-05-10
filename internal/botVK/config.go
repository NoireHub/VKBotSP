package botvk

type Config struct {
	DatabaseURL string `toml:"database_url"`
	Token string `toml:"token"`
}


func NewConfig() *Config {
	return &Config{}
}