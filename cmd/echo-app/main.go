package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/igorralexsander-corp/echoapp"
	"github.com/igorralexsander-corp/echoapp/internal/infrastructure/api"
	"github.com/igorralexsander-corp/echoapp/internal/infrastructure/logger"
	"github.com/labstack/echo"
)

func main() {

	logger.NewEasyZap()

	logger.Info("Starting")

	application := app.NewApplication()

	httpServer := api.NewHttpServer(&application)

	go api.Start(httpServer, "0.0.0.0:8080")

	shutDownHook(&application, httpServer, 5)

}

func shutDownHook(app *app.Application, apiServer *echo.Echo, terminationTime int32) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	switch <-quit {
	case os.Interrupt, syscall.SIGTERM:
		logger.Info("Initialize Gracefully shutdown")
		logger.Info(fmt.Sprintf("Wait %d seconds to process pending requests", terminationTime))

		time.Sleep(time.Duration(terminationTime) * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		logger.Info("Shutdown HTTP server...")

		if err := apiServer.Shutdown(ctx); err != nil {
			logger.Fatal(err, "Error to gracefully stop application, application stopped.")
		}
		logger.Info("Complete Gracefully shutdown")
	}

}
