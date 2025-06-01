package handlers

import (
	"notes-api/models"
	"notes-api/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req models.CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if strings.TrimSpace(req.Title) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	note := models.Note{
		Title:   strings.TrimSpace(req.Title),
		Content: req.Content,
		UserID:  userID,
	}

	if err := utils.DB.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create note",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// Pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Build query
	query := utils.DB.Where("user_id = ?", userID)

	// Add search functionality
	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	var total int64
	query.Model(&models.Note{}).Count(&total)

	// Get notes with pagination
	var notes []models.Note
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notes",
		})
	}

	return c.JSON(models.NotesResponse{
		Notes: notes,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

func GetNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	var req models.UpdateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update fields if provided
	if req.Title != "" {
		note.Title = strings.TrimSpace(req.Title)
	}
	if req.Content != "" {
		note.Content = req.Content
	}

	if err := utils.DB.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	result := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).Delete(&models.Note{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete note",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}
