package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mfauzirh/go-online-forum/internal/middleware"
	"github.com/mfauzirh/go-online-forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())
	route.POST("", h.CreatePost)
	route.POST("/:postID/comments", h.CreateComment)
	route.PUT("/:postID/user-activities", h.UpsertUserActivity)
	route.GET("", h.GetAllPost)
}
