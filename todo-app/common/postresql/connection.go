// Package postresql provides utilities for working with PostgreSQL databases,
// including connection pool management.
package postresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// GetConnectionPool establishes a connection pool to a PostgreSQL database using the provided configuration.
// It constructs a connection string from the PostreSQLConfig struct and uses it to create a pgxpool.Pool.
//
// Parameters:
// - context: A context.Context to control the lifetime of the connection pool creation request.
// - config: A PostreSQLConfig struct containing the database connection configuration.
//
// Returns:
// - A pointer to a pgxpool.Pool, which represents the established connection pool.
//
// Panics:
//   - If parsing the connection configuration fails or if establishing the connection pool fails,
//     the function will log the error and panic.
func GetConnectionPool(context context.Context, config PostreSQLConfig) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime,
	)

	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}

	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Error("Unable to connect to database:", err)
		panic(err)
	}

	return conn
}
