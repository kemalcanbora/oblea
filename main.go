package main

import (
	"oblea/config"
	m "oblea/pkg/databases/meilisearch"
)

func main() {
	index := m.NewMeiliSearch(config.Configure().MeiliSearch.Host,
		config.Configure().MeiliSearch.Port,
		config.Configure().MeiliSearchAPIKey,
		config.Configure().MeiliSearchIndex.Name,
	)

	data := map[string]interface{}{
		"id":   "55",
		"name": "test",
	}
	index.AddDocuments(data)
}
