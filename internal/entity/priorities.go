package entity

type Priorities struct {
	ID           int64  `json:"id" db:"id"`
	PriorityName string `json:"priority_name" db:"priority_name" binding:"required"`
}
