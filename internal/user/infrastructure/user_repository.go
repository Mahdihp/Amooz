package infrastructure

import (
	"Amooz/database/postgres"
	"Amooz/internal/user/domain"
	"Amooz/pkg/config"
)

type IUserRepository interface {
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(user domain.User) error
	SoftDelete(user domain.User) error
	FindByID(userId int64) (domain.User, error)
}

type UserRepositoryImpl struct {
	postgres *postgres.PostgresDb
}

func NewUserRepository(cfg config.Config) IUserRepository {
	newPostgres := postgres.NewPostgres(cfg)
	return &UserRepositoryImpl{
		postgres: newPostgres,
	}
}

func (r *UserRepositoryImpl) Update(user domain.User) error {
	tx := r.postgres.Db.Model(&user).Updates(domain.User{DisplayName: user.DisplayName,
		Password: user.Password})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepositoryImpl) SoftDelete(user domain.User) error {
	tx := r.postgres.Db.Model(&user).
		Updates(domain.User{Deleted: true})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (r *UserRepositoryImpl) Delete(user domain.User) error {
	tx := r.postgres.Db.Where("ID = ?", user.ID).Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepositoryImpl) Create(user domain.User) error {
	tx := r.postgres.Db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepositoryImpl) FindByID(userId int64) (domain.User, error) {
	var user domain.User
	tx := r.postgres.Db.First(&user, "ID = ?", userId)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return user, nil
}
