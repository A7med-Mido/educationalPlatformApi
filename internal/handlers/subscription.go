package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/internal/services"
	"github.com/A7med-Mido/educationalPlatformApi/internal/repositories"
)

var subscriptionService = services.NewSubscriptionService()
var courseRepoForSub = repositories.NewCourseRepo()

func GetAvailableCourses(c fiber.Ctx) error {
	// For students: Get all courses (since single-teacher platform)
	courses, err := courseRepoForSub.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get courses"})
	}
	return c.JSON(courses)
}

func SubscribeToCourse(c fiber.Ctx) error {
	idStr := c.Params("course_id")
	courseID, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	course, err := courseRepoForSub.FindByID(courseID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Course not found"})
	}
	studentID := c.Locals("userID").(uuid.UUID)
	sub, err := subscriptionService.Subscribe(studentID, courseID, course.IsFree)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to subscribe"})
	}
	if sub == nil {
		return c.JSON(fiber.Map{"message": "Payment required", "price": course.Price})
	}
	return c.JSON(sub)
}