package code

import "strings"

func Repository(nameService string) string {
	return `package ` + strings.ToLower(nameService) + `

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
	return `package ` + strings.ToLower(nameService) + `

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
	return `package ` + strings.ToLower(nameService) + `

type handler struct {
	service_ Service
}

func NewHandler(service_ Service) *handler {
	return &handler{service_}
}
	`
}

func Router(nameService string) string {
	return `package ` + strings.ToLower(nameService) + `

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	// repository := NewRepository(db)
	// service := NewService(repository)
	// handler := NewHandler(service)
	
	//` + nameService + ` := g.Group("")
	// ` + nameService + `.GET("", handler.get)
	
}
	`
}
