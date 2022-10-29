package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Status, resp)
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err.Error(),
	}
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Error:   err.Error(),
	}
}

func DataConflict(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: "Diplicate data",
		Error:   err.Error(),
	}
}

func DataNotFound(message string, err error) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   err.Error(),
	}
}

func Unauthorized(message string, err error) *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Error:   err.Error(),
	}
}

func DeleteSuccess(message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
	}
}
