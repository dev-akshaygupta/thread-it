package repository

import "github.com/dev-akshaygupta/thread-it/internal/domain/post/model"

type PostRepository interface {
	Save(post *model.Post) error
	FindByID(id string) (*model.Post, error)
	FindAll(limit, offset int) ([]*model.Post, error)
}
