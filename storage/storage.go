package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type storage struct {
	*gorm.DB
}

func New(connection string) (*storage, error) {
	gormDb, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := gormDb.AutoMigrate(&Product{}); err != nil {
		return nil, err
	}

	return &storage{gormDb}, nil
}

func (db *storage) CreateTriggers() error {
	if err := db.beforeProductInsertTrigger(); err != nil {
		return err
	}
	return db.beforeProductUpdateTrigger()
}

func (db *storage) CreateProducts(products ...Product) error {
	return db.Create(&products).Error
}

func (db *storage) UpdateProduct(product Product) error {
	if _, err := db.Product(product.Name); err != nil {
		return err
	}
	return db.Where("name = ?", product.Name).Updates(product).Error
}

func (db *storage) Product(name string) (Product, error) {
	var product Product
	return product, db.Where("name = ?", name).First(&product).Error
}
