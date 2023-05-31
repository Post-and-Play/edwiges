package main

import (
	"github.com/Post-and-Play/edwiges/ui"
	"github.com/joho/godotenv"
)

//v0.1

func main() {
	godotenv.Load(".env")
	ui.RunServer()
}
