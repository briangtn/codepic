package responses

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	INVALID_PAYLOAD_ERROR_CODE = "INVALID_PAYLOAD"
	INTERNAL_SERVER_ERROR_CODE = "INTERNAL_SERVER_ERROR"
)
