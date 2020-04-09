package mysqlDB 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Initialize(username string, password string, hostname string, port string,dbname string) *sql.DB {
	connectionString := username +":"+ password + "@tcp(" + hostname + ":" + port + ")/"+dbname
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
        return db;   

}
