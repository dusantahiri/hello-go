package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dusantahiri/hello-go/pkg/config"
	"github.com/dusantahiri/hello-go/pkg/handlers"
	"github.com/dusantahiri/hello-go/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Staring application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
