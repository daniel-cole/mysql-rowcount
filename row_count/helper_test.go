package row_count

import (
	"github.com/daniel-cole/mysql-rowcount/config"
	"github.com/daniel-cole/mysql-rowcount/db"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testConfig *config.RowCountConfig

func TestMain(m *testing.M) {

	if ok := os.Setenv("ROWCOUNT_CONFIG", filepath.Join("testdata", "rowcount-config.yaml")); ok != nil {
		log.Fatal(ok)
	}

	testConfig = &config.RowCountConfig{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1",
		Net:                  "tcp",
		AllowNativePasswords: true,
		DatabasesToIgnore:    []string{},
		DatabasesToInclude:   []string{},
		TablesToIgnore:       []string{},
		MaxWorkers:           2,
		OutputFile:           "test1234",
	}

	db.Setup(&mysql.Config{
		User:                 testConfig.User,
		Passwd:               testConfig.Passwd,
		Addr:                 testConfig.Addr,
		Net:                  testConfig.Net,
		AllowNativePasswords: testConfig.AllowNativePasswords,
	})
	defer db.DB.Close()

	loadTestData()

	os.Exit(m.Run())
}

func loadTestData() {
	sql := string(loadTestFile("db1.sql"))
	preparedSQL := strings.Split(string(sql), ";")
	for _, query := range preparedSQL {

		query = strings.TrimRight(query, "\r\n") // trim any newline characters to prevent empty queries
		if query == "" {
			continue
		}

		_, err := db.DB.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func loadTestFile(name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
