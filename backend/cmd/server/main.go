package main

import (
	"html/template"
	"net/http"

	"fmt"
	"github.com/SashimiDaBest/TastyTales/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "database/sql"
	// "github.com/lib/pq"
)
// http://localhost:8080/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
  )
  
func main() {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + 
	// "password=%s dbname=%s sslmode=disable", 
	// host, port, user, password, dbname)

	/*
	host - The host to connect to
	port - The port to bind to
	user - The user to sign in as
	password - The user’s password
	dbname - The name of the database to connect to
	sslmode - Whether or not to use SSL
	*/


	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()


	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Successfully connected to database!")

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
