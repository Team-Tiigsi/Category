package service

import (
	"github.com/Hamse/final_project/Back_end/Back_end/dtos"
	"github.com/Hamse/final_project/Back_end/Back_end/models"
	"github.com/Hamse/final_project/Back_end/Back_end/repository"
)

// "os/user"
// "strings"

type CategoryService struct {
	repo *repository.CategoryRepo
}

func RegisterCategoryService(repo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (svc CategoryService) AddCategory(data *dtos.AddCategorydtos) (int, error) {
	// name := strings.ToLower(data.CategoryName)

	category := models.Category{
		CategoryName: data.CategoryName,
	}

	err := svc.repo.CreateCategory(category)
	if err != nil {
		return 0, err
		
	}

	return category.ID, nil
}










func (svc *CategoryService) GetCategories() ([]models.Category, error) {
	return svc.repo.GetAllCategories()
}




func (svc *CategoryService) UpdateCategory(id uint, newName string) error {
    // return svc.Repo.UpdateCategory(id, newName)
	return  svc.repo.UpdateCategory(id,newName)
}
