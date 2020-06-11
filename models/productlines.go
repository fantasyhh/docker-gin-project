package models

import "github.com/jinzhu/gorm"

// ProductLine struct for productline model
type ProductLine struct {
	ProductLine     string    `gorm:"primary_key;column:productLine" json:"productLine" binding:"required"`
	TextDescription string    `gorm:"column:textDescription" json:"textDescription"`
	HTMLDescription string    `gorm:"column:htmlDescription" json:"htmlDescription"`
	Image           string    `gorm:"column:image" json:"image"`
	Products        []Product `gorm:"foreignKey:ProductLine"`
}

// TableName 指定表的名称，默认会将ProductLine转换为product_lines
func (ProductLine) TableName() string {
	return "productlines"
}

// ExistProductLine checks if a productline exists based on ProductLine (primary key)
func ExistProductLine(pk string) bool {
	var pl ProductLine
	err := db.Select("ProductLine").Where("ProductLine =?", pk).First(&pl).Error

	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

// ProductLineList get a list of productlines based on paging constraints
func ProductLineList(pageNum int, pageSize int) ([]ProductLine, error) {
	var (
		productlines []ProductLine
		err          error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&productlines).Error
	} else {
		err = db.Find(&productlines).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return productlines, nil
}

// ProductLineRetrieve retrieve one productline based on primary key
func ProductLineRetrieve(pk string) (ProductLine, error) {
	var pl ProductLine

	db.Where(&ProductLine{ProductLine: pk}).Find(&pl)
	//db.Preload("Products" ).Where(&ProductLine{ProductLine: pk}).Find(&pl)
	return pl, nil
}

// ProductLineCreate create one productline
func ProductLineCreate(productline ProductLine) (ProductLine, error) {
	if err := db.Create(&productline).Error; err != nil {
		return ProductLine{}, err
	}

	return productline, nil
}

// ProductLineUpdate update one productline  based on primary key
func ProductLineUpdate(pk string, maps interface{}) (ProductLine, error) {
	var (
		pl  ProductLine
		err error
	)

	if err = db.Model(&ProductLine{ProductLine: pk}).Updates(maps).Error; err != nil {
		return ProductLine{}, err
	}
	db.Where(&ProductLine{ProductLine: pk}).Find(&pl)
	return pl, nil
}

// ProductLineDestroy destroy one productline  based on primary key
func ProductLineDestroy(pk string) error {
	var (
		pl  ProductLine
		err error
	)

	db.Where(&ProductLine{ProductLine: pk}).Find(&pl)
	if err = db.Delete(&pl).Error; err != nil {
		return err
	}

	return nil
}

// ProductInLine get a list of products in a specify productline
func ProductInLine(pk string) ([]Product, error) {
	pl := ProductLine{ProductLine: pk}

	if err := db.Model(&pl).Related(&pl.Products, "products").Error; err != nil {
		return nil, err
	}

	return pl.Products, nil
}
