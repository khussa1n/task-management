package service

import (
	"github.com/khussa1n/task-management/internal/config"
	"github.com/khussa1n/task-management/internal/repository"
	"github.com/khussa1n/task-management/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.JWTToken
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.JWTToken, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
	}
}
