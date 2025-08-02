package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppEnv   = GetEnv("APP_ENV")
	APP_KEY  = GetEnv("APP_KEY")
	APP_NAME = GetEnv("APP_NAME")
	DBDriver = GetEnv("DB_DRIVER", "postgres")
	DBName   = GetEnv("DB_NAME", "local")
	DBHost   = GetEnv("DB_HOST", "localhost")
	DBPort   = GetEnv("DB_PORT", "5432")
	DBUser   = GetEnv("DB_USER", "root")
	DBPass   = GetEnv("DB_PASS", "")
	SSLMode  = GetEnv("SSL_MODE", "disable")
	// MONGOHost  = GetEnv("MONGO_HOST")
	// MONGOPort  = GetEnv("MONGO_PORT")
	// MONGODBDEV = MONGO_DB_DEV
	APPPort = GetEnv("APP_PORT")

	ProdUrl  = GetEnv("PROVIDER_PROD_URL")
	DevUrl   = GetEnv("PROVIDER_DEV_URL")
	LocalUrl = GetEnv("PROVIDER_LOCAL_URL")

	TrxPaymentPending = GetEnv("TRX_PAYMENT_PENDING")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != EMPTY_VALUE {
		log.Println(key, os.Getenv(key))
		return os.Getenv(key)
	} else {
		if len(value) > EMPTY_VALUE_INT {
			log.Println(key, value)
			return value[EMPTY_VALUE_INT]
		}
		return ""
	}
}
