package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input models.ProductRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Product.CreateProduct(input)
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

func (h *Handler) getProductById(c *gin.Context) {
	var id int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	data, err := h.services.Product.GetProductById(id)
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

func (h *Handler) getAllProducts(c *gin.Context) {

	data, err := h.services.Product.GetAllProducts()
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

func (h *Handler) updateProduct(c *gin.Context) {
	var id int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input models.ProductRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Product.UpdateProduct(id, input)
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

func (h *Handler) deleteProduct(c *gin.Context) {
	var id int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	data, err := h.services.Product.DeleteProduct(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "Deleted",
	})
}
