package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type GinServer interface {
	Start(ctx context.Context , httpAddress string) error
	ShutDown(ctx context.Context) error
	RegisterRoute(method string, path string, handlerFunc gin.HandlerFunc)
}

type GinserverBuilder struct {}

type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

//cerate a new instance of GinserverBuilder

func NewGinServerBuilder() *GinserverBuilder {
	return &GinserverBuilder{}
}

//method to build gin server

func (b *GinserverBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

// start server


func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	// go routine  to handle start server
	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Failed to start server: %s", err)
		}
	}()
	logrus.Infof("server started successfully at port 🚀🚀 %s", httpAddress)
	return nil
}

func (gs *ginServer) ShutDown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Failed to shutdown server: %s", err)
		return err
	}
	logrus.Info("Server shutdown successfully")
	return nil
}

// methdd to register a route

func (gs *ginServer) RegisterRoute(method string, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	default:
		logrus.Errorf("Invalid Http method")
	}
}
