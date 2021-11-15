package main

import (
	"oblea/models"
	m "oblea/pkg/databases/meilisearch"
	"testing"
)

//create a connection test
func TestConnection(t *testing.T) {
	index := m.Client()
	if index == nil {
		t.Error("connection is nil")
	}
}

func TestAddDocument(t *testing.T) {
	index := m.Client()
	if index == nil {
		t.Error("connection is nil")
	}

	d := models.Document{
		ID:   "55",
		Name: "test",
	}

	data := map[string]interface{}{
		"id":   d.ID,
		"name": d.Name,
	}

	result := m.AddDocument(index, data)
	if result != true {
		t.Error("add document is false")
	}
}
