package domain

type Role struct {
	ID        int8   `json:"ID,omitempty"`
	Name      string `json:"name,omitempty"`
	Users     []User `gorm:"many2many:user_roles;"`
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
