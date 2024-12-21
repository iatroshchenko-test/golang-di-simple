package post

type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Uuid        string `gorm:"not null"`
	Title	    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Content string     `gorm:"not null"`
}

func (Post) TableName() string {
	return "posts"
}
