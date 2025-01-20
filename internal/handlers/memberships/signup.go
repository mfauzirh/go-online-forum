package memberships

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.membershipSvc.SignUp(ctx, request); err != nil {
		log.Printf("Error ocurred: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusCreated)
}
