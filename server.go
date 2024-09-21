package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arsmoriendy/opor/gql-srv/graph"
	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	if internal.IsDevMode() {
		log.SetPrefix("[DEV] ")
	} else {
		log.SetPrefix("[PROD] ")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)

	rslvr := graph.Resolver{}
	rslvr.GetPorts()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &rslvr}))

	router.Handle("/query", srv)

	if internal.IsDevMode() {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	} else {
		log.Printf("connect to http://localhost:%s/query for GraphQL API", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
