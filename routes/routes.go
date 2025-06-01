package routes

import (
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected notes routes
	notes := api.Group("/notes")
	notes.Use(middleware.JWTMiddleware)
	notes.Post("/", handlers.CreateNote)
	notes.Get("/", handlers.GetNotes)
	notes.Get("/:id", handlers.GetNote)
	notes.Put("/:id", handlers.UpdateNote)
	notes.Delete("/:id", handlers.DeleteNote)
}
