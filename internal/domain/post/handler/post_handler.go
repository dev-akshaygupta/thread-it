package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/model"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/service"
)

type PostHandler struct {
	Service service.PostService
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req model.CreatePostRequest
	if dErr := json.NewDecoder(r.Body).Decode(&req); dErr != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(string)
	post, dErr := handler.Service.CreatePost(userID, req)
	if dErr != nil {
		http.Error(w, "could not create post", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(post)
}
