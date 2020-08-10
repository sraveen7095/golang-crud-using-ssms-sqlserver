package config

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func Connstr() (db *sql.DB) {

	dbDriver := "mssql"
	dbUser := "DESKTOP-1VA2HU8\\srave"
	//dbPass := "your_password"
	dbName := "portfolio"
	db, err := sql.Open(dbDriver, dbUser+":@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
