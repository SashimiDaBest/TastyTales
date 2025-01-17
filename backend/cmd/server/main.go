package main

import (
	_ "database/sql"
	_ "fmt"
	"html/template"
	_ "log"
	"net/http"
	_ "os"

	_ "github.com/SashimiDaBest/TastyTales/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)
  
func main() {
	/*
	// Load environment variables from .env file
	config.LoadEnv()

	// Retrieve required environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "3000" // Fallback to default port
	}

	// PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verify the connection to the database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database!")

	// Load environment variables from .env file
	config.LoadEnv()
	*/
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
	/*
	// Start the server
	log.Printf("Starting server on port %s...", appPort)
	err = http.ListenAndServe(":"+appPort, r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	*/
}

// Serve the main page
func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// Serve partial content for HTMX requests
func servePartial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := template.Must(template.ParseFiles("templates/partial.html")).Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}


/*
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
*/
