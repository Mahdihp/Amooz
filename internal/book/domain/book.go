package domain

import "time"

// Book نماینده یک کتاب در سیستم است
type Book struct {
	ID          string
	Title       string
	Author      Author
	Publisher   string
	PublishedAt time.Time
}
