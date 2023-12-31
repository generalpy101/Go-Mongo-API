package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/generalpy101/Go-Mongo-API/routers"

	_ "github.com/generalpy101/Go-Mongo-API/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = ":8080"

// Tip: It is expected that in a golang project there is only 1 file mostly named main.go
// Tip: Open go projet as a workspace in VSCode to get intellisense and other features

func main() {
	fmt.Println("Go API With MongoDB")
	fmt.Println("Server starting on port", port)
	r := routes.Router()

	// Server docs at /docs
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started on port", port)
}
