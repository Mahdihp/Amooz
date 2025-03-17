package postgres

import (
	"Amooz/internal/user/domain"
	"Amooz/pkg/common"
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

func NewPostgres(cfg config.Config) *PostgresDb {
	strConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Pass, cfg.Postgres.DbName, cfg.Postgres.Port)
	//db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  strConn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

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
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.Post{})

	seedData(db)

	return &PostgresDb{Db: db}
}
func seedData(db *gorm.DB) {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count == 0 {
		role := domain.Role{
			Name: common.KEY_ROLE_ADMIN,
		}
		db.Create(&domain.User{
			Username: "mahdi",
			Password: "123456",
			Deleted:  false,
			Roles:    []domain.Role{role},
		})
	}
}
