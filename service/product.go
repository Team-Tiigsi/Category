package service

import (
	"errors"

	"github.com/Hamse/final_project/Back_end/dtos"
	"github.com/Hamse/final_project/Back_end/models"
	"github.com/Hamse/final_project/Back_end/repository"
)

type ProductService struct {
    Repo repository.ProductRepo
}

func RegisterProductService(repo repository.ProductRepo) ProductService {
    return ProductService{Repo: repo}
}

func (s *ProductService) AddProduct(dto *dtos.AddProductdtos) (uint, error) {
    product := models.Product{
        Name:       dto.Name,
        CategoryID: dto.CategoryID,
    }
    err := s.Repo.AddProduct(&product)
    return product.ID, err
}














func (svc *CategoryService) GetProductCountPerCategory() ([]map[string]interface{}, error) {
    return svc.repo.GetProductCountPerCategory()
}

func (svc *CategoryService) DeleteCategory(id uint) error {
    hasProducts, err := svc.repo.HasProducts(id)
    if err != nil {
        return err
    }
    if hasProducts {
        return errors.New("cannot delete category: products exist under it")
    }
    return svc.repo.DeleteCategory(id)
}