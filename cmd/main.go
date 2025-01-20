package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/configs"
	"github.com/mfauzirh/go-online-forum/internal/handlers/memberships"
	"github.com/mfauzirh/go-online-forum/pkg/internalsql"

	membershipRepo "github.com/mfauzirh/go-online-forum/internal/repository/memberships"
	membershipSvc "github.com/mfauzirh/go-online-forum/internal/service/memberships"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}

	membershipRepository := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepository)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
