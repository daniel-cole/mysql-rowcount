package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Setup(mysqlConfig *mysql.Config) {

	database, err := sql.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	DB = database
}
