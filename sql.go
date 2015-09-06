package rets

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfiguration struct {
	Host     string `json:"host"`
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func OpenDatabase(config DatabaseConfiguration) (db *sql.DB, err error) {
	db, err = sql.Open(config.Driver, config.Username+":"+config.Password+
		"@tcp("+config.Host+")/"+config.Database)
	if err != nil {
		return
	}
	err = db.Ping()
	return
}
