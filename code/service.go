package code

import "strings"

func claerPackageName(serviceName string) string {
	return strings.NewReplacer(".", "", "/", "", "\\", "", "-", "").Replace(strings.ToLower(serviceName))
}

func Repository(serviceName string) string {
	return `package ` + claerPackageName(serviceName) + `

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

func Service(serviceName string) string {
	return `package ` + claerPackageName(serviceName) + `

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

func Handler(serviceName string) string {
	return `package ` + claerPackageName(serviceName) + `

type handler struct {
	service_ Service
}

func NewHandler(service_ Service) *handler {
	return &handler{service_}
}
	`
}

func Router(serviceName string) string {
	return `package ` + claerPackageName(serviceName) + `

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	// repository := NewRepository(db)
	// service := NewService(repository)
	// handler := NewHandler(service)
	
	//` + serviceName + ` := g.Group("/` + strings.ToLower(serviceName) + `")
	// ` + serviceName + `.GET("", handler.get)
	
}
	`
}
