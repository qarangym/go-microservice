package handlers

import (
    "encoding/json"
    "go-microservice/internal/models"
    "go-microservice/internal/repository"
    "log"
    "net/http"
)

type Handler struct {
    repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
    return &Handler{repo: repo}
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
    log.Println("GetItem handler called")
    id := r.URL.Query().Get("id")
    item, err := h.repo.GetItem(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(item)
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
    log.Println("GetItems handler called")
    items, err := h.repo.GetItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
    log.Println("CreateItem handler called")
    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.repo.CreateItem(item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
    log.Println("UpdateItem handler called")
    id := r.URL.Query().Get("id")
    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.repo.UpdateItem(id, item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
    log.Println("DeleteItem handler called")
    id := r.URL.Query().Get("id")
    if err := h.repo.DeleteItem(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
