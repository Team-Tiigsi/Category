package handler

import (
	"fmt"
	"net/http"

	"github.com/Hamse/final_project/Back_end/Back_end/dtos"
	"github.com/Hamse/final_project/Back_end/Back_end/infra"
	"github.com/Hamse/final_project/Back_end/Back_end/repository"
	"github.com/Hamse/final_project/Back_end/Back_end/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}
 
// RegisterCategoryHandler - same style as UserHandler
func RegisterCategoryHandler() CategoryHandler {
	categoryRepo := repository .RegisterCategoryRepo(infra.DB)
	categorySvc := service.RegisterCategoryService(categoryRepo)

	return CategoryHandler{
		CategoryService: categorySvc,
	}
}

// CreateCategoryHandler - POST /categories
func (h *CategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var requestBody dtos.AddCategorydtos

	// Bind JSON body
	err := c.ShouldBindJSON(&requestBody)
	fmt.Println("from body:", requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			
			"message":    "failed to bind body request",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.CategoryService.AddCategory(&requestBody)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"is_success": false,
			"err":    err.Error(),
			"message":"waa kagu fail garobey",
			
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":   "successfully created category",
		"is_success": true,
	})
}


func (h *CategoryHandler) GetCategoriesHandler(c *gin.Context) {
    categories, err := h.CategoryService.GetCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "is_success": false,
            "message":    "failed to fetch categories",
            "error":      err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "is_success": true,
        "data":       categories,
    })
}





// PUT /categories/:id
func (h *CategoryHandler) UpdateCategoryHandler(c *gin.Context) {
    // Get ID from URL
    idParam := c.Param("id")
    var requestBody struct {
        CategoryName string `json:"category_name" binding:"required"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "is_success": false,
            "message":    "invalid request body",
            "error":      err.Error(),
        })
        return
    }

    // Convert idParam to uint
    var id uint
    _, err := fmt.Sscan(idParam, &id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "is_success": false,
            "message":    "invalid category id",
        })
        return
    }

    // Call service
    err = h.CategoryService.UpdateCategory(id, requestBody.CategoryName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "is_success": false,
            "message":    "failed to update category",
            "error":      err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "is_success": true,
        "message":    "category updated successfully",
    })
}









// GET /categories/product-count
func (h *CategoryHandler) GetProductCountHandler(c *gin.Context) {
    results, err := h.CategoryService.GetProductCountPerCategory()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "is_success": false,
            "message":    "failed to fetch product counts",
            "error":      err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "is_success": true,
        "data":       results,
    })
}

// DELETE /categories/:id
func (h *CategoryHandler) DeleteCategoryHandler(c *gin.Context) {
    idParam := c.Param("id")
    var id uint
    _, err := fmt.Sscan(idParam, &id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "is_success": false,
            "message":    "invalid category id",
            "err":err.Error(),
        })
        return
    }

    err = h.CategoryService.DeleteCategory(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "is_success": false,
            "message":    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "is_success": true,
        "message":    "category deleted successfully",
    })
}