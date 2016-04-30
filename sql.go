package rets

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // SQL Driver must be blank
)

// DatabaseConfiguration represents a MySQL configuration that allows the user
// to specify the host, the username/password and the database to use when
// using the ReTS SQL API.
type DatabaseConfiguration struct {
	Host     string `json:"host"`
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// OpenDatabase opens a database connection from a given DatabaseConfiguration.
func OpenDatabase(config DatabaseConfiguration) (db *sql.DB, err error) {
	db, err = sql.Open(config.Driver, config.Username+":"+config.Password+
		"@tcp("+config.Host+")/"+config.Database+"?parseTime=true")
	if err != nil {
		return
	}
	err = db.Ping()
	return
}
