package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("postID tidak valid pada params"),
		})
		return
	}

	response, err := h.postSvc.GetPostByID(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("terjadi kesalahan server"),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
