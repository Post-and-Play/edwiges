package ui

import (
	"net/http"

	"github.com/Post-and-Play/edwiges/controllers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var healthCheck = []Route{
	{
		"/healthz",
		http.MethodGet,
		controllers.Health,
	},
	{
		"/readiness",
		http.MethodGet,
		controllers.Readiness,
	},
}

var mail = []Route{
	{
		"/mail",
		http.MethodPost,
		controllers.SendMail,
	},
}
