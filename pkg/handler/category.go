package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryPost struct {
	Name string `json:"name"`
}

// @Summary     Create Category
// @Security    ApiKeyAuth
// @Tags        category
// @Description create category
// @ID          create-category
// @Accept      json
// @Produce     json
// @Param       input   body      CategoryPost true "category info"
// @Success     200     {integer} integer      1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /admin/category [post]
func (h *Handler) createCategory(c *gin.Context) {
	var input CategoryPost

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.CreateCategory(input.Name)
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

// @Summary     Get Category By Id
// @Tags        category
// @Description get category by given id
// @ID          get-category-by-id
// @Produce     json
// @Param       id      path     string true "category id"
// @Success     200     {object} models.Category
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /admin/category/{id} [get]
func (h *Handler) getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.GetCategoryById(id)
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

// @Summary     Get All Categories
// @Tags        category
// @Description get all categories
// @ID          get-all-categories
// @Accept      json
// @Produce     json
// @Success     200     {object} []models.Category
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /admin/category [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	data, err := h.services.GetAllCategories()
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

// @Summary     Update Category
// @Security    ApiKeyAuth
// @Tags        category
// @Description update category
// @ID          update-category
// @Accept      json
// @Produce     json
// @Param       input   body      CategoryPost true "category info"
// @Param       id      path      string       true "category id"
// @Success     200     {integer} integer      1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /admin/category/{id} [put]
func (h *Handler) updateCategory(c *gin.Context) {
	var input CategoryPost
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.UpdateCategory(id, input.Name)
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

// @Summary     Delete Category
// @Security    ApiKeyAuth
// @Tags        category
// @Description delete category by given id
// @ID          delete-category
// @Accept      json
// @Produce     json
// @Param       id      path      string  true "category id"
// @Success     200     {integer} integer 1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /admin/category/{id} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.DeleteCategory(id)
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
