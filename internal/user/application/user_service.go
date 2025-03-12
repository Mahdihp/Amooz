package application

import (
	"Amooz/internal/user/domain"
	"Amooz/internal/user/infrastructure"
	"context"
)

// UserRepository اینترفیس برای مخزن کتاب‌ها
type UserRepository interface {
	Save(ctx context.Context, book *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

// UserService سرویس مدیریت کتاب‌ها
type UserService struct {
	repo infrastructure.UserRepository
}

// NewBookService سازنده UserService
func NewBookService(repo infrastructure.UserRepository) UserRepository {
	return &UserService{repo: repo}
}

func (u UserService) Save(ctx context.Context, book *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) FindByID(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
