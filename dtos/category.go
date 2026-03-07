package dtos


type AddCategorydtos struct {
    CategoryName string `json:"category_name" binding:"required"`
}




type AddProductdtos struct {
    Name       string `json:"name" binding:"required"`
    CategoryID uint   `json:"category_id" binding:"required"`
}
