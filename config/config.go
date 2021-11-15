package config

type Config struct {
	RedPanda struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redpanda"`
	RedPandaTopic struct {
		Name string `json:"name"`
	} `json:"redpanda_topic"`
	MeiliSearch struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"meilisearch"`
	MeiliSearchIndex struct {
		Name string `json:"name"`
	} `json:"meilisearch_index"`
	MeiliSearchAPIKey string `json:"meilisearch_api_key"`
}

func Configure() Config {
	var config Config
	config.RedPanda.Host = "localhost"
	config.RedPanda.Port = 9092
	config.RedPandaTopic.Name = "redpanda"
	config.MeiliSearch.Host = "localhost"
	config.MeiliSearch.Port = 7700
	config.MeiliSearchIndex.Name = "redpanda"
	config.MeiliSearchAPIKey = "masterKey"
	return config
}
