package models

type Product struct {
	Id           int    `json:"id" gorm:"column:id"`
	ProductName  string `json:"product_name" gorm:"column:product_name"`
	ProductCode  string `json:"product_code" gorm:"column:product_code"`
	ProductPrice string `json:"product_price" gorm:"column:product_price"`
}

func (Product) TableName() string {
	return "products"
}
