package routes

import (
	"github.com/Hamse/final_project/Back_end/Back_end/handler"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
    apiGroup := r.Group("/api")

    categoryHandler := handler.RegisterCategoryHandler()

    categoryGroup := apiGroup.Group("/categories")
    {
        categoryGroup.POST("/create", categoryHandler.CreateCategoryHandler)
        categoryGroup.GET("/all", categoryHandler.GetCategoriesHandler)
        categoryGroup.PUT("/:id", categoryHandler.UpdateCategoryHandler)
        categoryGroup.DELETE("/:id", categoryHandler.DeleteCategoryHandler)
        categoryGroup.GET("/product-count", categoryHandler.GetProductCountHandler) 
    }

    productHandler := handler.RegisterProductHandler()

    productGroup := apiGroup.Group("/products")
    {
        productGroup.POST("/create", productHandler.CreateProductHandler)
       
    }
}