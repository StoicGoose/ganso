// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: comment.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    user_id,
    post_id,
    content
) VALUES ($1, $2, $3) RETURNING id, user_id, post_id, edited, date_time, content
`

type CreateCommentParams struct {
	UserID  string `json:"user_id"`
	PostID  string `json:"post_id"`
	Content string `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.UserID, arg.PostID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PostID,
		&i.Edited,
		&i.DateTime,
		&i.Content,
	)
	return i, err
}

const createReply = `-- name: CreateReply :one
INSERT INTO replies (
    user_id,
    comment_id,
    content
) VALUES ($1, $2, $3) RETURNING id, user_id, comment_id, edited, date_time, content
`

type CreateReplyParams struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
	Content   string `json:"content"`
}

func (q *Queries) CreateReply(ctx context.Context, arg CreateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, createReply, arg.UserID, arg.CommentID, arg.Content)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CommentID,
		&i.Edited,
		&i.DateTime,
		&i.Content,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const deleteCommentsByPost = `-- name: DeleteCommentsByPost :exec
DELETE FROM comments WHERE post_id = $1
`

func (q *Queries) DeleteCommentsByPost(ctx context.Context, postID string) error {
	_, err := q.db.ExecContext(ctx, deleteCommentsByPost, postID)
	return err
}

const deleteCommentsByUser = `-- name: DeleteCommentsByUser :exec
DELETE FROM comments WHERE user_id = $1
`

func (q *Queries) DeleteCommentsByUser(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteCommentsByUser, userID)
	return err
}

const deleteRepliesByPost = `-- name: DeleteRepliesByPost :exec
DELETE FROM replies WHERE comment_id = $1
`

func (q *Queries) DeleteRepliesByPost(ctx context.Context, commentID string) error {
	_, err := q.db.ExecContext(ctx, deleteRepliesByPost, commentID)
	return err
}

const deleteRepliesByUser = `-- name: DeleteRepliesByUser :exec
DELETE FROM replies WHERE user_id = $1
`

func (q *Queries) DeleteRepliesByUser(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteRepliesByUser, userID)
	return err
}

const deleteReply = `-- name: DeleteReply :exec
DELETE FROM replies WHERE id = $1
`

func (q *Queries) DeleteReply(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReply, id)
	return err
}

const getCommentsForPost = `-- name: GetCommentsForPost :many
SELECT u.username, u.image, c.date_time, c.content, c.id
FROM comments AS c
INNER JOIN users AS u
ON c.user_id=u.id
WHERE post_id = $1
ORDER BY c.id
LIMIT $2
OFFSET $3
`

type GetCommentsForPostParams struct {
	PostID string `json:"post_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetCommentsForPostRow struct {
	Username sql.NullString `json:"username"`
	Image    sql.NullString `json:"image"`
	DateTime time.Time      `json:"date_time"`
	Content  string         `json:"content"`
	ID       int64          `json:"id"`
}

func (q *Queries) GetCommentsForPost(ctx context.Context, arg GetCommentsForPostParams) ([]GetCommentsForPostRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsForPost, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCommentsForPostRow{}
	for rows.Next() {
		var i GetCommentsForPostRow
		if err := rows.Scan(
			&i.Username,
			&i.Image,
			&i.DateTime,
			&i.Content,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRepliesForComment = `-- name: GetRepliesForComment :many
SELECT u.username, u.image, r.date_time, r.content, r.id
FROM replies AS r
INNER JOIN users AS u 
ON r.user_id=u.id
WHERE comment_id = $1
ORDER BY r.id
LIMIT $2
OFFSET $3
`

type GetRepliesForCommentParams struct {
	CommentID string `json:"comment_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

type GetRepliesForCommentRow struct {
	Username sql.NullString `json:"username"`
	Image    sql.NullString `json:"image"`
	DateTime time.Time      `json:"date_time"`
	Content  string         `json:"content"`
	ID       int64          `json:"id"`
}

func (q *Queries) GetRepliesForComment(ctx context.Context, arg GetRepliesForCommentParams) ([]GetRepliesForCommentRow, error) {
	rows, err := q.db.QueryContext(ctx, getRepliesForComment, arg.CommentID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRepliesForCommentRow{}
	for rows.Next() {
		var i GetRepliesForCommentRow
		if err := rows.Scan(
			&i.Username,
			&i.Image,
			&i.DateTime,
			&i.Content,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET content = $2, edited = 'true'
WHERE id = $1
RETURNING id, user_id, post_id, edited, date_time, content
`

type UpdateCommentParams struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PostID,
		&i.Edited,
		&i.DateTime,
		&i.Content,
	)
	return i, err
}

const updateReply = `-- name: UpdateReply :one
UPDATE replies
SET content = $2, edited = 'true'
WHERE id = $1
RETURNING id, user_id, comment_id, edited, date_time, content
`

type UpdateReplyParams struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func (q *Queries) UpdateReply(ctx context.Context, arg UpdateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, updateReply, arg.ID, arg.Content)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CommentID,
		&i.Edited,
		&i.DateTime,
		&i.Content,
	)
	return i, err
}
