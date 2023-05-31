package ui

import (
	"net/http"

	"github.com/Post-and-Play/edwiges/controllers"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

var swagg = []Route{
	{
		"/swagger/*any",
		http.MethodGet,
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	},
}

var mail = []Route{
	{
		"/mail",
		http.MethodPost,
		controllers.SendMail,
	},
}
