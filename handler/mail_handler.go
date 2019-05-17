package handler

import (
	"github.com/labstack/echo"
	"mibo/model"
	"mibo/pkg/mail"
	"net/http"
	"time"

	repo "mibo/repository"
)

type MailHandler struct {
	MailRepo repo.MailRepo
}

func (m *MailHandler) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Hello!",
	})
}

// Handle logic API send email
func (m *MailHandler) SendEmail(c echo.Context) error {
	data := model.Mail{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	mail := mail.New()
	templateData := model.Template{Name: data.Name}

	err := mail.Send(data.To, data.Subject, "./template/mail_template.html", templateData)
	if err == nil {
		err = m.MailRepo.Save(model.History{Name: data.Name, Email: data.To, Date: time.Now()})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, model.Response{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}

	return c.JSON(http.StatusInternalServerError, model.Response{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	})
}

// Get all mail sent
func (m *MailHandler) MailList(c echo.Context) error {
	results, err := m.MailRepo.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: results,
	})
}
