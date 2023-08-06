package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/khussa1n/task-management/api"
	"github.com/khussa1n/task-management/internal/entity"
	mock_service "github.com/khussa1n/task-management/internal/service/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := mock_service.NewMockService(controller)

	handler := New(mockService)

	recorder := httptest.NewRecorder()

	req := api.RegisterRequest{
		Users: entity.Users{
			ID:        1,
			Email:     "asdf",
			FirstName: "afd",
			LastName:  "adsf",
			Password:  "asdf",
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	require.NoError(t, err)

	u := &entity.Users{
		ID:        1,
		Email:     "asdf",
		FirstName: "afd",
		LastName:  "adsf",
		Password:  "asdf",
	}

	mockService.EXPECT().CreateUser(gomock.Any(), u).Return(u, nil).Times(1)

	url := fmt.Sprintf("/api/v1/auth/sign-up")
	request, err := http.NewRequest(http.MethodPost, url, &buf)
	require.NoError(t, err)

	handler.InitRouter().ServeHTTP(recorder, request)

	require.Equal(t, http.StatusCreated, recorder.Code)
}
