package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

// SignUp handles user registration by checking for existing users,
// hashing the password, and saving the new user.
//
// Parameters:
//   - ctx: Context for managing request lifecycle, such as timeouts or cancellations.
//   - req: A memberships.SignUpRequest containing the user's email, username, and password.
//
// Returns:
//   - error: Returns an error if validation, password hashing, or database insertion fails.
func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	existingUser, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("username or email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	user := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, user)
}
