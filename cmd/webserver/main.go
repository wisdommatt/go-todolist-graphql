package main

import (
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wisdommatt/go-todolist-graphql/cmd/webserver/graphql"
	"github.com/wisdommatt/go-todolist-graphql/cmd/webserver/graphql/generated"
	"github.com/wisdommatt/go-todolist-graphql/internal/todo"
)

const defaultPort = "8080"

func main() {
	dbConn, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/todo-list?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbConn.AutoMigrate(&todo.Todo{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
