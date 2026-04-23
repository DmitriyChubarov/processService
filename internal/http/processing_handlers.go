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
	firstPhotoHeader, err := c.FormFile("first_photo")
	if err != nil {
		return fmt.Errorf("cant get first photo header: %w", err)
	}
	firstPhoto, err := firstPhotoHeader.Open()
	if err != nil {
		return fmt.Errorf("cant get first photo: %w", err)
	}
	defer firstPhoto.Close()
	firstPhotoData, err := io.ReadAll(firstPhoto)
	secondPhotoHeader, err := c.FormFile("second_photo")
	if err != nil {
		return fmt.Errorf("cant get second photo header: %w", err)
	}
	secondPhoto, err := secondPhotoHeader.Open()
	if err != nil {
		return fmt.Errorf("cant get second photo: %w", err)
	}
	defer secondPhoto.Close()
	secondPhotoData, err := io.ReadAll(secondPhoto)
	if err != nil {
		return fmt.Errorf("cant get second photo data: %w", err)
	}
	var process entity.Process
	if err := c.Bind(&process); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}
	err = p.processingLogic.CreateProcess(ctx, firstPhotoData, secondPhotoData, process)
	if err != nil {
		return fmt.Errorf("cant create process: %w", err)
	}
	return c.JSON(http.StatusOK, "200")
}
