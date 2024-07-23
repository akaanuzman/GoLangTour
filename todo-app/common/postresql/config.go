package postresql

// PostreSQLConfig struct
// This struct is used to store the configuration of the PostgreSQL database
// Host: The host of the PostgreSQL database
// Port: The port of the PostgreSQL database
// Username: The username of the PostgreSQL database
// Password: The password of the PostgreSQL database
// DbName: The name of the PostgreSQL database
// MaxConnections: The maximum number of connections to the PostgreSQL database
// MaxConnectionIdleTime: The maximum idle time of the connections to the PostgreSQL database
type PostreSQLConfig struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime string
}
