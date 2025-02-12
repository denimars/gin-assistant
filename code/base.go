package code

func Base() string {
	return `
package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         ` + "`" + `gorm:"type:varchar(255)" json:"id"` + "`" + `
	CreatedAt time.Time      ` + "`" + `json:"-" gorm:"type:datetime(0)"` + "`" + `
	UpdatedAt time.Time      ` + "`" + `json:"-" gorm:"type:datetime(0)"` + "`" + `
	DeletedAt gorm.DeletedAt ` + "`" + `json:"-" gorm:"type:datetime(0)"` + "`" + `
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New(rand.NewSource(time.Now().UnixNano())))
	b.ID = id.String()
	return
}

func CustomJson(c *gin.Context, status int, obj interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(status)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Writer.Write(buffer.Bytes())
}

func JSON(c *gin.Context, err error, data interface{}, responStatus ...int) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "uuups..."})
		return
	} else if len(responStatus) > 0 {
		c.JSON(responStatus[0], data)
		return
	}
	c.JSON(http.StatusOK, data)
}

	`
}
