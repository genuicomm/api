package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/genuicomm/api/db"
	"github.com/genuicomm/api/modules"
	"github.com/genuicomm/api/pkg/config"
)

func main() {
	// Load konfigurasi aplikasi
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal("Gagal memuat konfigurasi aplikasi:", err)
	}

	// Koneksi ke database dengan retry mechanism
	database, err := connectWithRetry(cfg.DB)
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer database.Close()

	// Inisialisasi router Gin dengan middleware CORS
	router := modules.InitModules(cfg, database)

	// Konfigurasi server HTTP
	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: router,
	}

	// Channel untuk menangkap sinyal shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Jalankan server dalam goroutine
	go func() {
		log.Println("Server berjalan di port", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server gagal berjalan:", err)
		}
	}()

	// Tunggu sinyal shutdown
	<-quit
	log.Println("Shutting down server...")

	// Buat context dengan timeout untuk graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Coba shutdown server dengan graceful
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// connectWithRetry mencoba koneksi ke database dengan mekanisme retry
func connectWithRetry(dbConfig config.DBConfig) (*db.DB, error) {
	var database *db.DB
	var err error
	maxRetries := 5
	retryDelay := time.Second * 5

	for i := 0; i < maxRetries; i++ {
		database, err = db.NewConnection(dbConfig)
		if err == nil {
			// Jika koneksi berhasil, jalankan goroutine untuk monitoring koneksi
			go monitorDBConnection(database, dbConfig)
			return database, nil
		}

		log.Printf("Gagal terhubung ke database (percobaan %d/%d): %v\n", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryDelay)
			retryDelay *= 2 // Exponential backoff
		}
	}

	return nil, fmt.Errorf("gagal terhubung ke database setelah %d percobaan: %v", maxRetries, err)
}

// monitorDBConnection memonitor koneksi database dan mencoba reconnect jika terputus
func monitorDBConnection(dbConnection *db.DB, dbConfig config.DBConfig) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for range ticker.C {
		sqlDB, err := dbConnection.DB.DB()
		if err != nil {
			log.Printf("Gagal mengambil koneksi database: %v\n", err)
			continue
		}

		if err := sqlDB.Ping(); err != nil {
			log.Println("Koneksi database terputus, mencoba reconnect...")

			// Coba koneksi ulang menggunakan fungsi Reconnect dari package db
			if err := db.Reconnect(dbConnection, dbConfig); err != nil {
				log.Printf("Gagal reconnect ke database: %v\n", err)
				continue
			}

			log.Println("Berhasil reconnect ke database")
		}
	}
}
