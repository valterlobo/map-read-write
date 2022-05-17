package repositorio

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func CreateDatabase() string {

	dataSourceName := newDatabaseName(".", "person")
	file, err := os.Create(dataSourceName)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", dataSourceName)
	defer sqliteDatabase.Close() // Defer Closing the database
	createScript(sqliteDatabase) // Create Database

	return dataSourceName
}

func createScript(database *sql.DB) {

	streamScrypts, errorFile := ioutil.ReadFile("./sql/script.sql")

	strScript := string(streamScrypts)
	if errorFile != nil {
		panic(errorFile)
	}

	requests := strings.Split(strScript, ";")

	for _, request := range requests {
		_, err := database.Exec(request)
		if err != nil {
			panic(err)
		}
	}

}
func newDatabaseName(path string, dbName string) string {

	currentTime := time.Now()
	date := currentTime.Format("01-02-2006_150405")
	filePath := path + "/" + dbName + "_" + string(date) + ".db"
	return filePath
}

func GetConnection(dataSourceName string) *sql.DB {

	database, errorDB := sql.Open("sqlite3", dataSourceName)
	if errorDB != nil {
		panic(errorDB)
	}
	return database
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

