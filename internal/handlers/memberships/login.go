package memberships

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		log.Printf("Error ocurred: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	response := memberships.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, response)
}
