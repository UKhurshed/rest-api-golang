package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	name, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"name": name,
	})
}
