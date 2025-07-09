package delivery

import (
	"net/http"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *handler.PostHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/posts", handler.CreatePost)
	// r.Get("/posts/{id}", handler.Service.GetPostByID)
	// r.Get("/posts", handler.Service.ListPosts)

	return r
}
