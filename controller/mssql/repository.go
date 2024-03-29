package mssql

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo"
	"go.api.puro.mysql/models"
)

func AllBeers(c echo.Context) error {

	cervejas := &models.Cerveja{}
	lists := cervejas.GetBeers()

	return c.JSON(http.StatusOK, lists)
}

/*
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

func CreateEmployee(db *sql.DB, name string, location string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO TestSchema.Employees (Name, Location) VALUES ('%s','%s');",
		name, location)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func UpdateEmployee(db *sql.DB, name string, location string) (int64, error) {
	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = '%s' WHERE Name= '%s'",
		location, name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func DeleteEmployee(db *sql.DB, name string) (int64, error) {
	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name='%s';", name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}

func ExecuteAggregateStatement(db *sql.DB) {
	result, err := db.Prepare("SELECT SUM(Price) as sum FROM Table_with_5M_rows")
	if err != nil {
		fmt.Println("Error preparing query: " + err.Error())
	}

	row := result.QueryRow()
	var sum string
	err = row.Scan(&sum)
	fmt.Printf("Sum: %s\n", sum)
}
*/
