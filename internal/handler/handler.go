package handler

import "github.com/khussa1n/task-management/internal/service"

type Handler struct {
	srvs service.Service
}

func New(srvs service.Service) *Handler {
	return &Handler{
		srvs: srvs,
	}
}
