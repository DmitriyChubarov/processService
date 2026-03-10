package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, map[string]interface{}{
			"error": he.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}
