package entity

type Statuses struct {
	ID         int64  `json:"id" db:"id"`
	StatusName string `json:"status_name" db:"status_name" binding:"required"`
}
