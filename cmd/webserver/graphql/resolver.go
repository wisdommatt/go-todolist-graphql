package graphql

import "github.com/wisdommatt/go-todolist-graphql/internal/todo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoRepo todo.Repository
}
