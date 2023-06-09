package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"example/hello/graph/model"
	"fmt"
)

// CreatePart is the resolver for the CreatePart field.
func (r *mutationResolver) CreatePart(ctx context.Context, part *model.NewPart) (*model.Part, error) {
	panic(fmt.Errorf("not implemented: CreatePart - CreatePart"))
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - CreateUser"))
}

// CreateBear is the resolver for the CreateBear field.
func (r *mutationResolver) CreateBear(ctx context.Context, bear *model.NewTeddyBear) (*model.TeddyBear, error) {
	panic(fmt.Errorf("not implemented: CreateBear - CreateBear"))
}

// SellBear is the resolver for the SellBear field.
func (r *mutationResolver) SellBear(ctx context.Context, ownedBear *model.ModifyTeddyBear, newOwner *model.ModifyUser) (*model.TeddyBear, error) {
	panic(fmt.Errorf("not implemented: SellBear - SellBear"))
}

// ModifyPart is the resolver for the ModifyPart field.
func (r *mutationResolver) ModifyPart(ctx context.Context, part *model.ModifyPart) (*model.Part, error) {
	panic(fmt.Errorf("not implemented: ModifyPart - ModifyPart"))
}

// ModifyUser is the resolver for the ModifyUser field.
func (r *mutationResolver) ModifyUser(ctx context.Context, user *model.ModifyUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: ModifyUser - ModifyUser"))
}

// ModifyTeddybear is the resolver for the ModifyTeddybear field.
func (r *mutationResolver) ModifyTeddybear(ctx context.Context, ownedBear *model.ModifyTeddyBear) (*model.TeddyBear, error) {
	panic(fmt.Errorf("not implemented: ModifyTeddybear - ModifyTeddybear"))
}

// GetOwnedBearsList is the resolver for the GetOwnedBearsList field.
func (r *queryResolver) GetOwnedBearsList(ctx context.Context, user *model.ModifyUser) ([]*model.TeddyBear, error) {
	panic(fmt.Errorf("not implemented: GetOwnedBearsList - GetOwnedBearsList"))
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - GetUser"))
}

// GetOwnedParts is the resolver for the GetOwnedParts field.
func (r *queryResolver) GetOwnedParts(ctx context.Context, user *model.ModifyUser) ([]*model.Part, error) {
	panic(fmt.Errorf("not implemented: GetOwnedParts - GetOwnedParts"))
}

// GetOwnedPartsByType is the resolver for the GetOwnedPartsByType field.
func (r *queryResolver) GetOwnedPartsByType(ctx context.Context, user *model.ModifyUser, part *model.PartType) ([]*model.Part, error) {
	panic(fmt.Errorf("not implemented: GetOwnedPartsByType - GetOwnedPartsByType"))
}

// GetBearInfo is the resolver for the GetBearInfo field.
func (r *queryResolver) GetBearInfo(ctx context.Context, bear *model.ModifyTeddyBear) ([]*model.Part, error) {
	panic(fmt.Errorf("not implemented: GetBearInfo - GetBearInfo"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
