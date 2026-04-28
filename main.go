package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "mini-shop/config"
)

func main() {
    // 1. Load .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 2. Connect to database
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 3. Just to confirm connection works for now
    log.Println("Database connected!", db)

    // 4. Start server (we'll add router later)
    port := os.Getenv("SERVER_PORT")
    log.Println("Server running on port", port)
}