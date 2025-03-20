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

// UserService سرویس مدیریت کتاب‌ها
type UserService struct {
	repo infrastructure.IUserRepository
}

func (u UserService) Save(ctx context.Context, book *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FindByID(ctx context.Context, id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FindAll(ctx context.Context) ([]domain.User, error) {
	all, err := u.repo.FindAll()
	if err != nil {
		return []domain.User{}, err
	}
	return all, nil
}

// NewBookService سازنده UserService
func NewBookService(repo infrastructure.IUserRepository) IUserService {
	return UserService{repo: repo}
}
