package code

func Middleware(projectName string) string {
	return `
package middleware

import (
	"` + projectName + `/app/helper"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
}

type middleware struct {
	repo Repository
}

type message struct {
	Message string ` + "`" + `json:"message"` + "`" + `
}

func NewMiddleware(repo Repository) *middleware {
	return &middleware{repo}
}

func (m *middleware) cekTokenExist(c *gin.Context) bool {
	token := helper.GetToken(c)
	blackListTOken := m.repo.FindToken(token)
	return blackListTOken.Token == ""
}

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message message
		message.Message = "Unauthorized"
		isBlackList := m.cekTokenExist(c)
		statusNext := true
		var token *jwt.Token
		var err error
		if !isBlackList {
			message.Message = "Token has been blacklisted"
			statusNext = false
		} else {
			token, err = helper.TokenValidate(c)
			if err != nil {
				statusNext = false
				if err.Error() == "Token is expired" {
					message.Message = err.Error()
				}
			}
		}
		if !statusNext {
			c.JSON(http.StatusUnauthorized, message)
			c.Abort()
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("id", claims["id"])
		c.Next()

	}

}
	`
}

func ModelBlackListToken(projectName string) string {
	return `
package middleware

import "` + projectName + `/app/service"

type BlacklistToken struct {
	service.BaseModel
	Token string ` + "`" + `gorm:"type:varchar(255)"` + "`" + `
}

func (BlacklistToken) TableName() string {
	return "blacklist_tokens"
}
	`
}

func RepositoryBlackListToken() string {
	return `
package middleware

import "gorm.io/gorm"

type Repository interface {
	FindToken(token string) BlacklistToken
	Save(token string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindToken(token string) BlacklistToken {
	var blackListToken BlacklistToken
	r.db.Where("token = ?", token).Find(&blackListToken)
	return blackListToken
}

func (r *repository) Save(token string) error {
	err := r.db.Create(&BlacklistToken{
		Token: token,
	}).Error
	return err
}

	`
}
