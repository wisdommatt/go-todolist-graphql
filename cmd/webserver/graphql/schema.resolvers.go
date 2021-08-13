//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"errors"
	"fmt"
	"html"
	"time"

	"github.com/wisdommatt/go-todolist-graphql/cmd/webserver/graphql/generated"
	"github.com/wisdommatt/go-todolist-graphql/cmd/webserver/graphql/model"
	"github.com/wisdommatt/go-todolist-graphql/internal/todo"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

// CreateTodo is the mutation resolver to create new todo.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*todo.Todo, error) {
	newTodo := todo.Todo{
		Task:      html.EscapeString(input.Task),
		Status:    html.EscapeString(input.Status),
		TimeAdded: time.Now(),
	}
	if newTodo.Task == "" || newTodo.Status == "" {
		return nil, errors.New("All fields are required !")
	}
	err := r.TodoRepo.Save(&newTodo)
	if err != nil {
		return nil, errors.New("An error occured while adding new todo !")
	}
	return &newTodo, nil
}

// GetTodos is the query resolver to return all todos.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*todo.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
