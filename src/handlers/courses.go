package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/A7med-Mido/educationalPlatformApi/src/services"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"github.com/A7med-Mido/educationalPlatformApi/src/repositories"
	"path/filepath"
	"github.com/google/uuid"

)


var courseService = services.NewCourseService()
var courseRepo = repositories.NewCourseRepo()

func CreateCourse(c fiber.Ctx) error {
	var course models.Course
	if err := c.Bind().Body(&course); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	teacherID := c.Locals("userID").(uuid.UUID)
	if err := courseService.Create(&course, teacherID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create course"})
	}
	return c.JSON(course)
}

func GetCourses(c fiber.Ctx) error {
	teacherID := c.Locals("userID").(uuid.UUID)
	courses, err := courseService.GetAll(teacherID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get courses"})
	}
	return c.JSON(courses)
}

func UpdateCourse(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var updates models.Course
	if err := c.Bind().Body(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := courseService.Update(id, &updates); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update course"})
	}
	return c.JSON(fiber.Map{"message": "Course updated"})
}

func DeleteCourse(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := courseService.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete course"})
	}
	return c.JSON(fiber.Map{"message": "Course deleted"})
}

func UploadThumbnail(c fiber.Ctx) error {
	file, err := c.FormFile("thumbnail")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}
	ext := filepath.Ext(file.Filename)
	path := "./uploads/" + uuid.New().String() + ext
	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	course, err := courseRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Course not found"})
	}
	course.Thumbnail = path
	if err := courseRepo.Update(course); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update thumbnail"})
	}
	return c.JSON(fiber.Map{"thumbnail": path})
}