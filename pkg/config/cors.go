package config

// CORSConfig menyimpan konfigurasi CORS
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

// GetCORSConfig mengembalikan konfigurasi CORS
func GetCORSConfig() *CORSConfig {
	return &CORSConfig{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           86400, // 24 jam dalam detik
	}
}
