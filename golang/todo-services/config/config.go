package config

import (
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	DBUser 					string
	DBPassword			string
	DBName					string
	DBPort					string
	DBHost					string
	SECRETE_TOKEN		string
}

func LoadConfig() Config {
	return Config{
		DBUser: 				os.Getenv("DB_USER"),
		DBPassword: 		os.Getenv("DB_PASSWORD"),
		DBName: 				os.Getenv("DB_NAME"),
		DBPort:	 				os.Getenv("DB_PORT"),
		DBHost: 				os.Getenv("DB_HOST"),

		SECRETE_TOKEN: 	os.Getenv("SERCRET_TOKEN"),
	}
}

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)
	// Define Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)