package db

import (
	"fmt"
	"time"

	"gitlab.com/genuicomm/api/domain"
	"gitlab.com/genuicomm/api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah wrapper untuk gorm.DB
type DB struct {
	DB *gorm.DB
}

// NewConnection membuat koneksi database baru
func NewConnection(cfg config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Set koneksi pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleLifetime)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Minute)

	// Auto migrate semua model
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("error auto migrating models: %v", err)
	}

	return &DB{DB: db}, nil
}

// autoMigrate melakukan migrasi otomatis untuk semua model
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.Article{},
		&domain.Certificate{},
		&domain.Section{},
		&domain.File{},
	)
}

// Reconnect mencoba untuk membuat koneksi baru ke database
func Reconnect(db *DB, cfg config.DBConfig) error {
	// Tutup koneksi lama jika ada
	if db.DB != nil {
		sqlDB, err := db.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	// Buat koneksi baru
	newDB, err := NewConnection(cfg)
	if err != nil {
		return err
	}

	// Update koneksi database
	db.DB = newDB.DB
	return nil
}

// Close menutup koneksi database
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
