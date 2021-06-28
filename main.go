package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/andreasatle/bookings/config"
	"github.com/andreasatle/bookings/handlers"
	"github.com/andreasatle/bookings/render"
)

const portNumber = ":8080"

var app *config.AppConfig

func main() {

	// Initialize AppConfig struct
	app = &config.AppConfig{
		UseCache:     true,
		InProduction: false,
	}

	// Setup a session in the AppConfig struct
	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = false

	handlers.NewHandlers(handlers.NewRepo(app))

	// Create Template Cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("error creating template: %v\n", err)
	}
	app.TemplateCache = tc

	// Set the App Config struct to a module struct in render.
	render.NewTemplates(app)

	log.Printf("Starting application of port %s.", portNumber)

	// Create a new server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(app),
	}

	// Start the listener and serve
	log.Fatal(srv.ListenAndServe())
}
