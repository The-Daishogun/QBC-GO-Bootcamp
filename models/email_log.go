package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type EmailLog struct {
	gorm.Model
	UserID  uint
	SentAt  sql.NullTime
	Title   string
	Content string
}
