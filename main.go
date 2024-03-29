package main

import (
	"github.com/dyldawg/golf/handlers"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.Use(logRequests)

	courses := router.PathPrefix("/courses").Subrouter()
	courses.HandleFunc("", handlers.CoursesHandler).Methods(http.MethodGet)
	courses.HandleFunc(`/{id:\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b}`, handlers.GetCourseHandler).Methods(http.MethodGet)
	courses.HandleFunc("", handlers.NewCourseHandler).Methods(http.MethodPost)
	courses.HandleFunc("", handlers.UpdateCourseHandler).Methods(http.MethodPatch)
	courses.HandleFunc(`/{id:\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b}`, handlers.DeleteCourseHandler).Methods(http.MethodDelete)

	results := router.PathPrefix("/results").Subrouter()
	results.HandleFunc("", handlers.ResultsHandler).Methods(http.MethodGet)
	results.HandleFunc("", handlers.NewResultHandler).Methods(http.MethodPost)
	results.HandleFunc(`/{id:\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b}`, handlers.GetResultHandler).Methods(http.MethodGet)

	players := router.PathPrefix("/players").Subrouter()
	players.HandleFunc("", handlers.PlayersHandler).Methods(http.MethodGet)
	players.HandleFunc("", handlers.NewPlayerHandler).Methods(http.MethodPost)
	players.HandleFunc(`/{id:\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b}`, handlers.GetPlayerHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":3000", gorillaHandlers.CORS(gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), gorillaHandlers.AllowedOrigins([]string{"*"}))(router)))

}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

