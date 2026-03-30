package models


type role string

const(
    RoleAdmin role=`json:"role"`
)

type Product struct {
    ID          uint   `json:"ID"`
    Name        string `json:"name"`
    CategoryID  uint   `json:"ccategory_id"`
    Category    Category `json:"CategoryID"`
}