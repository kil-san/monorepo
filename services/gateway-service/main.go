package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kil-san/micro-serv/gateway-service/graph"
	"github.com/kil-san/micro-serv/gateway-service/graph/generated"
	gatewayMiddleware "github.com/kil-san/micro-serv/gateway-service/middleware"
	log "unknwon.dev/clog/v2"
)

const port = "8000"

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(gatewayMiddleware.CORS())

	err := log.NewConsole()
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{},
		},
	))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Trace("connect to http://localhost:%s/ for GraphQL playground", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("%+v\n", err)
	}

}
