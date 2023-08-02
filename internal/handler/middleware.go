package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/khussa1n/task-management/api"
	"net/http"
	"strings"
)

const authUserID = "auth_user_id"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")
		if authorizationHeader == "" {
			err := errors.New("authorization header is not set")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -1,
				Message: err.Error(),
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -2,
				Message: err.Error(),
			})
			return
		}

		userID, err := h.srvs.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -3,
				Message: err.Error(),
			})
			return
		}

		ctx.Set(authUserID, userID)
		ctx.Next()
	}
}
