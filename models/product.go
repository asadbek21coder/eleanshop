package models

import (
	"errors"
	"mime/multipart"
)

type ProductRequest struct {
	ProductName string `json:"product_name"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Color       string `json:"color"`
	Count       int    `json:"count"`
	Sizes       []int  `json:"sizes"`
	ImageUrl    string `json:"image_url"`
}

type FakeProduct struct {
	ProductName string                `form:"product_name"`
	CategoryId  int                   `form:"category_id"`
	Price       int                   `form:"price"`
	Color       string                `form:"color"`
	Count       int                   `form:"count"`
	Sizes       string                `form:"sizes"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
}

func (i ProductRequest) Validate() error {
	if i.ProductName == "" && i.CategoryId == 0 && i.Price == 0 && i.Color == "" && i.Count == 0 {
		return errors.New("update structure has no values")
	}

	return nil
}

type Product struct {
	ID             int    `json:"id"`
	ProductName    string `json:"product_name"`
	CategoryId     int    `json:"category_id"`
	Price          int    `json:"price"`
	Color          string `json:"color"`
	Count          int    `json:"count"`
	CategoryName   string `json:"category_name"`
	ImageUrl       string `json:"image_url"`
	AvailableSizes []int  `json:"available_sizes"`
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type Size struct {
	ID      int `json:"id" db:"id"`
	SizeNum int `json:"size_num" db:"size_num"`
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

type ResponseProduct struct {
	ProductName string                `json:"product_name"`
	Image       *multipart.FileHeader `json:"image"`
}

type QueryParams struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type FakeProductModel struct {
	ProductName    string `json:"product_name" db:"product_name"`
	CategoryName   string `json:"category_name" db:"category_name"`
	AvailableSizes []Size `json:"available_sizes"`
}
