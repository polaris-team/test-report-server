package resolvers

import (
	"context"
	"github.com/polaris-team/test-report-server/src/graphql/gqlgen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func ( mutationResolver) GetJsAPISign(ctx context.Context, input interface{}) (*interface{}, error) {
	panic("implement me")
}

func ( mutationResolver) CreateProject(ctx context.Context, input interface{}) (*interface{}, error) {
	panic("implement me")
}

func ( mutationResolver) UpdateProject(ctx context.Context, input interface{}) (*interface{}, error) {
	panic("implement me")
}

func ( mutationResolver) DeleteProject(ctx context.Context, input interface{}) (*interface{}, error) {
	panic("implement me")
}

type queryResolver struct{ *Resolver }

func ( queryResolver) Users(ctx context.Context) ([]*interface{}, error) {
	panic("implement me")
}

func ( queryResolver) Projects(ctx context.Context) ([]*interface{}, error) {
	panic("implement me")
}













