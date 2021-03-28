package config

import (
	"os"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	c.Server = os.Getenv("DB_SERVER") 
	c.Database = os.Getenv("DB_DATABASE")
}
