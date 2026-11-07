package config

import (
	"fmt"
	"log"
	"time"

	"e-piket-be/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *Config) *gorm.DB {
	// DSN format: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	gormLogLevel := logger.Info
	if cfg.AppEnv == "production" {
		gormLogLevel = logger.Error
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke MySQL database: %v", err)
	}

	// Konfigurasi connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Berhasil terhubung ke database MySQL!")

	// Auto Migration untuk membuat/memperbarui struktur tabel otomatis
	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Printf("Peringatan: Gagal menjalankan auto-migrate: %v", err)
	} else {
		log.Println("Database AutoMigrate berhasil dijalankan.")
	}

	DB = db
	return db
}
