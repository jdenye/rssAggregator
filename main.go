package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


// func handleReadiness(sen http.ResponseWriter, res *http.Request)  {
	
// }
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
	portNumber := os.Getenv("PORT")

	if portNumber == "" {
		log.Fatal("PORT number not found in the environment")
	}
	fmt.Println("Port: ", portNumber)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https//*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthy", handlerReadiness)

	router.Mount("/v1", v1Router)
	
	serve := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Printf("server running on port %v", portNumber)
	err = serve.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
