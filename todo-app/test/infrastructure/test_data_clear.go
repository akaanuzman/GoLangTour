// This Go code snippet provides a function to truncate the test data from the "todos" table in a PostgreSQL database. It uses the pgx library for database operations and the gommon log library for logging.

// Package infrastructure contains utilities for interacting with external resources such as databases.
package infrastructure

import (
	"context" // Import the context package for managing request-scoped values, cancelation signals, and deadlines.

	"github.com/jackc/pgx/v4/pgxpool" // Import the pgxpool package for PostgreSQL connection pool handling.
	"github.com/labstack/gommon/log"  // Import the log package from gommon for logging.
)

// TruncateTestData truncates the "todos" table in the database, effectively removing all test data.
func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	// Execute the TRUNCATE SQL command to remove all rows from the "todos" table and restart its identity values.
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE todos RESTART IDENTITY")
	if truncateResultErr != nil {
		// If an error occurs during the TRUNCATE operation, log the error.
		log.Error(truncateResultErr)
	} else {
		// If the operation is successful, log a confirmation message.
		log.Info("Todos table truncated")
	}
}
