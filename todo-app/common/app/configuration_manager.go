// Package app provides the core application configurations and utilities.
package app

import "todoapp/common/postresql"

// ConfigurationManager holds the configuration settings for the application,
// including database connection information.
type ConfigurationManager struct {
	// PostreSqlConfig contains the PostgreSQL database configuration.
	PostreSqlConfig postresql.PostreSQLConfig
}

// NewConfigurationManager creates and returns a new instance of ConfigurationManager.
// It initializes the PostgreSQL configuration by calling getPostgreSqlConfig.
// Returns a pointer to the newly created ConfigurationManager.
func NewConfigurationManager() *ConfigurationManager {
	postreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostreSqlConfig: postreSqlConfig,
	}
}

// getPostgreSqlConfig initializes and returns a PostreSQLConfig struct with
// predefined values for connecting to a PostgreSQL database.
// This includes the host, port, username, password, database name,
// maximum number of connections, and maximum idle time for connections.
// Returns the configured PostreSQLConfig.
func getPostgreSqlConfig() postresql.PostreSQLConfig {
	return postresql.PostreSQLConfig{
		Host:                  "localhost", // Database server host
		Port:                  "6432",      // Database server port
		Username:              "postgres",  // Database username
		Password:              "postgres",  // Database password
		DbName:                "todoapp",   // Database name
		MaxConnections:        "10",        // Maximum number of connections
		MaxConnectionIdleTime: "30s",       // Maximum idle time for connections
	}
}
