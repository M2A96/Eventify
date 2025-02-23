package helloworld

import (
	"net/http"

	"github.com/M2A96/Eventify.git/internal/core/ports/input"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	service input.HelloworldServicePort
}

func NewHTTPHandler(service input.HelloworldServicePort) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) GetHello(c echo.Context) error {
	name := c.Param("name")
	response, err := h.service.Greet(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}
