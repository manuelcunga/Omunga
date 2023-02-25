package main

import (
	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/routes"
)

func main() {
	// database.Connection()
	app := echo.New()
	routes.Setup(app)

	app.Start(":4002")
	// if err := app.Start(":4002"); err != nil {
	// 	log.Fatalln("Failed to start server:", err)
	// }
}
