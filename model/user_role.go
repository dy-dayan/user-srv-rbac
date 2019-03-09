package model

type UserRole struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`
	UID int64 `gorm:"index"`
	RoleID uint
	CreatedAt int64
}

func (u *UserRole) TableName() string {
	return "user_role"
}