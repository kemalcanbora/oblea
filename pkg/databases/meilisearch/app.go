package meilisearch

import (
	"fmt"
	meili "github.com/meilisearch/meilisearch-go"
)

type MeiliSearch struct {
	client *meili.Client
	index  *meili.Index
}

func NewMeiliSearch(host string, port int, apiKey string, indexName string) *MeiliSearch {
	address := fmt.Sprintf("http://%s:%d", host, port)
	client := meili.NewClient(meili.ClientConfig{
		Host:   address,
		APIKey: apiKey,
	})

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
	return &MeiliSearch{
		client: client,
		index:  index,
	}
}

func (m *MeiliSearch) AddDocuments(data map[string]interface{}) bool {
	documents := []map[string]interface{}{data}
	_, err := m.index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
