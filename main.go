package main

import (
	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/routes"
)

func main() {
	app := echo.New()
	routes.Home(app)
	routes.Setup(app)

	app.Logger.Fatal(app.Start(":4000"))
}
