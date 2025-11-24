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

func (r *CodePicsRepository) Save(codePic *domain.CodePic) (domain.CodePic, error) {
	result := r.DB.Create(codePic)
	if result.Error != nil {
		return domain.CodePic{}, result.Error
	}

	return *codePic, nil
}
