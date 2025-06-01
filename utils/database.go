package utils

import (
    "fmt"
    "log"
    "os"
    "notes-api/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    config := &gorm.Config{}
    if os.Getenv("ENV") == "development" {
        config.Logger = logger.Default.LogMode(logger.Info)
    }

    DB, err = gorm.Open(mysql.Open(dsn), config)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connected successfully")
    
    // Auto migrate the schema
    err = DB.AutoMigrate(&models.User{}, &models.Note{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    
    log.Println("Database migration completed")
}