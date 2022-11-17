package handler

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	// "github.com/asadbek21coder/eleanshop/models"
	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name   string                `form:"name"`
	Email  string                `form:"email"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
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

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   userObj,
	})

	// // data, err := h.services.Product.CreateProduct(input)
	// // if err != nil {
	// // 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// // 	return
	// // }
	// fmt.Print(response)

	// err := c.SaveUploadedFile(response.Image, response.Image.Filename)
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, "unknown error")
	// 	return
	// }

	// c.JSON(http.StatusCreated, map[string]interface{}{
	// 	"data":    response,
	// 	"isOk":    true,
	// 	"message": "OK",
	// })

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
