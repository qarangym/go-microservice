package transport

import (
    "go-microservice/internal/handlers"
    "net/http"

    "github.com/gorilla/mux"
)

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    })
}

func NewHTTPServer(handler *handlers.Handler) *http.Server {
    r := mux.NewRouter()
    r.HandleFunc("/items", handler.GetItems).Methods("GET")
    r.HandleFunc("/item", handler.GetItem).Methods("GET")
    r.HandleFunc("/item", handler.CreateItem).Methods("POST")
    r.HandleFunc("/item", handler.UpdateItem).Methods("PUT")
    r.HandleFunc("/item", handler.DeleteItem).Methods("DELETE")

    corsRouter := enableCORS(r)

    return &http.Server{
        Addr:    ":8080",
        Handler: corsRouter,
    }
}
