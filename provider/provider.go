package provider

import (
	"github.com/go-mpesa-integration/cmd/server"
	"github.com/go-mpesa-integration/internals/handler"
	"github.com/go-mpesa-integration/internals/repository"
	"github.com/go-mpesa-integration/internals/routes"
	"github.com/go-mpesa-integration/internals/services"
	"gorm.io/gorm"
)



func NewProvider(db  *gorm.DB , server server.GinServer){
	mpesaRepo := repository.NewMpesaRepository(db)
	mpesaService := services.NewMpesaService(mpesaRepo)
	mpesaHandler := handler.NewMpesaHandler(mpesaService)
	routes.RegisterMpesaRoutes(server , mpesaHandler)
}