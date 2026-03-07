package handler

import (
    "net/http"
    "github.com/Hamse/final_project/Back_end/dtos"
    "github.com/Hamse/final_project/Back_end/infra"
    "github.com/Hamse/final_project/Back_end/repository"
    "github.com/Hamse/final_project/Back_end/service"
    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    ProductService *service.ProductService
}

func RegisterProductHandler() ProductHandler {
    productRepo := repository.RegisterProductRepo(infra.DB)
    productSvc := service.RegisterProductService(productRepo)
    return ProductHandler{ProductService: &productSvc}
}

// POST /products/create
func (h *ProductHandler) CreateProductHandler(c *gin.Context) {
    var requestBody dtos.AddProductdtos
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "is_success": false,
            "message":    "failed to bind body request",
            "error":      err.Error(),
        })
        return
    }

    productID, err := h.ProductService.AddProduct(&requestBody)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "is_success": false,
            "message":    "failed to create product",
            "error":      err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "is_success": true,
        "message":    "product created successfully",
        "product_id": productID,
    })
}