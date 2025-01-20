package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/configs"
	"github.com/mfauzirh/go-online-forum/internal/handlers/memberships"
	"github.com/mfauzirh/go-online-forum/internal/handlers/posts"
	"github.com/mfauzirh/go-online-forum/pkg/internalsql"

	membershipRepo "github.com/mfauzirh/go-online-forum/internal/repository/memberships"
	postRepo "github.com/mfauzirh/go-online-forum/internal/repository/posts"
	membershipSvc "github.com/mfauzirh/go-online-forum/internal/service/memberships"
	postSvc "github.com/mfauzirh/go-online-forum/internal/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepository := membershipRepo.NewRepository(db)
	postRepository := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepository)
	postService := postSvc.NewService(cfg, postRepository)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
