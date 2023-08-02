package api

import "github.com/khussa1n/blog_api/internal/entity"

type RegisterRequest struct {
	entity.User
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
