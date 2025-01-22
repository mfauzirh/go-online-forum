package memberships

import (
	"context"
	"time"

	"github.com/mfauzirh/go-online-forum/internal/configs"
	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
)

// membershipRepository defines the interface for user-related database operations.
//
// Methods:
//   - GetUser: Retrieves a user based on email or username.
//   - CreateUser: Inserts a new user into the database.
type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user memberships.UserModel) error
	InsertRefereshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

// service provides higher-level operations related to memberships.
//
// Fields:
//   - membershipRepo: The repository implementing membership-related database operations.
type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

// NewService creates a new instance of the membership service.
//
// Parameters:
//   - membershipRepo: An implementation of the membershipRepository interface.
//
// Returns:
//   - *service: A new service instance to handle membership-related operations.
func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
