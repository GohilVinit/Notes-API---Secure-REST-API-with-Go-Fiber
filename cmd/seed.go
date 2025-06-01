package main

import (
	"fmt"
	"log"
	"notes-api/models"
	"notes-api/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	utils.ConnectDB()

	// Seed users
	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: hashPassword("password123"),
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: hashPassword("password123"),
		},
		{
			Name:     "Bob Johnson",
			Email:    "bob@example.com",
			Password: hashPassword("password123"),
		},
	}

	for _, user := range users {
		if err := utils.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to create user %s: %v", user.Email, err)
			continue
		}
		fmt.Printf("Created user: %s\n", user.Email)

		// Create sample notes for each user
		notes := []models.Note{
			{
				Title:   "Welcome Note",
				Content: "This is your first note! You can create, edit, and delete notes.",
				UserID:  user.ID,
			},
			{
				Title:   "Meeting Notes",
				Content: "Project discussion:\n- Implement authentication\n- Add note management\n- Deploy to production",
				UserID:  user.ID,
			},
			{
				Title:   "Shopping List",
				Content: "Groceries:\n- Milk\n- Bread\n- Eggs\n- Fruits",
				UserID:  user.ID,
			},
		}

		for _, note := range notes {
			if err := utils.DB.Create(&note).Error; err != nil {
				log.Printf("Failed to create note for user %s: %v", user.Email, err)
				continue
			}
		}
	}

	fmt.Println("Seeding completed successfully!")
	fmt.Println("Test users created:")
	fmt.Println("- john@example.com (password: password123)")
	fmt.Println("- jane@example.com (password: password123)")
	fmt.Println("- bob@example.com (password: password123)")
}

func hashPassword(password string) string {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}
	return hashed
}
