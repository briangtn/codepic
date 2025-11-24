package services

import (
	"github.com/briangtn/codepic/internal/domain"
	"github.com/briangtn/codepic/internal/repositories"
	"github.com/google/uuid"
)

type CodePicsService struct {
	repository *repositories.CodePicsRepository
}

func NewCodePicsService(repository *repositories.CodePicsRepository) *CodePicsService {
	return &CodePicsService{
		repository: repository,
	}
}

func (codePicsService *CodePicsService) CreateCodePic(title, content, language string, maxViews int) (domain.CodePic, error) {
	codepic := &domain.CodePic{
		ID:       uuid.New().String(),
		Title:    title,
		Content:  content,
		Language: language,
		MaxViews: maxViews,
	}

	return codePicsService.repository.Save(codepic)
}
