package main

import (
	"github.com/SemmiDev/go-graphql/internal/auth"
	database "github.com/SemmiDev/go-graphql/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SemmiDev/go-graphql/graph"
	"github.com/SemmiDev/go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}