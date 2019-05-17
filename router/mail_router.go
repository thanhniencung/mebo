package router

import (
	"github.com/labstack/echo"
	"mibo/handler"
)

func MailRouter(e *echo.Echo, handler handler.MailHandler) {
	e.GET("/", handler.Hello)
	e.POST("/mail-service/send", handler.SendEmail)
	e.GET("/mail-service/list", handler.MailList)
}
