package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iago-f-s-e/pix-code-go/src/domain/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	_, dir, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(dir)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env == "test" {
		dsn = os.Getenv("DSN_TEST")

		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {

		dsn = os.Getenv("DSN")

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.Logger.LogMode(logger.Silent)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
