package service

import (
	"time"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/model"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/repository"
	"github.com/google/uuid"
)

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

// Implement PostService
func (service *postService) CreatePost(userID string, req model.CreatePostRequest) (*model.Post, error) {
	post := &model.Post{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Body:        req.Body,
		UserID:      userID,
		SubredditID: req.SubredditID,
		CreatedAt:   time.Now().UTC(),
		VoteCount:   0,
	}

	dErr := service.repo.Save(post)
	if dErr != nil {
		return nil, dErr
	}

	return post, nil
}

func (service *postService) GetPostByID(postID string) (*model.Post, error) {
	return service.repo.FindByID(postID)
}

func (service *postService) ListPosts(limit, offset int) ([]*model.Post, error) {
	return service.repo.FindAll(limit, offset)
}
