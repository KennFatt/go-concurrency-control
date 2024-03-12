package handler

import "time"

type (
	CommonResponse struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
	}

	SuccessResponse struct {
		CommonResponse
		Data any `json:"data"`
	}

	ErrorResponse struct {
		CommonResponse
		Reason any `json:"reason"`
	}
)

func NewResponseSuccessData(data any) *SuccessResponse {
	return &SuccessResponse{
		CommonResponse: CommonResponse{Status: "success", Timestamp: time.Now().Format(time.RFC3339Nano)},
		Data:           data,
	}
}

func NewResponseError(err error) *ErrorResponse {
	return &ErrorResponse{
		CommonResponse: CommonResponse{
			Status:    "error",
			Timestamp: time.Now().Format(time.RFC3339Nano),
		},
		Reason: err.Error(),
	}
}
