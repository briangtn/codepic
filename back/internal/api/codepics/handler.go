package codepics

import (
	"github.com/briangtn/codepic/internal/api/responses"
	"github.com/briangtn/codepic/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	codePicsService *services.CodePicsService
}

func NewHandler(codePicsService *services.CodePicsService) *Handler {
	return &Handler{
		codePicsService: codePicsService,
	}
}

func (h *Handler) CreateCodePic(ctx *gin.Context) {
	var payload Payload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, responses.ErrorResponse{
			Code:    responses.INVALID_PAYLOAD_ERROR_CODE,
			Message: err.Error(),
		})
		return
	}

	codepic, err := h.codePicsService.CreateCodePic(payload.Title, payload.Content, payload.Language, payload.MaxViews)
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{
			Code:    responses.INTERNAL_SERVER_ERROR_CODE,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(201, responses.SuccessResponse{
		Data: codepic,
	})
}

func (h *Handler) GetCodePicByID(ctx *gin.Context) {
	id := ctx.Param("id")

	codepic, err := h.codePicsService.GetCodePicByID(id)
	if err != nil {
		ctx.JSON(500, responses.ErrorResponse{
			Code:    responses.INTERNAL_SERVER_ERROR_CODE,
			Message: err.Error(),
		})
		return
	}

	if codepic == nil {
		ctx.JSON(404, responses.ErrorResponse{
			Code:    CODEPIC_NOT_FOUND_ERROR_CODE,
			Message: "Codepic with id " + id + " not found",
		})
		return
	}

	ctx.JSON(200, responses.SuccessResponse{
		Data: codepic,
	})
}
