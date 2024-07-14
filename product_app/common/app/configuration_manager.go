package app

import "golangtour/common/postresql"

type ConfigurationManager struct {
	PostreSqlConfig postresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostreSqlConfig: postreSqlConfig,
	}
}

func getPostgreSqlConfig() postresql.Config {
	return postresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Username:              "postgres",
		Password:              "postgres",
		DbName:                "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
