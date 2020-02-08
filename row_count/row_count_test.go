package row_count

import (
	"github.com/daniel-cole/mysql-rowcount/util"
	"testing"
)

func TestGetTableRowCount(t *testing.T) {

	config := testConfig
	config.DatabasesToInclude = []string{"db1_test"}

	tableRowCount := FetchDatabaseTableRowCount(config)
	for _, table := range tableRowCount {
		if table.Table == "db1_test.test2" {
			expectedRowCount := 2
			if table.RowCount != expectedRowCount {
				t.Errorf("Expected row count for %s to be 5, instead got: %d", table.Table, table.RowCount)
			}
		}

		if table.Table == "db1_test.test1" {
			expectedRowCount := 5
			if table.RowCount != expectedRowCount {
				t.Errorf("Expected row count for %s to be 5, instead got: %d", table.Table, table.RowCount)
			}
		}
	}
}

func TestCompileDatabaseTableList(t *testing.T) {
	config := testConfig
	config.DatabasesToInclude = []string{"db1_test"}

	databaseTableList := buildDatabaseTableList(config)

	if !util.StringInSlice("db1_test.test1", databaseTableList) {
		t.Error("Expected table to exist in databaseTableList")
	}

	if !util.StringInSlice("db1_test.test2", databaseTableList) {
		t.Error("Expected table to exist in databaseTableList")
	}

	if util.StringInSlice("thisdoesntexist", databaseTableList) {
		t.Error("Did not expect table to exist in databaseTableList")
	}

}

func TestCompileDatabaseTableListIgnoreTable(t *testing.T) {
	config := testConfig
	config.DatabasesToInclude = []string{"db1_test"}
	config.TablesToIgnore = []string{"test2"}

	databaseTableList := buildDatabaseTableList(config)

	if !util.StringInSlice("db1_test.test1", databaseTableList) {
		t.Error("Expected table to exist in databaseTableList")
	}

	if util.StringInSlice("db1_test.test2", databaseTableList) {
		t.Error("Expected table to not exist in databaseTableList")
	}

	if util.StringInSlice("thisdoesntexist", databaseTableList) {
		t.Error("Did not expect table to exist in databaseTableList")
	}

}
