package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"testgql/graph/generated"
	"testgql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input1 string, input2 *model.Inputset2) (*model.Todo, error) {
	var output model.Todo
	fmt.Println(input1)
	fmt.Println(*input2)
	datatest := *input2

	fmt.Println(datatest.Data1)
	return &output, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
