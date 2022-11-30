package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/eleanshop/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create Question
// @Security ApiKeyAuth
// @Tags question
// @Description create question by given id
// @ID create-question
// @Accept  json
// @Produce  json
// @Param input body models.UpdateQuestionInput true "question info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /question [post]
func (h *Handler) createQuestion(c *gin.Context) {
	var input models.UpdateQuestionInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Question.CreateQuestion(input)
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

// @Summary Get Question By Id
// @Security ApiKeyAuth
// @Tags question
// @Description get question by given id
// @ID get-question-by-id
// @Produce  json
// @Param id path string true "question id"
// @Success 200 {object} models.Question
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /question/{id} [get]
func (h *Handler) getQuestionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.Question.GetQuestionById(id)
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

// @Summary Get All Questions
// @Tags question
// @Description get all questions
// @ID get-all-questions
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Question
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /question [get]
func (h *Handler) getAllQuestions(c *gin.Context) {
	data, err := h.services.Question.GetAllQuestions()
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

// @Summary Update Question
// @Security ApiKeyAuth
// @Tags question
// @Description update question
// @ID update-question
// @Accept  json
// @Produce  json
// @Param input body models.UpdateQuestionInput true "question info"
// @Param id path string true "question id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /question/{id} [put]
func (h *Handler) updateQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateQuestionInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.Question.UpdateQuestion(id, input)
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

// @Summary Delete Question
// @Security ApiKeyAuth
// @Tags question
// @Description delete question by given id
// @ID delete-question
// @Accept  json
// @Produce  json
// @Param id path string true "question id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /question/{id} [delete]
func (h *Handler) deleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.Question.DeleteQuestion(id)
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
