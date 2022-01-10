package role

type Role struct {
	Id      int    `gorm:"not null;uniqueIndex;primary_key"`
	Name    string `gorm:"size:255;not null"`
	Display string `gorm:"size:255;not null"`
}
