package main

import "github.com/aljrubior/go-facade/application"

func main() {
	app := application.NewApp()

	app.Run(":8080")
}
