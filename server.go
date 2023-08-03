package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	control "go.api.puro.mysql/controller/mssql"
)

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
	//e.GET("/persons", getPersons)
	//e.GET("/persons/:id", getPerson)
	//e.POST("/persons", CreatePerson)
	//e.PUT("/persons/:id", UpdatePerson)
	//e.DELETE("/persons/:id", DeletePerson)

	e.GET("/employees", control.AllBeers)

	e.Logger.Fatal(e.Start(":1323"))

}
