package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/khussa1n/task-management/internal/config"
	"github.com/khussa1n/task-management/internal/entity"
	mock_repository "github.com/khussa1n/task-management/internal/repository/mock"
	"github.com/khussa1n/task-management/pkg/jwttoken"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	cfg, err := config.InitConfig("../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("adsfkasdfk")

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := mock_repository.NewMockRepository(controller)

	ctx := context.Background()
	u := &entity.Users{
		ID:        1,
		Email:     "abc",
		FirstName: "abc",
		LastName:  "abc",
		Password:  "abc",
	}

	u2 := &entity.Users{
		Email:     "abadfc",
		FirstName: "aadsfbc",
		LastName:  "abasdfc",
		Password:  "aafdbc",
	}

	mockRepo.EXPECT().CreateUser(ctx, u).Return(u, nil).Times(1)
	mockRepo.EXPECT().CreateUser(ctx, u2).Return(u2, nil).Times(1)

	service := New(mockRepo, jwtToken, cfg)

	user, err := service.CreateUser(ctx, u)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	user, err = service.CreateUser(ctx, u2)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
