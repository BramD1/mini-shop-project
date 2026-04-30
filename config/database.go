package config

import (
    "os"
    "mini-shop/domain"
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

    err = db.AutoMigrate(
        &domain.User{},
        &domain.Toko{},
        &domain.Alamat{},
        &domain.Category{},
        &domain.Produk{},
        &domain.FotoProduk{},
        &domain.LogProduk{},
        &domain.Trx{},
        &domain.DetailTrx{},
    )
	if err != nil {
        return nil, err
    }

    return db, nil

}