// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/dependsonskipped"
	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DependsOnSkippedCreate is the builder for creating a DependsOnSkipped entity.
type DependsOnSkippedCreate struct {
	config
	mutation *DependsOnSkippedMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dosc *DependsOnSkippedCreate) SetName(s string) *DependsOnSkippedCreate {
	dosc.mutation.SetName(s)
	return dosc
}

// AddSkippedIDs adds the "skipped" edge to the ImplicitSkippedMessage entity by IDs.
func (dosc *DependsOnSkippedCreate) AddSkippedIDs(ids ...int) *DependsOnSkippedCreate {
	dosc.mutation.AddSkippedIDs(ids...)
	return dosc
}

// AddSkipped adds the "skipped" edges to the ImplicitSkippedMessage entity.
func (dosc *DependsOnSkippedCreate) AddSkipped(i ...*ImplicitSkippedMessage) *DependsOnSkippedCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dosc.AddSkippedIDs(ids...)
}

// Mutation returns the DependsOnSkippedMutation object of the builder.
func (dosc *DependsOnSkippedCreate) Mutation() *DependsOnSkippedMutation {
	return dosc.mutation
}

// Save creates the DependsOnSkipped in the database.
func (dosc *DependsOnSkippedCreate) Save(ctx context.Context) (*DependsOnSkipped, error) {
	return withHooks[*DependsOnSkipped, DependsOnSkippedMutation](ctx, dosc.sqlSave, dosc.mutation, dosc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dosc *DependsOnSkippedCreate) SaveX(ctx context.Context) *DependsOnSkipped {
	v, err := dosc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dosc *DependsOnSkippedCreate) Exec(ctx context.Context) error {
	_, err := dosc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dosc *DependsOnSkippedCreate) ExecX(ctx context.Context) {
	if err := dosc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dosc *DependsOnSkippedCreate) check() error {
	if _, ok := dosc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DependsOnSkipped.name"`)}
	}
	return nil
}

func (dosc *DependsOnSkippedCreate) sqlSave(ctx context.Context) (*DependsOnSkipped, error) {
	if err := dosc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dosc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dosc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dosc.mutation.id = &_node.ID
	dosc.mutation.done = true
	return _node, nil
}

func (dosc *DependsOnSkippedCreate) createSpec() (*DependsOnSkipped, *sqlgraph.CreateSpec) {
	var (
		_node = &DependsOnSkipped{config: dosc.config}
		_spec = sqlgraph.NewCreateSpec(dependsonskipped.Table, sqlgraph.NewFieldSpec(dependsonskipped.FieldID, field.TypeInt))
	)
	if value, ok := dosc.mutation.Name(); ok {
		_spec.SetField(dependsonskipped.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := dosc.mutation.SkippedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(implicitskippedmessage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DependsOnSkippedCreateBulk is the builder for creating many DependsOnSkipped entities in bulk.
type DependsOnSkippedCreateBulk struct {
	config
	builders []*DependsOnSkippedCreate
}

// Save creates the DependsOnSkipped entities in the database.
func (doscb *DependsOnSkippedCreateBulk) Save(ctx context.Context) ([]*DependsOnSkipped, error) {
	specs := make([]*sqlgraph.CreateSpec, len(doscb.builders))
	nodes := make([]*DependsOnSkipped, len(doscb.builders))
	mutators := make([]Mutator, len(doscb.builders))
	for i := range doscb.builders {
		func(i int, root context.Context) {
			builder := doscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DependsOnSkippedMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, doscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, doscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, doscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (doscb *DependsOnSkippedCreateBulk) SaveX(ctx context.Context) []*DependsOnSkipped {
	v, err := doscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (doscb *DependsOnSkippedCreateBulk) Exec(ctx context.Context) error {
	_, err := doscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (doscb *DependsOnSkippedCreateBulk) ExecX(ctx context.Context) {
	if err := doscb.Exec(ctx); err != nil {
		panic(err)
	}
}
