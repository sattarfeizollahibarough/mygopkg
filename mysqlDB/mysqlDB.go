package mysqlDB

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// CreateTable will create a table
func CreateTable(db *sql.DB, tablename string, columns map[string]string) {
	query := "CREATE TABLE IF NOT EXISTS " + tablename + "("
	for colname, coltype := range columns {
		query += colname + " " + coltype + ","
	}
	lastindx := strings.LastIndex(query, ",")
	query = query[:lastindx] + strings.Replace(query[lastindx:], ",", ")", 1)
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
}

// ExecuteQuery will execute insert and update statements
func ExecuteQuery(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
}

// Initialize will initialize database connection
func Initialize(username string, password string, hostname string, port string, dbname string) *sql.DB {
	connectionString := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname + "?autocommit=true"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(120)
	db.SetMaxOpenConns(120)
	return db
}

// SelectQuery will execute select statements
func SelectQuery(db *sql.DB, query string) *sql.Rows {
	res, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	return res
}
