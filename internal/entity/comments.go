package entity

import "time"

type Comments struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id" binding:"required"`
	CreatedDate     time.Time `json:"created_date" db:"created_date" binding:"required"`
	Comment         string    `json:"comment" db:"comment" binding:"required"`
	ParentCommentID int64     `json:"parent_comment_id" db:"parent_comment_id" binding:"required"`
}
