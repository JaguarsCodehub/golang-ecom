package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string // host:port
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	cfg := Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "TdmxJIuUTCjRrIWViEZxmkeIhabqvMeG"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}

	// Support full DATABASE_URL (e.g. mysql://user:pass@host:port/dbname)
	if dbURL := getEnv("DATABASE_URL", ""); dbURL != "" {
		if u, err := url.Parse(dbURL); err == nil {
			if u.User != nil {
				cfg.DBUser = u.User.Username()
				if p, ok := u.User.Password(); ok {
					cfg.DBPassword = p
				}
			}

			// Host contains host:port
			if u.Host != "" {
				cfg.DBAddress = u.Host
			}

			// Path is /dbname
			if u.Path != "" {
				cfg.DBName = strings.TrimPrefix(u.Path, "/")
			}
		}
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
