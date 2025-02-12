package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime(0)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime(0)"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"tyupe:datetime(0)"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New(rand.NewSource(uint64(time.Now().UnixNano()))))
	b.ID = id.String()
	return
}

func CustomJson(c *gin.Context, status int, obj interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(status)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false) // Disable HTML escaping
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Write(buffer.Bytes())
}

func JSON(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "uuups..."})
		return
	}
	c.JSON(http.StatusOK, data)
}