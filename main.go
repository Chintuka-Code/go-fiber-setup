package main

import (
	"struct-validation/src/types"
)

func main() {
	app := types.NewApplication(":8000", "/api/v1")

	app.SetMiddleware()
	app.SetUpRoutes()
	app.SetValidator()
	app.StartListen()
}
