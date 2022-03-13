package main

import "github.com/aljrubior/amc-ui-rest-facade/application"

func main() {
	app := application.NewApp()

	app.Run(":8080")
}
