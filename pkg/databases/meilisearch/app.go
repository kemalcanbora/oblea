package meilisearch

import (
	"fmt"
	meili "github.com/meilisearch/meilisearch-go"
	"oblea/config"
)

func Client() *meili.Index {
	address := fmt.Sprintf("http://%s:%d", config.Configure().MeiliSearch.Host, config.Configure().MeiliSearch.Port)
	client := meili.NewClient(meili.ClientConfig{
		Host:   address,
		APIKey: config.Configure().MeiliSearchAPIKey,
	})

	indexName := config.Configure().MeiliSearchIndex.Name
	var index *meili.Index

	index, _ = client.GetIndex(indexName)
	if index == nil {
		_, createErr := client.CreateIndex(&meili.IndexConfig{
			Uid:        indexName,
			PrimaryKey: "id",
		})
		if createErr != nil {
			fmt.Println(createErr)
		}
		index, _ = client.GetIndex(indexName)
	}
	return index
}

func AddDocument(index *meili.Index, data map[string]interface{}) bool {
	documents := []map[string]interface{}{data}
	_, err := index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
