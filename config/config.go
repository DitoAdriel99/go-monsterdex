package config

import (
	"os"
	"strconv"
)

type Cfg struct {
	ENV      string
	Port     int
	Database Database
	Redis    Redis
	GCS      GCS
	JWT      JWT
}

type Database struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	URL      string
}

type Redis struct {
	Host     string
	Password string
}

type GCS struct {
	AccountPath string
	ProjectID   string
	Storage     Storage
}

type Storage struct {
	Bucket string
	Prefix string
}

type JWT struct {
	Key     string
	Expired int
}

func PopulateConfigFromEnv() Cfg {
	var conf Cfg

	// Set default values or use os.Getenv() for each field
	conf.ENV = os.Getenv("ENVIRONMENT")
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// handle error (e.g., set a default port)
		port = 8080
	}
	conf.Port = port

	conf.Database.Driver = os.Getenv("DB_DRIVER")
	conf.Database.Username = os.Getenv("DB_USERNAME")
	conf.Database.Password = os.Getenv("DB_PASSWORD")
	conf.Database.Host = os.Getenv("DB_HOST")
	conf.Database.Port = os.Getenv("DB_PORT")
	conf.Database.URL = os.Getenv("DB_URL")

	conf.Redis.Host = os.Getenv("REDIS_HOST")
	conf.Redis.Password = os.Getenv("REDIS_PASSWORD")

	conf.GCS.AccountPath = os.Getenv("GCS_ACCOUNT_PATH")
	conf.GCS.ProjectID = os.Getenv("GCS_PROJECT_ID")
	conf.GCS.Storage.Bucket = os.Getenv("GCS_BUCKET")
	conf.GCS.Storage.Prefix = os.Getenv("GCS_PREFIX")

	conf.JWT.Key = os.Getenv("JWT_KEY")
	expiredStr := os.Getenv("JWT_EXPIRED")
	expired, err := strconv.Atoi(expiredStr)
	if err != nil {
		// handle error (e.g., set a default expiration)
		expired = 3600 // default expiration in seconds
	}
	conf.JWT.Expired = expired

	return conf
}
