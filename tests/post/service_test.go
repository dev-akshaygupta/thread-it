package post

import (
	"testing"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/model"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (mock *mockRepo) Save(post *model.Post) error {
	args := mock.Called(post)
	return args.Error(0)
}

func (mock *mockRepo) FindByID(id string) (*model.Post, error) {
	args := mock.Called(id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (mock *mockRepo) FindAll(limit, offset int) ([]*model.Post, error) {
	args := mock.Called(limit, offset)
	return args.Get(0).([]*model.Post), args.Error(1)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := service.NewPostService(mockRepo)

	input := model.CreatePostRequest{
		Title:       "First blog",
		Body:        "This is first blog",
		SubredditID: "Only blog",
	}
	userID := "u12322"

	mockRepo.On("Save", mock.AnythingOfType("*model.Post")).Return(nil)

	result, dErr := svc.CreatePost(userID, input)
	assert.NoError(t, dErr)
	assert.Equal(t, input.Title, result.Title)
	mockRepo.AssertCalled(t, "Save", mock.AnythingOfType("*model.Post"))
}
