package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	POSTGRES = "postgres"
	NEON     = "neon"
	LOCAL    = "local"
)

type StorageInput struct {
	DB         *gorm.DB
	BucketName string
}

func NewPostgresDB(ctx context.Context) *gorm.DB {
	newLogger := logger.New(
		log.New(nil, "", 0),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	dsn := "host=" + config.DB_POSTGRES_CONTAINER.DB_HOST + " user=" + config.DB_POSTGRES_CONTAINER.DB_USER + " password=" + config.DB_POSTGRES_CONTAINER.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_CONTAINER.DB_NAME + " port=" + config.DB_POSTGRES_CONTAINER.DB_PORT + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil
	}
	return db
}

func NewPostgresDBLocal(ctx context.Context) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	dsn := "host=" + config.DB_POSTGRES_LOCAL.DB_HOST + " user=" + config.DB_POSTGRES_LOCAL.DB_USER + " password=" + config.DB_POSTGRES_LOCAL.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_LOCAL.DB_NAME + " port=" + config.DB_POSTGRES_LOCAL.DB_PORT + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil
	}
	return db
}

func NeonConnection(ctx context.Context) *gorm.DB {
	newLogger := logger.New(
		log.New(nil, "", 0),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	dsn := "postgresql://" + config.DB_NEON.DB_USER + ":" + config.DB_NEON.DB_PASSWORD + "@" + config.DB_NEON.DB_HOST + "/" + config.DB_NEON.DB_NAME + "?sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil
	}
	return db
}

func SetupDatabaseConnection(ctx context.Context, SGBD string) (*gorm.DB, *sql.DB) {
	var db *gorm.DB

	switch SGBD {
	case POSTGRES:
		db = NewPostgresDB(ctx)
	case NEON:
		db = NeonConnection(ctx)
	case LOCAL:
		db = NewPostgresDBLocal(ctx)
	default:
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	return db, sqlDB
}

func CheckConnection(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil {
		return false
	}

	if err := sqlDB.Ping(); err != nil {
		return false
	}

	return true
}

func Shutdown(ctx context.Context, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	if err := sqlDB.Close(); err != nil {

	}
}
