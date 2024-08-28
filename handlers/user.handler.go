package handlers

import (
	"net/http"

	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary Create new user
// @Schemes
// @Description Creates a new user
// @Param   request   body     models.SignUpDto   true  "data to sign up new user"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseObject "desc"
// @Router /auth [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input models.SignUpDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {

		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}

	// send to controller
	result := h.controller.RegisterUser(&input)
	c.JSON(result.Code, result)
}

// @Tags User
// @Summary Login  user
// @Schemes
// @Description Logs in a user
// @Param   request   body     models.SignInDto   true  "data to log in a user"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseObject "desc"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input models.SignInDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	// send to controller
	result := h.controller.Login(&input)
	c.JSON(200, result)
}
