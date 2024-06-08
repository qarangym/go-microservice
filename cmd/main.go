
package main

import (
    "go-microservice/internal/handlers"
    "go-microservice/internal/repository"
    "go-microservice/internal/transport"
    "log"
    "net/http"
)

func main() {
    log.Println("Starting HTTP server on port 8080")
    repo := repository.NewInMemoryRepo()
    handler := handlers.NewHandler(repo)

    go func() {
        log.Println("Starting gRPC server on port 50051")
        transport.StartGRPCServer(repo)
    }()

    server := transport.NewHTTPServer(handler)
    log.Fatal(http.ListenAndServe(":8080", server.Handler))
}
