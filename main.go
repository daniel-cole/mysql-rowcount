package main

import (
	"fmt"
	"github.com/daniel-cole/mysql-rowcount/config"
	"github.com/daniel-cole/mysql-rowcount/db"
	"github.com/daniel-cole/mysql-rowcount/row_count"
	"github.com/daniel-cole/mysql-rowcount/util"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var version = "undefined"

var rowCountConfig *config.RowCountConfig

func main() {

	var arg string

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	if arg == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	rowCountConfig = config.LoadConfig()

	if util.CheckFileExists(rowCountConfig.OutputFile) {
		log.Fatalf("Output file already exists: %s", rowCountConfig.OutputFile)
	}

	db.Setup(&mysql.Config{
		User:                 rowCountConfig.User,
		Passwd:               rowCountConfig.Passwd,
		Addr:                 rowCountConfig.Addr,
		Net:                  rowCountConfig.Net,
		AllowNativePasswords: rowCountConfig.AllowNativePasswords,
	})
	defer db.DB.Close()

	for _, tableRowCount := range row_count.FetchDatabaseTableRowCount(rowCountConfig) {
		util.WriteBytesToFile(rowCountConfig.OutputFile, []byte(fmt.Sprintf("%s %d\n", tableRowCount.Table, tableRowCount.RowCount)))
	}

}
