package user

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
