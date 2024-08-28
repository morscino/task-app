package handlers

import (
	"net/http"

	"task-app/common/messages"
	"task-app/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthenticatedUserMiddleware() gin.HandlerFunc {
	// add the middleware function
	return func(c *gin.Context) {
		user, err := h.controller.Middleware().JwtUserAuth(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
			c.Abort()
		} else {
			c.Set("authUser", user)
		}
		c.Next()
	}
}
