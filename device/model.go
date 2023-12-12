package device

import "time"

type Device struct {
	UUID      string `gorm:"PrimaryKey;UniqueIndex;AutoIncrement"`
	Name      string
	CreatedAt time.Time
	ActiveAt  time.Time
}
