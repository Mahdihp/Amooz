package application

import (
	"Amooz/internal/user/domain"
	"Amooz/internal/user/infrastructure"
	"context"
)

// IUserService اینترفیس برای مخزن کتاب‌ها
type IUserService interface {
	Save(ctx context.Context, book *domain.User) error
	FindByID(ctx context.Context, id string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
}

// UserServiceImpl سرویس مدیریت کتاب‌ها
type UserServiceImpl struct {
	userRepository infrastructure.IUserRepository
}

func NewUserService(repo infrastructure.IUserRepository) IUserService {
	return UserServiceImpl{
		userRepository: repo,
	}
}

func (u UserServiceImpl) Save(ctx context.Context, book *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) FindByID(ctx context.Context, id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	all, err := u.userRepository.FindAll()
	if err != nil {
		return []domain.User{}, err
	}
	return all, nil
}
