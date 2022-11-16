package models

import "errors"

type ProductRequest struct {
	ProductName string `json:"product_name"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
	Count       int    `json:"count"`
	Sizes       []int  `json:"sizes"`
}

func (i ProductRequest) Validate() error {
	if i.ProductName == "" && i.CategoryId == 0 && i.Price == 0 && i.Color == "" && i.Count == 0 {
		return errors.New("update structure has no values")
	}

	return nil
}

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	CategoryId   int    `json:"category_id"`
	Price        int    `json:"price"`
	Color        string `json:"color"`
	Count        int    `json:"count"`
	CategoryName string `json:"category_name"`
	Sizes        []int  `json:"available_sizes"`
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type Size struct {
	ID      int `json:"id"`
	SizeNum int `json:"size_num"`
}

type AvailableSize struct {
	AvailableSize int `json:"available_size"`
}

type SizeInput struct {
	SizeNum int `json:"size_num"`
}

type ProductSizes struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	SizeID    int `json:"size_id"`
}
