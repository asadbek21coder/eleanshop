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
	// fmt.Println(s)
	a := &IntSlice{}
	// fmt.Println(a)

	err := json.Unmarshal([]byte(`{"Ints":`+s+"}"), a)
	// fmt.Println(a.Ints)
	return a.Ints, err
}

// @Summary     Create Product
// @Security    ApiKeyAuth
// @Tags        products
// @Description create product
// @ID          create-product
// @Accept      multipart/form-data
// @Produce     json
// @Param       form   formData  models.FakeProduct true "product info"
// @Param       file   formData  file true "image info"
// @Success     200     {integer} integer            1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /product [post]
func (h *Handler) createProduct(c *gin.Context) {

	var userObj models.FakeProduct
	var request models.ProductRequest

	if err := c.ShouldBind(&userObj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body1: "+err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body2: "+err.Error())
		return
	}
	err = c.SaveUploadedFile(file, "assets/images/"+file.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body3: "+err.Error())
		return
	}

	sizes, err := convertToIntSlice(userObj.Sizes)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body4: "+err.Error())
		return
	}
	request.ProductName = userObj.Name
	request.CategoryId = userObj.Category
	request.Price = userObj.Price
	request.Color = userObj.Color
	request.Count = userObj.Count
	request.Sizes = sizes
	request.ImageUrl = "assets/images/" + file.Filename

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

// @Summary     Get Product By Id
// @Tags        products
// @Description get product by given id
// @ID          get-product-by-id
// @Produce     json
// @Param       id      path     string true "product id"
// @Success     200     {object} models.Product
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /product/{id} [get]
func (h *Handler) getProductById(c *gin.Context) {
	var id int

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id")
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

// @Summary     Get All Products
// @Tags        products
// @Description get all products
// @ID          get-all-products
// @Accept      json
// @Produce     json
// @Param       search  query    string false "search"
// @Param       limit   query    string false "limit"
// @Param       offset  query    string false "offset"
// @Success     200     {object} []models.Product
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /product [get]
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

// @Summary     Update Product
// @Security    ApiKeyAuth
// @Tags        products
// @Description update product
// @ID          update-product
// @Accept      multipart/form-data
// @Produce     json
// @Param       id      path      string       true "product id"
// @Param       form   formData  models.FakeProduct true "product info"
// @Param       file   formData  file true "image info"
// @Success     200     {integer} integer            1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /product/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	var userObj models.FakeProduct
	var request models.ProductRequest

	if err := c.ShouldBind(&userObj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body: "+err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body: "+err.Error())
		return
	}
	err = c.SaveUploadedFile(file, "assets/images/"+file.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		fmt.Println("error: ", err)
	}
	sizes, err := convertToIntSlice(userObj.Sizes)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id param")
		return
	}
	request.ProductName = userObj.Name
	request.CategoryId = userObj.Category
	request.Price = userObj.Price
	request.Color = userObj.Color
	request.Count = userObj.Count
	request.Sizes = sizes
	request.ImageUrl = "assets/images/" + file.Filename

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

// @Summary     Delete Product
// @Security    ApiKeyAuth
// @Tags        products
// @Description delete product by given id
// @ID          delete-products
// @Accept      json
// @Produce     json
// @Param       id      path      string  true "product id"
// @Success     200     {integer} integer 1
// @Failure     400,404 {object}  errorResponse
// @Failure     500     {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /product/{id} [delete]
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
