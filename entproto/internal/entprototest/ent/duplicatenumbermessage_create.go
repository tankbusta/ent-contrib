// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/duplicatenumbermessage"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DuplicateNumberMessageCreate is the builder for creating a DuplicateNumberMessage entity.
type DuplicateNumberMessageCreate struct {
	config
	mutation *DuplicateNumberMessageMutation
	hooks    []Hook
}

// SetHello sets the "hello" field.
func (dnmc *DuplicateNumberMessageCreate) SetHello(s string) *DuplicateNumberMessageCreate {
	dnmc.mutation.SetHello(s)
	return dnmc
}

// SetWorld sets the "world" field.
func (dnmc *DuplicateNumberMessageCreate) SetWorld(s string) *DuplicateNumberMessageCreate {
	dnmc.mutation.SetWorld(s)
	return dnmc
}

// Mutation returns the DuplicateNumberMessageMutation object of the builder.
func (dnmc *DuplicateNumberMessageCreate) Mutation() *DuplicateNumberMessageMutation {
	return dnmc.mutation
}

// Save creates the DuplicateNumberMessage in the database.
func (dnmc *DuplicateNumberMessageCreate) Save(ctx context.Context) (*DuplicateNumberMessage, error) {
	return withHooks(ctx, dnmc.sqlSave, dnmc.mutation, dnmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dnmc *DuplicateNumberMessageCreate) SaveX(ctx context.Context) *DuplicateNumberMessage {
	v, err := dnmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dnmc *DuplicateNumberMessageCreate) Exec(ctx context.Context) error {
	_, err := dnmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmc *DuplicateNumberMessageCreate) ExecX(ctx context.Context) {
	if err := dnmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dnmc *DuplicateNumberMessageCreate) check() error {
	if _, ok := dnmc.mutation.Hello(); !ok {
		return &ValidationError{Name: "hello", err: errors.New(`ent: missing required field "DuplicateNumberMessage.hello"`)}
	}
	if _, ok := dnmc.mutation.World(); !ok {
		return &ValidationError{Name: "world", err: errors.New(`ent: missing required field "DuplicateNumberMessage.world"`)}
	}
	return nil
}

func (dnmc *DuplicateNumberMessageCreate) sqlSave(ctx context.Context) (*DuplicateNumberMessage, error) {
	if err := dnmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dnmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dnmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dnmc.mutation.id = &_node.ID
	dnmc.mutation.done = true
	return _node, nil
}

func (dnmc *DuplicateNumberMessageCreate) createSpec() (*DuplicateNumberMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &DuplicateNumberMessage{config: dnmc.config}
		_spec = sqlgraph.NewCreateSpec(duplicatenumbermessage.Table, sqlgraph.NewFieldSpec(duplicatenumbermessage.FieldID, field.TypeInt))
	)
	if value, ok := dnmc.mutation.Hello(); ok {
		_spec.SetField(duplicatenumbermessage.FieldHello, field.TypeString, value)
		_node.Hello = value
	}
	if value, ok := dnmc.mutation.World(); ok {
		_spec.SetField(duplicatenumbermessage.FieldWorld, field.TypeString, value)
		_node.World = value
	}
	return _node, _spec
}

// DuplicateNumberMessageCreateBulk is the builder for creating many DuplicateNumberMessage entities in bulk.
type DuplicateNumberMessageCreateBulk struct {
	config
	err      error
	builders []*DuplicateNumberMessageCreate
}

// Save creates the DuplicateNumberMessage entities in the database.
func (dnmcb *DuplicateNumberMessageCreateBulk) Save(ctx context.Context) ([]*DuplicateNumberMessage, error) {
	if dnmcb.err != nil {
		return nil, dnmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dnmcb.builders))
	nodes := make([]*DuplicateNumberMessage, len(dnmcb.builders))
	mutators := make([]Mutator, len(dnmcb.builders))
	for i := range dnmcb.builders {
		func(i int, root context.Context) {
			builder := dnmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DuplicateNumberMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, dnmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dnmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dnmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dnmcb *DuplicateNumberMessageCreateBulk) SaveX(ctx context.Context) []*DuplicateNumberMessage {
	v, err := dnmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dnmcb *DuplicateNumberMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := dnmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmcb *DuplicateNumberMessageCreateBulk) ExecX(ctx context.Context) {
	if err := dnmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
