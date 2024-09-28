package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arsmoriendy/opor/gql-srv/db"
	"github.com/arsmoriendy/opor/gql-srv/graph"
	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
	"github.com/arsmoriendy/opor/gql-srv/server"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	internal.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	loglvl.Init()
	db.InitPool()
	defer db.Pool.Close()

	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)

	rslvr := graph.New()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: rslvr}))

	router.Handle("/query", server.Auth(srv))

	if internal.IsDevMode() {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	} else {
		log.Printf("connect to http://%s:%s/query for GraphQL API", host, port)
	}

	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
