package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.Abort()
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, isAdmin, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	// fmt.Println(userId)
	c.Set("userId", userId)
	c.Set("isAdmin", isAdmin)
}

// func getUserId(c *gin.Context) (int, error) {
// 	id, ok := c.Get("userId")
// 	if !ok {
// 		return 0, errors.New("user id not found")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		return 0, errors.New("user id is of invalid type")
// 	}

// 	return idInt, nil
// }

func (h *Handler) isAdmin(c *gin.Context) {
	h.userIdentity(c)
	isAdmin, ok := c.Get("isAdmin")
	if !ok {
		return
	}

	isAdm, ok := isAdmin.(bool)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "isAdmin is invalid type")
		return
	}

	if !isAdm {
		newErrorResponse(c, http.StatusUnauthorized, "you are not admin")
		return
	}

}
