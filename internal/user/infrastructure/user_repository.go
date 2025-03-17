package infrastructure

import (
	"Amooz/database/postgres"
	"Amooz/internal/user/domain"
	"Amooz/pkg/config"
	"time"
)

type IUserRepository interface {
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(user domain.User) error
	SoftDelete(user domain.User) error
	FindByID(userId int64) (domain.User, error)
}

type UserRepository struct {
	postgres *postgres.PostgresDb
}

func NewUserRepository(cfg config.Config) IUserRepository {
	//newPostgres := postgres.NewPostgres(cfg, []interface{}{&domain.User{}})

	return &UserRepository{
		postgres: nil,
	}
}

func (r *UserRepository) Update(user domain.User) error {
	tx := r.postgres.Db.Model(&user).Updates(domain.User{DisplayName: user.DisplayName,
		Password: user.Password, UpdatedAt: time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepository) SoftDelete(user domain.User) error {
	tx := r.postgres.Db.Model(&user).
		Updates(domain.User{Deleted: true, UpdatedAt: time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (r *UserRepository) Delete(user domain.User) error {
	tx := r.postgres.Db.Where("ID = ?", user.ID).Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepository) Create(user domain.User) error {
	tx := r.postgres.Db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepository) FindByID(userId int64) (domain.User, error) {
	var user domain.User
	tx := r.postgres.Db.First(&user, "ID = ?", userId)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return user, nil
}
