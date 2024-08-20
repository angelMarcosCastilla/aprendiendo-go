package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() *sql.DB {
	server := "127.0.0.1"
	port := 1433
	database := "todoListApp"

	connString := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes;",
		server, port, database)

	conexion, err := sql.Open("mssql", connString)

	if err != nil {
		panic(err)
	}

	return conexion
}
