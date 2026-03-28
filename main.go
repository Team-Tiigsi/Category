package main

import (
	"log/slog"

	"github.com/Hamse/final_project/Back_end/Back_end/infra"
	"github.com/Hamse/final_project/Back_end/Back_end/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// infra.InitEnv()
	// config := infra.Configurations

	infra .ConnectedDB()
	// infra.ConnectDB()
	slog.Info("connected database successfully ✅")

	r := gin.Default()

	routes.RegisterCategoryRoutes(r)

	// slo'g.Info("Application is running on ", "port", config.Port)
	r.Run(":8080") // config.Port should be ":9000"
}