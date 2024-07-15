package config

import "os"

var Port = "8080"

var Database = struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Username: os.Getenv("DB_USERNAME"),
	Password: os.Getenv("DB_PASSWORD"),
	Database: os.Getenv("DB_DATABASE"),
}
