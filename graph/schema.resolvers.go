package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql/database"
	"go-graphql/graph/generated"
	"go-graphql/graph/model"
)

func (r *mutationResolver) CreateCreator(ctx context.Context, input *model.NewCreator) (*model.Creator, error) {
	return db.SaveCreator(input), nil
}

func (r *mutationResolver) CreateArt(ctx context.Context, input *model.NewArt) (*model.Art, error) {
	return db.SaveArt(input), nil
}

func (r *queryResolver) CreatorByID(ctx context.Context, id string) (*model.Creator, error) {
	return db.FindCreatorByID(id), nil
}

func (r *queryResolver) Creators(ctx context.Context) ([]*model.Creator, error) {
	return db.FindAllCreators(), nil
}

func (r *queryResolver) ArtByID(ctx context.Context, id string) (*model.Art, error) {
	return db.FindArtByID(id), nil
}

func (r *queryResolver) Arts(ctx context.Context) ([]*model.Art, error) {
	return db.FindAllArts(), nil
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
var db = database.Connect()
