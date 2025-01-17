package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/configs"
	"github.com/mfauzirh/go-online-forum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("Failed to initalize config: %v", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
