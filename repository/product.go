package repository

import (
	"github.com/Hamse/final_project/Back_end/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
    DB *gorm.DB
}

func RegisterProductRepo(db *gorm.DB) ProductRepo {
    return ProductRepo{DB: db}
}

func (r *ProductRepo) AddProduct(product *models.Product) error {
    return r.DB.Create(product).Error
}