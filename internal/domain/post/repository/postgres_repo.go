package repository

import (
	"database/sql"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/model"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) PostRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Save(post *model.Post) error {
	query := `
	INSERT INTO posts (id, title, body, user_id, subreddit_id, created_at, vote_count)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query,
		post.ID, post.Title, post.Body, post.UserID,
		post.SubredditID, post.CreatedAt, post.VoteCount)
	return err
}

func (r *postgresRepo) FindByID(id string) (*model.Post, error) {
	query := `SELECT id, title, body, user_id, subreddit_id, created_at, vote_count from posts WHERE id=$1`
	row := r.db.QueryRow(query, id)

	var p model.Post
	if err := row.Scan(&p.ID, &p.Title, &p.Body, &p.UserID, &p.SubredditID, &p.CreatedAt, &p.VoteCount); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postgresRepo) FindAll(limit, offset int) ([]*model.Post, error) {
	query := `
	SELECT id, title, body, user_id, subreddit_id, created_at, vote_count
	FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []*model.Post
	for rows.Next() {
		var p model.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.UserID, &p.SubredditID, &p.CreatedAt, &p.VoteCount); err != nil {
			return nil, err
		}
		posts = append(posts, &p)
	}

	return posts, nil
}
