package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/A7med-Mido/educationalPlatformApi/src/config"       // Assuming module is edu-platform; adjust if needed
	"github.com/A7med-Mido/educationalPlatformApi/src/database"
	"github.com/A7med-Mido/educationalPlatformApi/src/handlers"
	"github.com/A7med-Mido/educationalPlatformApi/src/middleware"
)

func main() {
	app := fiber.New()

	// Load environment variables and connect to the database
	config.LoadEnv()
	database.ConnectDB()

	// Public routes for authentication
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.RegisterStudent)

	// Protected dashboard routes for teacher (admin)
	admin := app.Group("/dashboard", middleware.AuthMiddleware("teacher"))
	admin.Post("/courses", handlers.CreateCourse)
	admin.Get("/courses", handlers.GetCourses)
	admin.Put("/courses/:id", handlers.UpdateCourse)
	admin.Delete("/courses/:id", handlers.DeleteCourse)
	admin.Post("/courses/:id/thumbnail", handlers.UploadThumbnail)
	admin.Get("/reports", handlers.CreateCourse)

	// Protected API routes for students
	student := app.Group("/api", middleware.AuthMiddleware("student"))
	student.Get("/courses", handlers.GetAvailableCourses)
	student.Post("/subscribe/:course_id", handlers.SubscribeToCourse)

	// Start the server
	app.Listen(":3000")
}