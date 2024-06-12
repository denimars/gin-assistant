package code

import "strings"

func Repository(nameService string) string {
	return `package ` + nameService + `

import (
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

`

}

func Service(nameService string) string {
	return `package ` + nameService + `

type Service interface {

}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}
`
}

func Handler(nameService string) string {
	return `package ` + nameService + `

type handler struct {
	service_ Service
}

func NewHandler(service_ Service) *handler {
	return &handler{service_}
}
	`
}

func Router(serviceName string) string {
	return `package ` + serviceName + `

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ` + strings.ToTitle(strings.ToLower(serviceName)) + `Router(g *gin.RouterGroup, db *gorm.DB) {
	// repository := NewRepository(db)
	// service := NewService(repository)
	// handler := NewHandler(service)
	
	//project := g.Group("")
	// project.GET("", handler.get)
	
}
	`
}
