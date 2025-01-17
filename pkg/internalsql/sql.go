package internalsql

import (
	"database/sql"
	"log"
)

// Connect establishes a connection to a MySQL database using the provided data source name.
// It returns the established database connection or an error if the connection fails.
//
// Parameters:
//   - dataSourceName (string): The data source name used to connect to the MySQL database.
//
// Returns:
//   - (*sql.DB, error): A pointer to the established *sql.DB connection and any error encountered.
func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}
	return db, nil
}
