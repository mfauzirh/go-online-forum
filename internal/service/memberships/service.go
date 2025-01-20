package memberships

import (
	"context"

	"github.com/mfauzirh/go-online-forum/internal/configs"
	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
)

// membershipRepository defines the interface for user-related database operations.
//
// Methods:
//   - GetUser: Retrieves a user based on email or username.
//   - CreateUser: Inserts a new user into the database.
type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user memberships.UserModel) error
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
