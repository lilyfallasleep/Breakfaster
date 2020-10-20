package config

import (
	"os"
	"strconv"
	"time"
)

// GetEnvWithDefault is a helper function for specifying a default env value
func GetEnvWithDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Config is a type for general configuration
type Config struct {
	DBconfig               *DBConfig
	DatabaseDsn            string
	Port                   string
	ChannelSecret          string
	AccessToken            string
	BotVersion             string
	OrderPageURI           string
	GinMode                string
	LogPath                string
	Logger                 *Logger
	DefaultCacheExpiration time.Duration
	CleanCacheInterval     time.Duration
	ClovaSecretKey         string
	ClovaBuilderURL        string
}

// NewConfig is a factory for Config instance
func NewConfig() (*Config, error) {
	maxDBIdleConns, err := strconv.Atoi(GetEnvWithDefault("MAX_DB_IDLE_CONNS", "10"))
	if err != nil {
		return &Config{}, err
	}
	maxDBOpenConns, err := strconv.Atoi(GetEnvWithDefault("MAX_DB_OPEN_CONNS", "100"))
	if err != nil {
		return &Config{}, err
	}
	defaultCacheExpiration, err := strconv.ParseInt(GetEnvWithDefault("DEFAULT_CACHE_EXPIRATION", "300"), 10, 64)
	if err != nil {
		return &Config{}, err
	}
	cleanCacheInterval, err := strconv.ParseInt(GetEnvWithDefault("CLEAN_CACHE_INTERVAL", "600"), 10, 64)
	if err != nil {
		return &Config{}, err
	}

	ginMode := GetEnvWithDefault("GIN_MODE", "debug")
	logPath := GetEnvWithDefault("LOG_PATH", "server.log")
	logger, err := getLogger(ginMode, logPath)
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		DBconfig: &DBConfig{
			MaxIdleConns: maxDBIdleConns,
			MaxOpenConns: maxDBOpenConns,
		},
		DatabaseDsn:            os.Getenv("DB_DSN"),
		Port:                   GetEnvWithDefault("PORT", "80"),
		ChannelSecret:          os.Getenv("CHANNEL_SECRET"),
		AccessToken:            os.Getenv("ACCESS_TOKEN"),
		BotVersion:             GetEnvWithDefault("BOT_VERSION", "v1"),
		OrderPageURI:           os.Getenv("ORDER_PAGE_URI"),
		GinMode:                ginMode,
		LogPath:                logPath,
		Logger:                 logger,
		DefaultCacheExpiration: time.Duration(defaultCacheExpiration) * time.Second,
		CleanCacheInterval:     time.Duration(cleanCacheInterval) * time.Second,
		ClovaSecretKey:         os.Getenv("CLOVA_SECRET_KEY"),
		ClovaBuilderURL:        os.Getenv("CLOVA_BUILDER_URL"),
	}, nil
}
