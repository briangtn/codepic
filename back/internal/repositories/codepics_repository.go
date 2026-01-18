package repositories

import (
	"github.com/briangtn/codepic/internal/domain"
	"gorm.io/gorm"
)

type CodePicsRepository struct {
	DB *gorm.DB
}

func NewCodePicsRepository(db *gorm.DB) *CodePicsRepository {
	return &CodePicsRepository{
		DB: db,
	}
}

func (r *CodePicsRepository) Save(codePic *domain.CodePic) (*domain.CodePic, error) {
	result := r.DB.Create(codePic)
	if result.Error != nil {
		return nil, result.Error
	}

	return codePic, nil
}

func (r *CodePicsRepository) FindByID(id string) (*domain.CodePic, error) {
	var codePic domain.CodePic
	result := r.DB.First(&codePic, "id = ?", id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &codePic, nil
}

func (r *CodePicsRepository) Delete(codePic *domain.CodePic) error {
	result := r.DB.Delete(codePic)
	return result.Error
}

func (r *CodePicsRepository) Update(codePic *domain.CodePic) (*domain.CodePic, error) {
	result := r.DB.Save(codePic)
	if result.Error != nil {
		return nil, result.Error
	}

	return codePic, nil
}
