package models

import (
	"github.com/jinzhu/gorm"
)

// Product struct for product model
type Product struct {
	ProductCode        string  `gorm:"primary_key;column:productCode" json:"productCode" binding:"required"`
	ProductName        string  `gorm:"column:productName" json:"productName" binding:"required"`
	ProductLine        string  `gorm:"column:productLine" json:"productLine" binding:"required"`
	ProductScale       string  `gorm:"column:productScale" json:"productScale" binding:"required"`
	ProductVendor      string  `gorm:"column:productVendor" json:"productVendor" binding:"required"`
	ProductDescription string  `gorm:"column:productDescription" json:"productDescription" binding:"required"`
	QuantityInStock    int64   `gorm:"column:quantityInStock" json:"quantityInStock" binding:"required"`
	BuyPrice           float32 `gorm:"column:buyPrice" json:"buyPrice" binding:"required"`
	MSRP               float32 `gorm:"column:MSRP" json:"MSRP" binding:"required"`
}

// ExistProduct checks if a product exists based on ProductCode
func ExistProduct(pk string) bool {
	var p Product
	err := db.Select("ProductCode").Where("ProductCode =?", pk).First(&p).Error

	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

// ProductList get a list of  products based on paging constraints
func ProductList(pageNum int, pageSize int) ([]Product, error) {
	var (
		products []Product
		err      error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&products).Error
	} else {
		err = db.Find(&products).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return products, nil
}

// ProductRetrieve retrieve one product based on pprimary key
func ProductRetrieve(pk string) (Product, error) {
	var p Product

	db.Where("productCode =?", pk).Find(&p)
	return p, nil
}

// ProductCreate create one product
func ProductCreate(product Product) (Product, error) {
	if err := db.Create(&product).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

// ProductUpdate update one product based on primary key
func ProductUpdate(pk string, maps interface{}) (Product, error) {
	var (
		p   Product
		err error
	)

	if err = db.Model(&Product{ProductCode: pk}).Updates(maps).Error; err != nil {
		return Product{}, err
	}
	db.Where(&Product{ProductCode: pk}).Find(&p)
	return p, nil
}

// ProductDestroy destroy one product based on primary key
func ProductDestroy(pk string) error {
	var (
		p   Product
		err error
	)

	db.Where("productCode =?", pk).Find(&p)
	if err = db.Delete(&p).Error; err != nil {
		return err
	}

	return nil
}
