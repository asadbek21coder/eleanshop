package handler

import (
	"net/http"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

// @Summary     SignUp
// @Tags        auth
// @Description create account
// @ID          create-account
// @Accept      json
// @Produce     json
// @Param       input   body      models.User true "account info"
// @Success     200     {integer} integer     1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /auth/signup [post]
func (h *Handler) signup(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "OK",
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary     SignIn
// @Tags        auth
// @Description login
// @ID          login
// @Accept      json
// @Produce     json
// @Param       input   body     signInInput true "credentials"
// @Success     200     {string} string      "token"
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /auth/signin [post]
func (h *Handler) signin(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

// @Summary     Set Admin
// @Security    ApiKeyAuth
// @Tags        auth
// @Description login
// @ID          admin
// @Accept      json
// @Produce     json
// @Param       input   body     models.SetAdmin true "credentials"
// @Success     200     {string} string          "Message"
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /admin/set-admin [put]
func (h *Handler) setAdmin(c *gin.Context) {
	var input models.SetAdmin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.SetAdmin(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "OK",
	})

}
