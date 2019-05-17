package router

import (
	"github.com/labstack/echo"
	"mibo/handler"
)

func WebRouter(e *echo.Echo, handler handler.WebHandler) {
	e.GET("/", handler.Home)
	e.GET("/history", handler.MailList)
}