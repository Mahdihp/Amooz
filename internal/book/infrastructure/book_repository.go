package infrastructure

import (
	"Amooz/internal/book/domain"
	"context"
	"errors"
)

// BookRepository مخزن کتاب‌ها است که داده‌ها را ذخیره می‌کند
type BookRepository struct {
	books map[string]*domain.Book
}

// NewBookRepository سازنده برای BookRepository
func NewBookRepository() *BookRepository {
	return &BookRepository{
		books: make(map[string]*domain.Book),
	}
}

// Save کتاب را ذخیره می‌کند
func (r *BookRepository) Save(ctx context.Context, book *domain.Book) error {
	r.books[book.ID] = book
	return nil
}

// FindByID کتابی را با شناسه پیدا می‌کند
func (r *BookRepository) FindByID(ctx context.Context, id string) (*domain.Book, error) {
	book, exists := r.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}
