package database

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"upper.io/db.v3/lib/sqlbuilder"

	homedir "github.com/mitchellh/go-homedir"
	"upper.io/db.v3/sqlite"
)

var settings = sqlite.ConnectionURL{
	Database: filepath.Join(getDBDir(), "cvpm-database.db"), // Path to a sqlite3 database file.
}

func runSQL(sql string, database *sql.DB) {
	statement, err := database.Prepare(sql)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func initDatabase() {
	database, err := sql.Open("sqlite3", filepath.Join(getDBDir(), "cvpm-database.db"))
	if err != nil {
		panic(err)
	}
	datasetSQLString := "CREATE TABLE IF NOT EXISTS dataset (id INTEGER PRIMARY KEY, Name TEXT, Desc TEXT, Tags TEXT, Files TEXT, Link TEXT)"
	requestSQLString := "CREATE TABLE IF NOT EXISTS request (id INTEGER PRIMARY KEY, Ip TEXT, Vendor TEXT, Package TEXT, Solver TEXT, Ray TEXT, Token Text, Timestamp Text)"
	environmentSQLString := "CREATE TABLE IF NOT EXISTS environment (id INTEGER PRIMARY KEY, Key TEXT, Value TEXT, Vendor TEXT, PackageName TEXT)"
	runSQL(datasetSQLString, database)
	runSQL(requestSQLString, database)
	runSQL(environmentSQLString, database)
	database.Close()
}

func getDBDir() string {
	homePath, _ := homedir.Dir()
	cvpmPath := filepath.Join(homePath, "cvpm")
	dbPath := filepath.Join(cvpmPath, "database")
	return dbPath
}

func GetDBInstance() sqlbuilder.Database {
	initDatabase()
	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}
	return sess
}

func CloseDB(sess sqlbuilder.Database) {
	defer sess.Close()
}
