package memberships

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// repository is a private struct that holds the database connection.
// It is used to interact with the database for CRUD operations related to memberships.
type repository struct {
	db *sql.DB
}

// NewRepository creates a new instance of the repository struct.
// It accepts a pointer to a database connection and returns a repository object.
//
// Parameters:
//   - db (*sql.DB): The database connection to be used by the repository.
//
// Returns:
//   - *repository: A new instance of the repository struct with the database connection.
func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}
