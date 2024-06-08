package transport

import (
    "go-microservice/internal/repository"
    pb "go-microservice/proto"
    "google.golang.org/grpc"
    "log"
    "net"
)

func StartGRPCServer(repo repository.Repository) {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterItemServiceServer(grpcServer, NewItemServiceServer(repo))

    log.Println("Starting gRPC server on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
