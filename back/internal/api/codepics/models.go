package codepics

type Payload struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Language string `json:"language" binding:"required"`
	MaxViews int    `json:"max_views" binding:"gte=0"`
}

const (
	CODEPIC_NOT_FOUND_ERROR_CODE = "CODEPIC_NOT_FOUND"
)
