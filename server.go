package main
 
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type PersonScan struct {
	id, name, salary, age string
}

type Person struct {
	Id     string `json: "id"`
	Name   string `json: "name"`
	Salary string `json: "salary"`
	Age    string `json: "age"`
}

type Persons struct {
	Persons []Person `json:"person"`
}

/*Create mysql connection*/
func CreateCon() *sql.DB {
	//db, err := sql.Open("mysql", "db_user:db_password@tcp(SERVER_IP:PORT)/database_name")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/my_db_test")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
func getPersons(c echo.Context)   error {
	var employee PersonScan
	sql := "SELECT id, name, salary, age FROM person ORDER BY id DESC"
	db := CreateCon()	
	selDB, err := db.Query(sql)
	
	if err != nil {
		fmt.Println(err)
	}
	emp := Person{}
	res := []Person{}	
	for selDB.Next() {
		err = selDB.Scan(&employee.id, &employee.name, &employee.salary, &employee.age)
		if err != nil {
			panic(err.Error())
		}
		emp.Id     = employee.id
		emp.Name   = employee.name
		emp.Salary = employee.salary			
		emp.Age    = employee.age	
		res = append(res, emp)
		//Persons{emp}
	}
	response := res
	return c.JSON(http.StatusOK, response)
}
func getPerson(c echo.Context)    error {
	var employee PersonScan
	sql := "SELECT id, name, salary, age FROM person WHERE id = ?"
	db  := CreateCon()
	err := db.QueryRow(sql, c.Param("id")).Scan(&employee.id, &employee.name, &employee.salary, &employee.age)
	
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound , "Erro geral : "+ err.Error())
	}
	
	response := Person{Id: employee.id, Name: employee.name, Salary: employee.salary, Age: employee.age}
	return c.JSON(http.StatusOK, response)

}
func CreatePerson(c echo.Context) error {
	emp := new(Person)
	if err := c.Bind(emp); err != nil {
		return err
	}
	db  := CreateCon()
	sql := "INSERT INTO person(name,  salary, age) VALUES( ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	
	if err != nil {
		fmt.Print(err.Error())
		return c.JSON(http.StatusNotFound , "Erro ao inserir o registro : "+ err.Error())
	}
	
	defer stmt.Close()
	result, err2 := stmt.Exec(emp.Name, emp.Salary, emp.Age)
	
	// Exit if we get an error
	if err2 != nil {
		//panic(err2)
		return c.JSON(http.StatusNotFound , "Panico : "+ err2.Error())
	}
	fmt.Println(result.LastInsertId())	
	return c.JSON(http.StatusCreated, emp)
}
func UpdatePerson(c echo.Context) error {
	emp := new(Person)
	if err := c.Bind(emp); err != nil {
		return err
	}

	id := c.Param("id")
	if id == "" {
		panic("error")
	}

	db  := CreateCon()
	sql := `UPDATE
				person
			SET
				name   = ?,
				salary = ?,
				age    = ?
			WHERE
				id = ?`

	stmt, err := db.Prepare(sql)
	
	if err != nil {
		fmt.Print(err.Error())
		return c.JSON(http.StatusNotFound , "Erro ao inserir o registro : "+ err.Error())
	}
	
	defer stmt.Close()
	result, err2 := stmt.Exec(emp.Name, emp.Salary, emp.Age, emp.Id)
	
	// Exit if we get an error
	if err2 != nil {
		//panic(err2)
		return c.JSON(http.StatusNotFound , "Panico : "+ err2.Error())
	}
	fmt.Println(result.LastInsertId())	
	return c.JSON(http.StatusAccepted, emp)
}
func DeletePerson(c echo.Context) error {
	db  := CreateCon()
	sql := "Delete FROM person Where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "Erro : "+ err.Error())
	}
	result, err2 := stmt.Exec(c.Param("id"))
	if err2 != nil {
		//panic(err2)
		return c.JSON(http.StatusNotFound , "Erro ao inserir o registro : "+ err2.Error())
	}
	fmt.Println(result.RowsAffected())
	return c.JSON(http.StatusOK, "Deleted") 
}
func Welcome(c echo.Context) error {
	//return c.JSON( http.StatusOK, "<strong>Hello, World!</strong>" )
	//return c.String(http.StatusOK, "Hello, World!")	
	return c.HTML(http.StatusOK, "<strong> High performance, minimalist Go web framework ⇨ http server started on [::]:1323 </strong>")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", Welcome)
	e.GET("/persons",        getPersons)
	e.GET("/persons/:id",    getPerson)
	e.POST("/persons",       CreatePerson)
	e.PUT("/persons/:id",    UpdatePerson)
	e.DELETE("/persons/:id", DeletePerson)
	e.Logger.Fatal(e.Start(":1323"))
}