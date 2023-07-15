// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	ent1 "entgo.io/contrib/entgql/internal/todo/ent"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/contrib/entgql/internal/todouuid/ent"
	"github.com/google/uuid"
)

func (r *categoryResolver) Types(ctx context.Context, obj *ent.Category) (*CategoryTypes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationResolver) ID(ctx context.Context, obj *ent1.Workspace) (uuid.UUID, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) BillProducts(ctx context.Context) ([]*ent.BillProduct, error) {
	return r.client.BillProduct.Query().All(ctx)
}

func (r *queryResolver) Categories(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.CategoryOrder, where *ent.CategoryWhereInput) (*ent.CategoryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGroupFilter(where.Filter),
		)
}

func (r *queryResolver) OneToMany(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *OneToManyOrder, where *OneToManyWhereInput) (*OneToManyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserFilter(where.Filter),
		)
}

func (r *todoResolver) Status(ctx context.Context, obj *ent.Todo) (todo.Status, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Username(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Friendships(ctx context.Context, obj *ent.User, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.FriendshipWhereInput) (*ent.FriendshipConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *createCategoryInputResolver) Types(ctx context.Context, obj *ent.CreateCategoryInput, data *CategoryTypesInput) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createTodoInputResolver) Status(ctx context.Context, obj *ent.CreateTodoInput, data todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *createUserInputResolver) Username(ctx context.Context, obj *ent.CreateUserInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) Status(ctx context.Context, obj *ent.TodoWhereInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusNeq(ctx context.Context, obj *ent.TodoWhereInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusIn(ctx context.Context, obj *ent.TodoWhereInput, data []todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusNotIn(ctx context.Context, obj *ent.TodoWhereInput, data []todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateCategoryInputResolver) Types(ctx context.Context, obj *ent.UpdateCategoryInput, data *CategoryTypesInput) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateTodoInputResolver) Status(ctx context.Context, obj *ent.UpdateTodoInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateUserInputResolver) Username(ctx context.Context, obj *ent.UpdateUserInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) Username(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) UsernameLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Organization returns OrganizationResolver implementation.
func (r *Resolver) Organization() OrganizationResolver { return &organizationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// CreateCategoryInput returns CreateCategoryInputResolver implementation.
func (r *Resolver) CreateCategoryInput() CreateCategoryInputResolver {
	return &createCategoryInputResolver{r}
}

// CreateTodoInput returns CreateTodoInputResolver implementation.
func (r *Resolver) CreateTodoInput() CreateTodoInputResolver { return &createTodoInputResolver{r} }

// CreateUserInput returns CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() CreateUserInputResolver { return &createUserInputResolver{r} }

// TodoWhereInput returns TodoWhereInputResolver implementation.
func (r *Resolver) TodoWhereInput() TodoWhereInputResolver { return &todoWhereInputResolver{r} }

// UpdateCategoryInput returns UpdateCategoryInputResolver implementation.
func (r *Resolver) UpdateCategoryInput() UpdateCategoryInputResolver {
	return &updateCategoryInputResolver{r}
}

// UpdateTodoInput returns UpdateTodoInputResolver implementation.
func (r *Resolver) UpdateTodoInput() UpdateTodoInputResolver { return &updateTodoInputResolver{r} }

// UpdateUserInput returns UpdateUserInputResolver implementation.
func (r *Resolver) UpdateUserInput() UpdateUserInputResolver { return &updateUserInputResolver{r} }

// UserWhereInput returns UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() UserWhereInputResolver { return &userWhereInputResolver{r} }

type categoryResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type createCategoryInputResolver struct{ *Resolver }
type createTodoInputResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
type todoWhereInputResolver struct{ *Resolver }
type updateCategoryInputResolver struct{ *Resolver }
type updateTodoInputResolver struct{ *Resolver }
type updateUserInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
