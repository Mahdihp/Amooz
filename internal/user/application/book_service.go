package application

import (
	"Amooz/internal/user/domain"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

// BookRepository اینترفیس برای مخزن کتاب‌ها
type BookRepository interface {
	Save(ctx context.Context, book *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

// BookService سرویس مدیریت کتاب‌ها
type BookService struct {
	repo BookRepository
}

// NewBookService سازنده BookService
func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook کتاب جدیدی ایجاد می‌کند
func (s *BookService) CreateBook(ctx context.Context, title string, role domain.Role, publisher string, publishedAt time.Time) (*domain.User, error) {
	book := &domain.User{
		ID:          uuid.NewString(),
		Title:       title,
		Role:        role,
		Publisher:   publisher,
		PublishedAt: publishedAt,
	}
	err := s.repo.Save(ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// FindBookByID کتابی را با شناسه پیدا می‌کند
func (s *BookService) FindBookByID(ctx context.Context, id string) (*domain.User, error) {
	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return book, nil
}
