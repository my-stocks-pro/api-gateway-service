package infrastructure

import "time"

type Config struct {
	StartDate string
	Name      string
	Port      string
}

func NewConfig() Config {
	return Config{
		StartDate: time.Now().Format("2006-01-02 15:04"),
		Name:      "api-gateway",
		Port:      ":9000",
	}
}
