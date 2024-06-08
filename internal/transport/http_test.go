package transport

import (
	"go-microservice/internal/handlers"
	"go-microservice/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPServer(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	handler := handlers.NewHandler(repo)
	server := NewHTTPServer(handler)

	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server.Handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
