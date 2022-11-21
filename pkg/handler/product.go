package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

type IntSlice struct {
	Ints []int
}

func convertToIntSlice(s string) ([]int, error) {
	a := &IntSlice{}
	err := json.Unmarshal([]byte(`{"Ints":`+s+"}"), a)
	return a.Ints, err
}

func (h *Handler) createProduct(c *gin.Context) {

	var userObj models.FakeProduct
	var request models.ProductRequest

	if err := c.ShouldBind(&userObj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request: "+err.Error())
		return
	}

	err := c.SaveUploadedFile(userObj.Image, "assets/images/"+userObj.Image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}

	sizes, err := convertToIntSlice(userObj.Sizes)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	request.ProductName = userObj.ProductName
	request.CategoryId = userObj.CategoryId
	request.Price = userObj.Price
	request.Color = userObj.Color
	request.Count = userObj.Count
	request.Sizes = sizes
	request.ImageUrl = "assets/images/" + userObj.Image.Filename

	fmt.Println(request)
	data, err := h.services.Product.CreateProduct(request)
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

	search := c.Query("search")

	offset, err := h.parseOffsetQueryParam(c)
	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return

	}

	limit, err := h.parseLimitQueryParam(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return

	}

	data, err := h.services.Product.GetAllProducts(models.QueryParams{
		Search: search,
		Offset: offset,
		Limit:  limit,
	})
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
	var userObj models.FakeProduct
	var request models.ProductRequest

	if err := c.ShouldBind(&userObj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request: "+err.Error())
		return
	}

	err := c.SaveUploadedFile(userObj.Image, "assets/images/"+userObj.Image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}

	sizes, err := convertToIntSlice(userObj.Sizes)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	request.ProductName = userObj.ProductName
	request.CategoryId = userObj.CategoryId
	request.Price = userObj.Price
	request.Color = userObj.Color
	request.Count = userObj.Count
	request.Sizes = sizes
	request.ImageUrl = "assets/images/" + userObj.Image.Filename

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	data, err := h.services.Product.UpdateProduct(id, request)
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
