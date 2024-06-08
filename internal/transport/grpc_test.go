package transport

import (
	"context"
	"go-microservice/internal/repository"
	pb "go-microservice/proto"
	"google.golang.org/grpc"
	"testing"
)

func TestGRPCServer(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	go StartGRPCServer(repo)

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewItemServiceClient(conn)

	// Create an item
	item := &pb.Item{Id: "1", Name: "TestItem", Price: 100}
	_, err = client.CreateItem(context.Background(), &pb.CreateItemRequest{Item: item})
	if err != nil {
		t.Fatalf("could not create item: %v", err)
	}

	// Get the item
	resp, err := client.GetItem(context.Background(), &pb.GetItemRequest{Id: "1"})
	if err != nil {
		t.Fatalf("could not get item: %v", err)
	}

	if resp.Item.Name != "TestItem" || resp.Item.Price != 100 {
		t.Errorf("item returned is not correct: got %v", resp.Item)
	}
}
