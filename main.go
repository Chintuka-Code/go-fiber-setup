package main

import (
	// _ "struct-validation/docs"
	"struct-validation/src/types"
)

// @title Go Fiber Setup
// @version 3.0
// @description This is a sample Go Fiber Application.
// @host localhost:8000
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @securityDefinitions.oauth2 Bearer
// @tokenUrl http://localhost:8000/oauth/token
// @scopes read:items=Grants read access to items
// @scopes write:items=Grants write access to items
// @scheme bearer
// @Security Bearer
// @param Accept-Language header string true "Language" default(en-US)
func main() {
	app := types.NewApplication(":8000", "/api/v1")

	app.SetMiddleware()
	app.InitSwagger()
	app.SetUpRoutes()
	app.SetValidator()
	app.StartListen()
}
