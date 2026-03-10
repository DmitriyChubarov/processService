package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DmitriyChubarov/processing/internal/entity"
	"github.com/labstack/echo/v4"
)

type processingHandlers struct {
	processingLogic entity.ProcessLogic
}

func NewProcessingHandlers(processingLogic entity.ProcessLogic) *processingHandlers {
	return &processingHandlers{processingLogic: processingLogic}
}

func (p *processingHandlers) Register(e *echo.Echo) {
	e.POST("/start-processing/", p.StartProcessing)
}

func (p *processingHandlers) StartProcessing(c echo.Context) (err error) {
	ctx := c.Request().Context()
	photoHeader, err := c.FormFile("original_photo")
	if err != nil {
		return fmt.Errorf("cant get original photo header")
	}
	photo, err := photoHeader.Open()
	if err != nil {
		return fmt.Errorf("cant get original photo")
	}
	defer photo.Close()
	photoData, err := io.ReadAll(photo)
	if err != nil {
		return fmt.Errorf("cant get original photo data")
	}
	result, err := p.processingLogic.CreateProcess(ctx, photoData)
	if err != nil {
		return fmt.Errorf("cant create process: %w", err)
	}
	return c.JSON(http.StatusOK, result)
}
