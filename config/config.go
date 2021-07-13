package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

// Config represents service configuration for dp-collection-api
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	DefaultMaxLimit            int           `envconfig:"DEFAULT_MAXIMUM_LIMIT"`
	DefaultLimit               int           `envconfig:"DEFAULT_LIMIT"`
	DefaultOffset              int           `envconfig:"DEFAULT_OFFSET"`
	MongoConfig                MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr          string `envconfig:"MONGODB_BIND_ADDR"           json:"-"` // This line contains sensitive data and the json:"-" tells the json marshaller to skip serialising it.
	ContentDatabase   string `envconfig:"MONGODB_CONTENT_DATABASE"`
	ContentCollection string `envconfig:"MONGODB_CONTENT_COLLECTION"`
	Username          string `envconfig:"MONGODB_USERNAME"    json:"-"`
	Password          string `envconfig:"MONGODB_PASSWORD"    json:"-"`
	IsSSL             bool   `envconfig:"MONGODB_IS_SSL"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                   "localhost:26400",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		DefaultMaxLimit:            1000,
		DefaultLimit:               20,
		DefaultOffset:              0,
		MongoConfig: MongoConfig{
			BindAddr:          "localhost:27017",
			ContentDatabase:   "content",
			ContentCollection: "content",
			Username:          "",
			Password:          "",
			IsSSL:             false,
		},
	}

	return cfg, envconfig.Process("", cfg)
}
