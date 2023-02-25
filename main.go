package main

import (
	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/routes"
)

func main() {
	app := echo.New()
	routes.Setup(app)

	app.Start(":4002")

}
