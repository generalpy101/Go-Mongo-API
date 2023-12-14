package routes

import (
	"log"
	"net/http"

	"github.com/generalpy101/Go-Mongo-API/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.HandleFunc("/api/movies", controllers.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkMovieAsWatchedController).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteMovieController).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMoviesController).Methods("DELETE")

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log information about the request
		log.Printf("Request: %s %s", r.Method, r.RequestURI)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
