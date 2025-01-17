package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()

	r.Run(":8080")
}
