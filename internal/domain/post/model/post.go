package model

import "time"

type Post struct {
	ID          string
	Title       string
	Body        string
	UserID      string
	SubredditID string
	CreatedAt   time.Time
	VoteCount   int
}
