package entity

type Users struct {
	ID        int64  `json:"id" db:"id"`
	Email     string `json:"email" db:"email" binding:"required"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Password  string `json:"password" db:"hashed_password" binding:"required"`
}
