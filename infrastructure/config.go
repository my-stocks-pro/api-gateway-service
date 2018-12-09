package infrastructure


type Config struct {
	Name string
	Port string
}

func NewConfig() Config {
	return Config{
		Name: "api-handler",
		Port: ":9000",
	}
}
