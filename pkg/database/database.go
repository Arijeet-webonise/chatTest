package database

import (
	"database/sql"
	"fmt"

	"github.com/knq/dburl"
	_ "github.com/lib/pq"
)

// DatabaseConnectionInitialiser encapsilate database connection
type DatabaseConnectionInitialiser interface {
	InitialiseConnection() (*sql.DB, error)
}

// DatabaseConfig wrapper for DB data
type DatabaseConfig struct {
	//DB *sql.DB
	Protocol         string
	Username         string
	Password         string
	Host             string
	DatabaseName     string
	ConnectionParams string
}

// InitialiseConnection inicilize db
func (dw *DatabaseConfig) InitialiseConnection() (*sql.DB, error) {
	s := fmt.Sprintf("%s://%s:%s@%s/%s?%s", dw.Protocol, dw.Username, dw.Password, dw.Host, dw.DatabaseName, dw.ConnectionParams)
	u, err := dburl.Open(s)
	return u, err
}
