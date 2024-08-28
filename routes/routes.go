package routes

import (
	"net/http"
	"task-app/handlers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler handlers.Operations
}

func NewRoutes(h handlers.Operations) Routes {
	return Routes{handler: h}
}

func (ro Routes) RegisterRoutes(r *gin.Engine, handler handlers.Operations) {
	CheckRoutes(r)

	// tasks
	tasks := r.Group("tasks", handler.AuthenticatedUserMiddleware())
	{
		tasks.POST("", handler.CreateTask)
		tasks.GET("", handler.GetTasks)
		tasks.GET("/:id", handler.GetTaskById)
		tasks.PUT("/:id", handler.UpdateTask)
		tasks.DELETE("/:id", handler.DeleteTask)
		tasks.PUT("/:id/complete", handler.CompletesTask)

	}

	// auth
	auth := r.Group("auth")
	{
		auth.POST("", handler.SignUp)
		auth.POST("/login", handler.Login)
	}

}

func CheckRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Task App API",
		})
	})
}
