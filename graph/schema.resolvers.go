package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql/database"
	"go-graphql/graph/generated"
	"go-graphql/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateCreator(ctx context.Context, input *model.NewCreator) (*model.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateArt(ctx context.Context, input *model.NewArt) (*model.Art, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Creator(ctx context.Context, id string) (*model.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Creators(ctx context.Context) ([]*model.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Art(ctx context.Context, id string) (*model.Art, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Arts(ctx context.Context) ([]*model.Art, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

//func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
//	return db.Save(input), nil
//}
//func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
//	return db.FindByID(id), nil
//}
//func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
//	return db.All(), nil
//}
