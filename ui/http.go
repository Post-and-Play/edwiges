package ui

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Post-and-Play/edwiges/docs"
	"github.com/Post-and-Play/edwiges/services"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	handleRoutes(r)
	return r
}

func RunServer() {
	var port string

	host := os.Getenv("HOST")
	port = os.Getenv("PORT")

	done := services.MakeDoneSignal()

	server := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: Router(),
	}

	go func() {
		log.Printf("Server started at %s:%s", host, port)

		if err := server.ListenAndServe(); err != nil {
			log.Panicf("Error trying to start server: %+v", err)
		}
	}()

	<-done
	log.Println("Stopping server...")
}

func handleRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"

	for _, route := range healthCheck {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	for _, route := range swagg {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := r.Group("/api")

	routers := [][]Route{mail}

	for _, router := range routers {
		for _, route := range router {
			apiGroup.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
