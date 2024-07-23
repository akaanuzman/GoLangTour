// This Go code snippet provides a function to initialize test data in the "todos" table of a PostgreSQL database. It uses the pgx library for database operations and the gommon log library for logging.

// Package infrastructure contains utilities for interacting with external resources such as databases.
package infrastructure

import (
	"context" // Import the context package for managing request-scoped values, cancelation signals, and deadlines.
	"fmt"     // Import the fmt package for formatting strings.

	"github.com/jackc/pgx/v4/pgxpool" // Import the pgxpool package for PostgreSQL connection pool handling.
	"github.com/labstack/gommon/log"  // Import the log package from gommon for logging.
)

// INSERT_TODOS contains the SQL query to insert initial test data into the "todos" table.
var INSERT_TODOS = `INSERT INTO todos (title, description, created_at, updated_at)
VALUES('Buy Milk', 'Buy 1lt milk from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Bread', 'Buy 1 loaf of bread from the bakery', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Eggs', 'Buy 1 dozen eggs from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Butter', 'Buy 1 stick of butter from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
`

// TestDataInitialize inserts initial test data into the "todos" table using the INSERT_TODOS query.
func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	// Execute the INSERT_TODOS SQL command to insert test data into the "todos" table.
	insertTodosResult, insertTodosErr := dbPool.Exec(ctx, INSERT_TODOS)
	if insertTodosErr != nil {
		// If an error occurs during the insert operation, log the error.
		log.Error(insertTodosErr)
	} else {
		// If the operation is successful, log the number of rows affected.
		log.Info(fmt.Sprintf("Todos data created with %d rows", insertTodosResult.RowsAffected()))
	}
}
