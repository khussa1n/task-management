package api

import "github.com/khussa1n/task-management/internal/entity"

type RegisterRequest struct {
	entity.Users
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
