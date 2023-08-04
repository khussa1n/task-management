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
	task.PUT("/", h.updateTask)
	task.DELETE("/", h.updateTask)

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

	priority := apiV1.Group("/priorities")
	priority.Use(h.authMiddleware())
	priority.POST("/", h.createPriority)
	priority.GET("/", h.getAllPriorities)
	priority.DELETE("/:id", h.deletePriority)

	action := apiV1.Group("/actions")
	action.Use(h.authMiddleware())
	action.POST("/", h.createAction)
	action.GET("/", h.getAllActions)
	action.DELETE("/:id", h.deleteAction)

	comment := apiV1.Group("/comments")
	comment.Use(h.authMiddleware())
	comment.POST("/", h.createComment)
	comment.GET("/", h.getAllComments)
	comment.DELETE("/:id", h.deleteComment)

	taskLog := apiV1.Group("/task-logs")
	taskLog.Use(h.authMiddleware())
	taskLog.POST("/", h.createTaskLog)
	taskLog.GET("/:taskID", h.getAllTaskLogsByTaskID)
	taskLog.PUT("/", h.updateTaskLog)

	event := apiV1.Group("/events")
	event.Use(h.authMiddleware())
	event.POST("/", h.createEvent)
	event.GET("/:taskID", h.getAllEventsByTaskID)

	return router
}
