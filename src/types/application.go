package types

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"struct-validation/src/localize"
	"struct-validation/src/middleware"
	"struct-validation/src/routes"
	"struct-validation/src/validation"
	customValidation "struct-validation/src/validation/custom-validation"
)

type Application struct {
	server    *fiber.App
	port      string
	Validator *validator.Validate
	baseRoute string
}

func NewApplication(port, baseRoute string) *Application {
	server := fiber.New()
	return &Application{
		server:    server,
		port:      port,
		baseRoute: baseRoute,
	}
}

func (app *Application) SetUpRoutes() {
	routes.InitRoutes(app.server, app.baseRoute)

}

func (app *Application) SetValidator() {
	app.Validator = validation.InitValidator()
	customValidation.RegisterCustomValidation(app.Validator)
}

func (app *Application) SetMiddleware() {
	app.server.Use(recover.New())
	app.server.Use(cors.New())
	app.server.Use(localize.NewLocalization())
	app.server.Use(middleware.RequestLogger)
	app.server.Use(middleware.GlobalErrorCatch)
	app.server.Use(middleware.ValidationErrorCatch)
}

func (app *Application) StartListen() {
	if err := app.server.Listen(app.port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (app *Application) InitSwagger() {
	// @Security Bearer
	app.server.Get("/docs/*", middleware.NewSwagger())
}
