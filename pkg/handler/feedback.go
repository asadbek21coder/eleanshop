package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create Feedback
// @Security ApiKeyAuth
// @Tags feedback
// @Description create feedback
// @ID create-feedback
// @Accept  json
// @Produce  json
// @Param input body models.UpdateFeedbackInput true "feedback info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /feedback [post]
func (h *Handler) createFeedback(c *gin.Context) {
	var input models.UpdateFeedbackInput
	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(int)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id type")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Feedback.CreateFeedback(input, userIdInt)
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

func (h *Handler) getFeedbackById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.Feedback.GetFeedbackById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "OK",
	})
}

func (h *Handler) getAllFeedbacks(c *gin.Context) {
	data, err := h.services.Feedback.GetAllFeedbacks()
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

func (h *Handler) updateFeedback(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(int)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id type")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateFeedbackInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Feedback.UpdateFeedback(id, input, &userIdInt)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "OK",
	})
}

// @Summary Delete Feedback
// @Security ApiKeyAuth
// @Tags feedback
// @Description delete feedback by given id
// @ID delete-feedback
// @Accept  json
// @Produce  json
// @Param id path string true "feedback id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /feedback/{id} [delete]
func (h *Handler) deleteFeedback(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdInt, ok := userId.(int)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id type")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.Feedback.DeleteFeedback(id, userIdInt)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    id,
		"isOk":    true,
		"message": "OK",
	})
}
