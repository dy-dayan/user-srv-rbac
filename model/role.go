package model

type Role struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name string // 角色名称
	State int //状态 1：有效， 0：无效
	CreatedAt int64
	UpdatedAt int64
}

const (
	RoleStateActive = iota
	RoleStateInactive
)

func (r *Role) TableName() string {
	return "role"
}

