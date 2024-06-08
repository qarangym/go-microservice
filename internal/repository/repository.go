package repository

import (
    "errors"
    "go-microservice/internal/models"
    "sync"
)

type Repository interface {
    GetItem(id string) (*models.Item, error)
    GetItems() ([]*models.Item, error)
    CreateItem(item models.Item) error
    UpdateItem(id string, item models.Item) error
    DeleteItem(id string) error
}

type inMemoryRepo struct {
    mu    sync.RWMutex
    items map[string]*models.Item
}

func NewInMemoryRepo() Repository {
    return &inMemoryRepo{
        items: make(map[string]*models.Item),
    }
}

func (r *inMemoryRepo) GetItem(id string) (*models.Item, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    item, ok := r.items[id]
    if !ok {
        return nil, errors.New("item not found")
    }
    return item, nil
}

func (r *inMemoryRepo) GetItems() ([]*models.Item, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    items := make([]*models.Item, 0, len(r.items))
    for _, item := range r.items {
        items = append(items, item)
    }
    return items, nil
}

func (r *inMemoryRepo) CreateItem(item models.Item) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.items[item.ID]; exists {
        return errors.New("item already exists")
    }
    r.items[item.ID] = &item
    return nil
}

func (r *inMemoryRepo) UpdateItem(id string, item models.Item) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.items[id]; !exists {
        return errors.New("item not found")
    }
    r.items[id] = &item
    return nil
}

func (r *inMemoryRepo) DeleteItem(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.items[id]; !exists {
        return errors.New("item not found")
    }
    delete(r.items, id)
    return nil
}
