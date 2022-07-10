package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/guicostaarantes/go-auth/graph/generated"
	users_commands "github.com/guicostaarantes/go-auth/modules/users/commands"
	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
)

func (r *mutationResolver) AuthenticateUser(ctx context.Context, input users_commands.AuthenticateUserInput) (string, error) {
	return r.Commands.AuthenticateUser.ExecuteCommand(input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input users_commands.CreateUserInput) (bool, error) {
	return r.Commands.CreateUser.ExecuteCommand(input)
}

func (r *queryResolver) MyUser(ctx context.Context) (*users_models.User, error) {
	userID := ctx.Value("userID").(string)

	return r.Queries.UserByID.ExecuteQuery(userID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
