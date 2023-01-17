package main

import (
	"better/app"
	_ "better/docs"
)

// @title Better - The better backend
// @version 0.0.1
// @contact.name Cyteon

// @host localhost:8080
// @BasePath /_/
func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
