package routes

import (
	"github.com/go-mpesa-integration/cmd/server"
	"github.com/go-mpesa-integration/internals/handler"
)


func RegisterMpesaRoutes(server server.GinServer ,mpesaHandler *handler.MpesaHandler) {
	server.RegisterRoute("POST","/api/v1/stkpush", mpesaHandler.InitiateSTKPush)
	server.RegisterRoute("POST","/api/v1/stkcallback",mpesaHandler.STKCallback)
}
