package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	repo "mibo/repository"
	"net/http"
)

type WebHandler struct {
	MailRepo repo.MailRepo
}

func (m WebHandler) Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func (m WebHandler) MailList(c echo.Context) error {
	Results, err := m.MailRepo.List()
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error.html", map[string]interface{}{})
	}

	return c.Render(http.StatusOK, "history.html", echo.Map{
		"Results": Results,
	})
}