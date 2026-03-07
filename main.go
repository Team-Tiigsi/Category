package main

import (
	"log/slog"

	"github.com/Hamse/final_project/Back_end/infra"
	"github.com/Hamse/final_project/Back_end/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.InitEnv()
	config := infra.Configurations

	infra .ConnectedDB()
	// infra.ConnectDB()
	slog.Info("connected database successfully ✅")

	r := gin.Default()

	routes.RegisterCategoryRoutes(r)

	slog.Info("Application is running on ", "port", config.Port)
	r.Run(config.Port) // config.Port should be ":9000"
}