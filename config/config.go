package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "dave",
			Password: os.Getenv("DBPASS"),
			Name:     "test_db",
			Charset:  "utf8",
		},
	}
}
