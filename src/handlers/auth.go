package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"github.com/A7med-Mido/educationalPlatformApi/src/utils"
	"github.com/A7med-Mido/educationalPlatformApi/src/services"
	"golang.org/x/crypto/bcrypt"
)

var userService = services.NewUserService()

func Login(c fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := userService.GetByEmail(input.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func RegisterStudent(c fiber.Ctx) error {
	var user models.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := userService.RegisterStudent(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}
	return c.JSON(fiber.Map{"message": "Student registered"})
}