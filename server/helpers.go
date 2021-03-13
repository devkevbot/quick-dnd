package main

import (
	"github.com/labstack/echo/v4"
)

type standardResponse struct {
	Event      string      `json:"event"`
	URI        string      `json:"request_uri"`
	HTTPMethod string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// sendJSONResponse is a helper method which will return a
// JSON-formatted HTTP response with status = `statusCode`.
//
// This function should be called when returning either an error or an
// OK message inside an HTTP handler.
func sendJSONResponse(c echo.Context, statusCode int, event, message string, data interface{}) error {
	resp := standardResponse{
		Event:      event,
		URI:        c.Request().RequestURI,
		HTTPMethod: c.Request().Method,
		Message:    message,
		Data:       data,
	}

	return c.JSON(statusCode, resp)
}
