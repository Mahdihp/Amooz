package infrastructure

import (
	"Amooz/database/postgres"
	"Amooz/internal/user/domain"
	"Amooz/pkg/config"
	"context"
	"errors"
)

// UserRepository مخزن کتاب‌ها است که داده‌ها را ذخیره می‌کند
type UserRepository struct {
	postgres *postgres.PostgresDb
}

// NewUserRepository سازنده برای UserRepository
func NewUserRepository(cfg config.Config) UserRepository {
	//newPostgres := postgres.NewPostgres(cfg, []interface{}{&domain.User{}})

	return UserRepository{
		postgres: nil,
	}
}

// Save کتاب را ذخیره می‌کند
func (r *UserRepository) Save(ctx context.Context, book *domain.User) error {
	r.postgres.Db.Create(book)
	return errors.New("not implemented")
}

// FindByID کتابی را با شناسه پیدا می‌کند
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, errors.New("not implemented")
}
