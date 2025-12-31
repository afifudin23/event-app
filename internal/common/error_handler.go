package common

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// CUSTOM ERRORS
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, ErrorResponse[any](
			ErrorSchema{
				Code:    appErr.Code,
				Message: appErr.Message,
				Details: appErr.Details,
			},
		))
		c.Errors = nil
		return
	}

	// SERVER ERRORS
	log.Print(err)
	c.JSON(http.StatusInternalServerError, InternalServerError())
}
