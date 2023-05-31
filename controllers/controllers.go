package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/Post-and-Play/edwiges/models"
	"github.com/Post-and-Play/edwiges/templates"
	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary      Healthcheck
// @Description  Route to health check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /healthz [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Readiness godoc
// @Summary      Healthcheck
// @Description  Route to readiness check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /readiness [get]
func Readiness(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// SendMail godoc
// @Summary      Sends a mail
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.MailRequest  true  "Mail Model"
// @Success      200  {object}  models.MailRequest
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /mail [post]
func SendMail(c *gin.Context) {
	var sender models.Sender
	var receiver models.Receiver

	sender.SenderMail = os.Getenv("SENDER_MAIL")
	sender.SenderPass = os.Getenv("SENDER_PASS")

	to := []string{
		receiver.ReceiverMail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", sender.SenderMail, sender.SenderPass, smtpHost)

	t, body := templates.BuildTemplate()

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender.SenderMail, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
