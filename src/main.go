package main

import "github.com/aljrubior/amc-ui-rest-facade/application"

func main() {
	app := application.Application{}

	app.Run(":8080")
}
