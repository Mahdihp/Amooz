package postgres

import (
	"Amooz/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type PostgresDb struct {
	Db *gorm.DB
}

func NewPostgres(cfg config.Config, tables []interface{}) *PostgresDb {
	strConn := fmt.Sprintf("host=%s user=y%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Pass, cfg.Postgres.DbName, cfg.Postgres.Port)
	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance!", err)
	}

	// تنظیمات اتصال به دیتابیس
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxConns) // حداکثر تعداد اتصالات باز
	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdle)  // حداکثر تعداد اتصالات بی‌کار
	sqlDB.SetConnMaxLifetime(time.Minute * 30)   // مدت زمان استفاده از یک اتصال قبل از بسته شدن

	// Migrate the schema
	db.AutoMigrate(&tables)

	return &PostgresDb{Db: db}
}
