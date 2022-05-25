package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/romshark/gqlgen-2207-reproduction/graph/generated"
	"github.com/romshark/gqlgen-2207-reproduction/graph/model"
)

func (r *queryResolver) E(ctx context.Context, e1 model.E1, e2 model.E2, e3 model.E3) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
