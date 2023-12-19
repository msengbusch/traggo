package user

type Role struct {
	ID   int    `gorm:"primary_key;unique_index;auto_increment"`
	Name string `gorm:"unique"`
}

const (
	CreateRole Permission = iota
	ReadRole
)
