package main

import (
	"github.com/labstack/echo"
	"io"
	"mibo/driver"
	"mibo/handler"
	"mibo/pkg/mail"
	"mibo/repository/repoimpl"
	"mibo/router"
	"html/template"
)

const (
	DATABASE  = "mebo"
	MAIL_FROM = "" // your email
	MAIL_PASS = "" // your email password
	MAIL_SERVER = "smtp.gmail.com"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main()  {
	mail.Init(MAIL_FROM, MAIL_PASS, MAIL_SERVER)
	driver.ConnectMongoDB()
	e := echo.New()

	e.Static("/static", "./template/vendor")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./template/*.html")),
	}
	e.Renderer = renderer

	mailRepo := repoimpl.NewUserRepo(driver.Mongo.Client.Database(DATABASE))

	router.MailRouter(e, handler.MailHandler{MailRepo: mailRepo})
	router.WebRouter(e, handler.WebHandler{MailRepo: mailRepo})

	e.Logger.Fatal(e.Start(":3000"))
}
