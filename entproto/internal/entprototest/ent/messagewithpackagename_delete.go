// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithPackageNameDelete is the builder for deleting a MessageWithPackageName entity.
type MessageWithPackageNameDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithPackageNameMutation
}

// Where appends a list predicates to the MessageWithPackageNameDelete builder.
func (mwpnd *MessageWithPackageNameDelete) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameDelete {
	mwpnd.mutation.Where(ps...)
	return mwpnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwpnd *MessageWithPackageNameDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithPackageNameMutation](ctx, mwpnd.sqlExec, mwpnd.mutation, mwpnd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpnd *MessageWithPackageNameDelete) ExecX(ctx context.Context) int {
	n, err := mwpnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwpnd *MessageWithPackageNameDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(messagewithpackagename.Table, sqlgraph.NewFieldSpec(messagewithpackagename.FieldID, field.TypeInt))
	if ps := mwpnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwpnd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwpnd.mutation.done = true
	return affected, err
}

// MessageWithPackageNameDeleteOne is the builder for deleting a single MessageWithPackageName entity.
type MessageWithPackageNameDeleteOne struct {
	mwpnd *MessageWithPackageNameDelete
}

// Where appends a list predicates to the MessageWithPackageNameDelete builder.
func (mwpndo *MessageWithPackageNameDeleteOne) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameDeleteOne {
	mwpndo.mwpnd.mutation.Where(ps...)
	return mwpndo
}

// Exec executes the deletion query.
func (mwpndo *MessageWithPackageNameDeleteOne) Exec(ctx context.Context) error {
	n, err := mwpndo.mwpnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithpackagename.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpndo *MessageWithPackageNameDeleteOne) ExecX(ctx context.Context) {
	if err := mwpndo.Exec(ctx); err != nil {
		panic(err)
	}
}
