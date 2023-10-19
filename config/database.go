package config

import (
	"fmt"
	"go-jwt-api/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB{
// 	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	helper.ErrorPanic(err)
	fmt.Println("🚀 Connected Successful")

	return db
}