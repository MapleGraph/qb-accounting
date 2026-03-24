package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response is the standard API response structure
type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Error      []string    `json:"error"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}

// SuccessResponse sends a success response with data
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:     true,
		StatusCode: statusCode,
		Error:      nil,
		Data:       data,
		Message:    message,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	var errMsgs []string
	if err != nil {
		errMsgs = []string{err.Error()}
	}

	c.JSON(statusCode, Response{
		Status:     false,
		StatusCode: statusCode,
		Error:      errMsgs,
		Data:       nil,
		Message:    message,
	})
}

// ValidationError creates a response for validation errors
func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be a valid email", err.Field()))
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be at least %s", err.Field(), err.Param()))
		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be at most %s", err.Field(), err.Param()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status:     false,
		StatusCode: http.StatusBadRequest,
		Error:      errMsgs,
		Data:       nil,
		Message:    "Validation error",
	}
}

// GeneralError creates a general error response
func GeneralError(err error, message string) Response {
	return Response{
		Status:     false,
		StatusCode: http.StatusInternalServerError,
		Error:      []string{err.Error()},
		Data:       nil,
		Message:    message,
	}
}
