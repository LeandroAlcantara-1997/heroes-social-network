package config

import "github.com/Netflix/go-env"

type environment struct {
	APIPort       string `env:"API_PORT"`
	DBName        string `env:"DB_NAME"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBHost        string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
	SplunkHost    string `env:"SPLUNK_HOST"`
	SplunkToken   string `env:"SPLUNK_TOKEN"`
	CacheHost     string `env:"CACHE_HOST"`
	CachePort     string `env:"CACHE_PORT"`
	CachePassword string `env:"CACHE_PASSWORD"`
	AllowOrigins  string `env:"ALLOW_ORIGINS"`
}

var Env environment

func LoadEnv() (err error) {
	_, err = env.UnmarshalFromEnviron(&Env)
	return
}
