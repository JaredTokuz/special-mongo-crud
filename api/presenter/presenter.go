package presenter

import (
	"github.com/gofiber/fiber/v2"
)

// Default Success
func SuccessResponse(data ...interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
// Default ErrorResponse
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

// ! Custom Responses below ...

// * Example: BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
func FindSuccessResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

