package cmd

import (
	"context"
	"time"

	"github.com/go-mpesa-integration/cmd/server"
	"github.com/go-mpesa-integration/config"
	"github.com/go-mpesa-integration/db"
	"github.com/go-mpesa-integration/provider"
	"github.com/sirupsen/logrus"
)

func Run(){
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background();
	envConfig := config.NewEnvConfig()

	db :=  db.Init(envConfig , db.DBMigrator)

	provider.NewProvider(db , server)

	go func(){
		if err := server.Start(ctx , envConfig.AppPort); err != nil {
			logrus.Errorf("Error starting server %v", err)
		}

	}()
	<-ctx.Done()
	logrus.Info("Server stopped")

	shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.ShutDown(shutdown); err != nil {
		logrus.Errorf("Error shutting down server %v", err)
	}
	

}