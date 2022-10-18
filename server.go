package main

import (
	"fmt"
	"log"
	"net/http"

	model "server/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home. ")
	fmt.Fprintf(w, "\n For Api: Go to /api/Kittys for api")
	fmt.Println("Endpoint Hit: Endpoint hit (homePage)")
}

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/api/kittys", model.Kittys).Methods("GET")
	r.HandleFunc("/", homePage).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handler := c.Handler(r)
	fmt.Printf("Starting server at port 8080\n")

	log.Fatal((http.ListenAndServe(":8080", handler)))
	http.Handle("/", r)
}

func main() {

	handleRequests()
}
