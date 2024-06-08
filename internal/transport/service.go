package transport

import (
    "context"
    "log"
    "go-microservice/internal/models"
    "go-microservice/internal/repository"
    pb "go-microservice/proto"
)

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer struct {
    repo repository.Repository
    pb.UnimplementedItemServiceServer
}

// NewItemServiceServer creates a new gRPC server for ItemService.
func NewItemServiceServer(repo repository.Repository) *ItemServiceServer {
    return &ItemServiceServer{repo: repo}
}

// GetItem handles the GetItem gRPC request.
func (s *ItemServiceServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
    log.Printf("Received GetItem request: %v", req)
    item, err := s.repo.GetItem(req.GetId())
    if err != nil {
        return nil, err
    }
    return &pb.GetItemResponse{Item: &pb.Item{
        Id:    item.ID,
        Name:  item.Name,
        Price: int32(item.Price),
    }}, nil
}

// ListItems handles the ListItems gRPC request.
func (s *ItemServiceServer) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
    log.Printf("Received ListItems request")
    items, err := s.repo.GetItems()
    if err != nil {
        return nil, err
    }
    pbItems := make([]*pb.Item, len(items))
    for i, item := range items {
        pbItems[i] = &pb.Item{
            Id:    item.ID,
            Name:  item.Name,
            Price: int32(item.Price),
        }
    }
    return &pb.ListItemsResponse{Items: pbItems}, nil
}

// CreateItem handles the CreateItem gRPC request.
func (s *ItemServiceServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
    log.Printf("Received CreateItem request: %v", req)
    item := models.Item{
        ID:    req.GetItem().GetId(),
        Name:  req.GetItem().GetName(),
        Price: int(req.GetItem().GetPrice()),
    }
    if err := s.repo.CreateItem(item); err != nil {
        return nil, err
    }
    return &pb.CreateItemResponse{}, nil
}

// UpdateItem handles the UpdateItem gRPC request.
func (s *ItemServiceServer) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
    log.Printf("Received UpdateItem request: %v", req)
    item := models.Item{
        ID:    req.GetItem().GetId(),
        Name:  req.GetItem().GetName(),
        Price: int(req.GetItem().GetPrice()),
    }
    if err := s.repo.UpdateItem(req.GetId(), item); err != nil {
        return nil, err
    }
    return &pb.UpdateItemResponse{}, nil
}

// DeleteItem handles the DeleteItem gRPC request.
func (s *ItemServiceServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
    log.Printf("Received DeleteItem request: %v", req)
    if err := s.repo.DeleteItem(req.GetId()); err != nil {
        return nil, err
    }
    return &pb.DeleteItemResponse{}, nil
}
