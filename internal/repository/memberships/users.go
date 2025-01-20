package memberships

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
)

// GetUser retrieves a user from the database based on their email or username.
//
// Parameters:
//   - ctx: The context for the database query, used for cancellation or timeouts.
//   - email: The email address of the user to retrieve.
//   - username: The username of the user to retrieve.
//
// Returns:
//   - *memberships.UserModel: The user model containing user information if found, or nil if no user matches.
//   - error: An error if the query fails or the database encounters an issue.
func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ?`

	row := r.db.QueryRowContext(ctx, query, email, username)

	var user memberships.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.CreatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser inserts a new user into the database.
//
// Parameters:
//   - ctx: The context for the database operation, used for cancellation or timeouts.
//   - user: The memberships.UserModel containing the user's details to insert.
//
// Returns:
//   - error: An error if the insertion fails or the database encounters an issue.
func (r *repository) CreateUser(ctx context.Context, user memberships.UserModel) error {
	query := `INSERT INTO users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
