package domain

type User struct {
	ID        int64 `json:"ID,omitempty" gorm:"primaryKey"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`

	DisplayName string `json:"displayName,omitempty"`
	Username    string `json:"username" gorm:"uniqueIndex;type:varchar(100);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	Deleted     bool   `json:"deleted"`
	Roles       []Role `json:"roles" gorm:"many2many:user_roles;"`
	Posts       []Post `json:"posts" gorm:"foreignKey:UserID"`
}
