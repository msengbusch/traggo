package user

type User struct {
	ID       int    `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Name     string `gorm:"unique"`
	Password []byte
}
