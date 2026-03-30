package repository

import (
	"github.com/Hamse/final_project/Back_end/Back_end/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}



func RegisterCategoryRepo(db *gorm .DB) *CategoryRepo{
	return  &CategoryRepo {
	DB: db,
	}

}

func (repo CategoryRepo) CreateCategory(data models.Category) error{
	return repo.DB.Create(&data).Error
}





func (repo *CategoryRepo) GetAllCategories() ([]models.Category, error) {
    var categories []models.Category
    err := repo.DB.Find(&categories).Error
    return categories, err
}










func (repo *CategoryRepo) UpdateCategory(id uint, newName string) error {
    var category models.Category
    if err := repo.DB.First(&category, id).Error; err != nil {
        return err
    }
    category.CategoryName = newName
    return repo.DB.Save(&category).Error
}






func (r *CategoryRepo) DeleteCategory(id uint) error {
    var category models.Category
    if err := r.DB.First(&category, id).Error; err != nil {
        return err
    }
    return r.DB.Delete(&category).Error
}





// Count products per category
func (r *CategoryRepo) GetProductCountPerCategory() ([]map[string]interface{}, error) {
    var results []map[string]interface{}
    err := r.DB.Model(&models.Product{}).
        Select("category_id, COUNT(*) as product_count").
        Group("category_id").
        Find(&results).Error
    return results, err
}

// Check if products exist under category
func (r *CategoryRepo) HasProducts(id uint) (bool, error) {
    var count int64
    err := r.DB.Model(&models.Product{}).Where("category_id = ?", id).Count(&count).Error
    return count > 0, err
}