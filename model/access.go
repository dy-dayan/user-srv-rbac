package model

type Access struct {
	ID uint `gorm:"primary_key; AUTO_INCREMENT"`
	Title string
	State int
	CreatedAt int64
	UpdatedAt int64
}


const (
	AccessStateActive = iota
	AccessStateInactive
)

func (r *Access) TableName() string {
	return "access"
}