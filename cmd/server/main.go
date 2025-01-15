package main

import (
	"html/template"
	"net/http"

	"fmt"
	"config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load environment variables from .env file
	config.LoadEnv()

	// Retrieve environment variables
	port := config.GetEnv("APP_PORT", "3000")
	dbUser := config.GetEnv("DB_USER", "default_user")
	dbPassword := config.GetEnv("DB_PASSWORD", "default_password")

	fmt.Println("Port:", port)
	fmt.Println("Database User:", dbUser)
	fmt.Println("Database Password:", dbPassword)

	// Create a new router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Define routes
	r.Get("/", serveIndex)
	r.Get("/partial", servePartial)

	// Start server
	http.ListenAndServe(":8080", r)
}

// Serve the main page
func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// Serve partial content for HTMX requests
func servePartial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template.Must(template.ParseFiles("templates/partial.html")).Execute(w, nil)
}
