package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

//const connStr = "Server=172.20.13.96;Database=DESENVOLVIMENTO;User Id=axel;Password=axel123;"

var (
	server   = "127.0.0.1"
	port     = 1433
	user     = "sa"
	password = "Dev1234!"
	database = "api_teste"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	fmt.Printf("Connected!\n")
	//defer conn.Close()
	return conn
}
