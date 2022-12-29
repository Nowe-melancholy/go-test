package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"go-test/graph/model"
)

// CreateL is the resolver for the createL field.
func (r *mutationResolver) CreateL(ctx context.Context, input model.CreateLIdInput) (int, error) {
	panic(fmt.Errorf("not implemented: CreateL - createL"))
}

// CreateD is the resolver for the createD field.
func (r *mutationResolver) CreateD(ctx context.Context, input model.CreateDIdInput) (int, error) {
	panic(fmt.Errorf("not implemented: CreateD - createD"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
