package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	auth := apiV1.Group("/auth")
	auth.POST("/sign-up", h.createUser)
	auth.POST("/sign-in", h.login)

	user := apiV1.Group("/users")
	user.Use(h.authMiddleware())
	user.GET("/:id", h.GetUserByID)
	user.POST("/", h.UpdateUser)
	user.DELETE("/", h.DeleteUser)

	return router
}
