package routes

import (
	"github.com/briangtn/codepic/internal/api/codepics"
	"github.com/briangtn/codepic/internal/repositories"
	"github.com/briangtn/codepic/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	CodePicsHandler *codepics.Handler
}

func RouterSetup(r *gin.Engine, db *gorm.DB) {
	MiddlewaresSetup(r)
	handlers := SetupDependencies(db)

	codepicsGroup := r.Group("/codepics")
	{
		codepicsGroup.POST("/", handlers.CodePicsHandler.CreateCodePic)
		codepicsGroup.GET("/:id", handlers.CodePicsHandler.GetCodePicByID)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

func MiddlewaresSetup(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func SetupDependencies(db *gorm.DB) *Handlers {
	codePicsRepository := repositories.NewCodePicsRepository(db)
	codePicsService := services.NewCodePicsService(codePicsRepository)

	return &Handlers{
		CodePicsHandler: codepics.NewHandler(codePicsService),
	}
}
