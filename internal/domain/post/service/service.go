package service

import (
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/model"
)

type PostService interface {
	CreatePost(userID string, req model.CreatePostRequest) (*model.Post, error)
	GetPostByID(postID string) (*model.Post, error)
	ListPosts(limit, offset int) ([]*model.Post, error)
}
