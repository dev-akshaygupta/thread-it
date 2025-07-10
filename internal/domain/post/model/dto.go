package model

type CreatePostRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	SubredditID string `json:"subreddit_id"`
}

type PostResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	UserID      string `json:"user_id"`
	SubredditID string `json:"subreddit_id"`
	CreatedAt   string `json:"created_at"`
	VoteCount   int    `json:"vote_count"`
}
