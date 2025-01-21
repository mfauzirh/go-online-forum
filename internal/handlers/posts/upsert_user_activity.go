package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/model/posts"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("postID tidak valid pada params"),
		})
		return
	}

	userID := c.GetInt64("userID")

	if err := h.postSvc.UpsertUserActivity(ctx, postID, userID, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("terjadi kesalahan server"),
		})
		return
	}

	c.Status(http.StatusOK)
}
