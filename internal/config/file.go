package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Server struct {
		HTTP struct {
			Address string `env:"SERVER_HTTP_ADDRESS" envDefault:":8000"`
		}
	}
	Logger struct {
	}
	Database struct {
		Driver   string `env:"DB_DRIVER"`
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
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

func (c *Config) GetDbDriver() string {
	return c.Database.Driver
}

func (c *Config) GetDbHost() string {
	return c.Database.Host
}

func (c *Config) GetDbPort() string {
	return c.Database.Port
}

func (c *Config) GetDbName() string {
	return c.Database.Name
}

func (c *Config) GetDbUser() string {
	return c.Database.User
}

func (c *Config) GetDbPassword() string {
	return c.Database.Password
}
