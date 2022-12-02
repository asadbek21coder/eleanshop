package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

// @Summary     Create Size
// @Security    ApiKeyAuth
// @Tags        size
// @Description create size by given id
// @ID          create-size
// @Accept      json
// @Produce     json
// @Param       input   body      models.SizeInput true "size info"
// @Success     200     {integer} integer          1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /admin/sizes [post]
func (h *Handler) createSize(c *gin.Context) {
	var input models.SizeInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Size.CreateSize(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "OK",
	})
}


// @Summary     Get All Size
// @Security    ApiKeyAuth
// @Tags        size
// @Description get all sizes
// @ID          get-all-sizes
// @Accept      json
// @Produce     json
// @Success     200     {object} []models.Size
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /admin/sizes [get]
func (h *Handler) getAllSizes(c *gin.Context) {

	data, err := h.services.Size.GetAllSize()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("working")

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "OK",
	})
}

// @Summary     Delete Size
// @Security    ApiKeyAuth
// @Tags        size
// @Description delete size by given id
// @ID          delete-size
// @Accept      json
// @Produce     json
// @Param       id      path      string  true "size id"
// @Success     200     {integer} integer 1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /admin/sizes/{id} [delete]
func (h *Handler) deleteSize(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Size.DeleteSize(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isOk":    true,
		"message": "Deleted",
	})
}

// @Summary     Get Size By Id
// @Security    ApiKeyAuth
// @Tags        size
// @Description get size by given id
// @ID          get-size-by-id
// @Produce     json
// @Param       id      path     string true "size id"
// @Success     200     {object} models.Size
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /admin/sizes/{id} [get]
func (h *Handler) getSizesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.Size.GetSizesById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isOk":    data,
		"message": "Ok",
	})
}
