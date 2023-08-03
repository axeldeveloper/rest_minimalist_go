package models

import (
	"fmt"
	"time"

	"go.api.puro.mysql/db"

	_ "github.com/denisenkom/go-mssqldb"
)

type Cerveja struct {
	Id        int       `json:id`
	Score     string    `json:score`
	Winner    string    `json:winner`
	Bar       string    `json:bar`
	CreatedOn time.Time `json:created_on`
	IsFinal   bool      `json:is_final`
}

// ReadEmployees read all employees
func (c Cerveja) GetBeers() []Cerveja {
	db := db.CreateCon()

	tsql := " SELECT * FROM Cerveja "
	//tsql := fmt.Sprintf(q)
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil
	}
	defer rows.Close()

	var cervejas []Cerveja

	count := 0
	for rows.Next() {
		var cerveja Cerveja
		err := rows.Scan(&cerveja.Id, &cerveja.Score, &cerveja.Winner, &cerveja.Bar, &cerveja.CreatedOn, &cerveja.IsFinal)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil
		}
		fmt.Printf("ID: %d, Score: %s, Winner: %s\n", cerveja.Id, cerveja.Score, cerveja.Winner)
		cervejas = append(cervejas, cerveja)
		count++
	}
	return cervejas
}

/*

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
