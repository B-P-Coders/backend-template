package config

var Conf *Config

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     uint16 `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

type ServerConfig struct {
	Port         uint16 `env:"PORT"`
	AllowOrigins string `env:"ALLOW_ORIGINS"`
}

type Config struct {
	DbConfig  DatabaseConfig `envPrefix:"DB_"`
	SrvConfig ServerConfig   `envPrefix:"SERVER_"`
}
