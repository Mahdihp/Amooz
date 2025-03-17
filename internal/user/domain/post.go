package domain

type Post struct {
	ID        int64 `json:"ID,omitempty" gorm:"primaryKey"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`

	Title   string `json:"title,omitempty"`
	Deleted bool   `json:"deleted"`
	UserID  uint   `json:"userID"`
}
