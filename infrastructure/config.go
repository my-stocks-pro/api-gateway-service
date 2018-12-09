package infrastructure


type Config struct {
	Name string
	Port string
}

func NewConfig() Config {
	return Config{
		Name: "api-gateway",
		Port: ":9000",
	}
}
