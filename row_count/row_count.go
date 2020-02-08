package row_count

import (
	"fmt"
	"github.com/daniel-cole/mysql-rowcount/config"
	"github.com/daniel-cole/mysql-rowcount/db"
	"log"
	"sort"
	"sync"
)

type TableRowCount struct {
	Table    string
	RowCount int
}

func FetchDatabaseTableRowCount(config *config.RowCountConfig) []TableRowCount {
	var tableRowCount []TableRowCount
	var wg sync.WaitGroup

	databaseTableList := buildDatabaseTableList(config)

	limit := make(chan struct{}, config.MaxWorkers)
	defer close(limit)

	rowCountChan := make(chan TableRowCount)
	for _, databaseTable := range databaseTableList {
		wg.Add(1)
		go func(databaseTable string) {
			defer wg.Done()
			limit <- struct{}{}
			log.Printf("Retrieving row count for: %s", databaseTable)
			rowCountChan <- TableRowCount{
				Table:    databaseTable,
				RowCount: countRows(databaseTable),
			}

			<-limit
		}(databaseTable)
	}

	go func() {
		wg.Wait()
		close(rowCountChan)
	}()

	for results := range rowCountChan {
		tableRowCount = append(tableRowCount, results)
	}

	sort.Slice(tableRowCount, func(i, j int) bool {
		return tableRowCount[i].Table < tableRowCount[j].Table
	})

	return tableRowCount
}

func buildDatabaseTableList(config *config.RowCountConfig) []string {

	databasesToIgnore := make(map[string]bool)
	databasesToInclude := make(map[string]bool)
	tablesToIgnore := make(map[string]bool)

	for _, database := range config.DatabasesToInclude {
		databasesToInclude[database] = true
	}

	for _, database := range config.DatabasesToIgnore {
		databasesToIgnore[database] = true
	}

	for _, table := range config.TablesToIgnore {
		tablesToIgnore[table] = true
	}

	var databaseTableList []string

	databases := discoverDatabases()
	for _, database := range databases {
		log.Println("Checking database: " + database)
		if databasesToIgnore[database] {
			log.Printf("Ignoring database %s as it exists in the databases to ignore list\n", database)
			continue
		}

		if len(databasesToInclude) > 0 && !databasesToInclude[database] {
			log.Printf("Ignoring database %s as it does not exist in the databases to include list\n", database)
			continue
		}

		tables := discoverTables(database)
		for _, table := range tables {
			if tablesToIgnore[table] {
				log.Printf("Ignoring table %s as it exists in the tables to ignore list", table)
				continue
			}
			databaseTableList = append(databaseTableList, database+"."+table)
		}
	}
	return databaseTableList
}

func discoverDatabases() []string {
	return queryDBForList("SHOW DATABASES")
}

func discoverTables(database string) []string {
	return queryDBForList(fmt.Sprintf("SHOW TABLES FROM %s", database))
}

func countRows(databaseTable string) int {
	var count int
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM %s", databaseTable)
	err := db.DB.QueryRow(queryString).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func queryDBForList(query string) []string {
	var list []string
	rows, err := db.DB.Query(query)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item string
		if err := rows.Scan(&item); err != nil {
			log.Fatal(err)
		}
		list = append(list, item)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}
