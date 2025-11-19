package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config menyimpan konfigurasi aplikasi
type Config struct {
	App AppConfig `envPrefix:"APP_"`
	DB  DBConfig  `envPrefix:"POSTGRES_"`
}

var Cfg Config

// AppConfig menyimpan konfigurasi aplikasi
type AppConfig struct {
	Port string `env:"PORT" envDefault:"3000"`
}

// DBConfig menyimpan konfigurasi database
type DBConfig struct {
	Host               string `env:"HOST" envDefault:"localhost"`
	Port               string `env:"PORT" envDefault:"5432"`
	User               string `env:"USER" envDefault:"postgres"`
	Password           string `env:"PASSWORD" envDefault:""`
	DBName             string `env:"DBNAME" envDefault:"certificate-api"`
	MaxOpenConnections int    `env:"MAX_OPEN_CONNECTIONS" envDefault:"5"`
	MaxConnLifetime    int    `env:"MAX_CONN_LIFETIME" envDefault:"10"`
	MaxIdleLifetime    int    `env:"MAX_IDLE_LIFETIME" envDefault:"5"`
}

// NewAppConfig membaca environment variable dan menginisialisasi konfigurasi aplikasi
func NewAppConfig() (*Config, error) {
	_ = godotenv.Load()
	if err := env.Parse(&Cfg); err != nil {
		return nil, fmt.Errorf("[Config-NewAppConfig] gagal memuat konfigurasi: %w", err)
	}
	return &Cfg, nil
}

// GetConfig mengembalikan instance konfigurasi aplikasi
func GetConfig() *Config {
	return &Cfg
}
