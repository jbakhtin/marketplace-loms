package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Server struct {
		HTTP struct {
			Address string `env:"SERVER_HTTP_ADDRESS" envDefault:":8080"`
		}
	}
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) GetServerHTTPAddress() string {
	return c.Server.HTTP.Address
}
