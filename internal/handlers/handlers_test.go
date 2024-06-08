package handlers

import (
	"bytes"
	"encoding/json"
	"go-microservice/internal/models"
	"go-microservice/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItem(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	handler := NewHandler(repo)

	item := models.Item{ID: "1", Name: "TestItem", Price: 100}
	repo.CreateItem(item)

	req, err := http.NewRequest("GET", "/item?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetItem(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotItem models.Item
	if err := json.NewDecoder(rr.Body).Decode(&gotItem); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if gotItem != item {
		t.Errorf("handler returned unexpected body: got %v want %v", gotItem, item)
	}
}

func TestCreateItem(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	handler := NewHandler(repo)

	item := models.Item{ID: "2", Name: "NewItem", Price: 200}
	itemBytes, _ := json.Marshal(item)

	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(itemBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.CreateItem(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	createdItem, err := repo.GetItem("2")
	if err != nil {
		t.Fatal(err)
	}

	if createdItem.Name != item.Name || createdItem.Price != item.Price {
		t.Errorf("handler did not create item correctly: got %v want %v", createdItem, item)
	}
}
