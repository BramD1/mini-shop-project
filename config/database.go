package config

import (
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func ConnectDB() (*gorm.DB, error) {
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
           "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
           os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}