package config

type Config struct {
	RedPanda struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redpanda"`

	RedPandaTopic struct {
		Name string `json:"name"`
	} `json:"redpanda_topic"`
}

func Configure() Config {
	var config Config
	config.RedPanda.Host = "localhost"
	config.RedPanda.Port = 9092
	config.RedPandaTopic.Name = "redpanda"
	return config
}
