package config

type Config struct {
	DatabaseURL string
	ServerAddr  string
}

func Load() *Config {
	return &Config{
		DatabaseURL: "host=localhost port=5432 dbname=booking user=postgres password=123 sslmode=disable",
		ServerAddr:  ":8080",
	}
}
