package handlers

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App, h *Database) {
	api := app.Group("/api", HeaderAuth)
	api.Post("/login", h.LoginHandler)
	v1 := api.Group("/v1", h.AuthMiddleware)
	v1.Get("/", h.RootHandler)

}

func (h *Database) RootHandler(ctx *fiber.Ctx) error {

	return ctx.SendString("Authorized")

}
