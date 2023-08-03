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
	user.GET("/:id", h.getUserByID)
	user.PUT("/", h.updateUser)
	user.DELETE("/", h.deleteUser)

	task := apiV1.Group("/tasks")
	task.Use(h.authMiddleware())
	task.POST("/", h.createTask)
	task.GET("/", h.getAllTasks)
	//task.GET("/:id")
	//task.PUT("/")
	//task.DELETE("/")

	status := apiV1.Group("/statuses")
	status.Use(h.authMiddleware())
	status.POST("/", h.createStatus)
	status.GET("/", h.getAllStatuses)
	status.DELETE("/:id", h.deleteStatus)

	role := apiV1.Group("/roles")
	role.Use(h.authMiddleware())
	role.POST("/", h.createRole)
	role.GET("/", h.getAllRoles)
	role.DELETE("/:id", h.deleteRole)

	return router
}
