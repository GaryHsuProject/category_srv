package category

import (
	"shop/models"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(*models.Category) error
	FindAll() ([]*models.Category, error)
}

type CategoryRepository struct {
	orm *gorm.DB
}

func NewCategoryRepository(orm *gorm.DB) Repository {
	return &CategoryRepository{orm: orm}
}

func (u *CategoryRepository) Insert(category *models.Category) error {
	if err := u.orm.Create(&category).Error; err != nil {
		return err
	}
	return nil
}
func (u *CategoryRepository) FindAll() ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	if err := u.orm.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
