package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var SqlClient *sql.DB

func GetSqlConnection() (*sql.DB, error) {
	db, errorDb := sql.Open("mysql", "ali:aA@09380962410@/nothing")
	failOnError(errorDb, "Failed to open mysql")
	SqlClient = db
	return db, errorDb
}
