package code

func Base() string {
	return `
package service

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         ` + "`" + `gorm:"primaryKey" json:"id"` + "`" + `
	CreatedAt time.Time      ` + "`" + `json:"created_at" gorm:"type:datetime(0)"` + "`" + `
	UpdatedAt time.Time      ` + "`" + `json:"updated_at" gorm:"type:datetime(0)"` + "`" + `
	DeletedAt gorm.DeletedAt ` + "`" + `json:"deleted_at" gorm:"tyupe:datetime(0)"` + "`" + `
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type Response struct {
	DeletedAt gorm.DeletedAt ` + `json:"-"
}
	`
}
