// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-test/ent/middleware"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MiddlewareCreate is the builder for creating a Middleware entity.
type MiddlewareCreate struct {
	config
	mutation *MiddlewareMutation
	hooks    []Hook
}

// SetLID sets the "l_id" field.
func (mc *MiddlewareCreate) SetLID(i int) *MiddlewareCreate {
	mc.mutation.SetLID(i)
	return mc
}

// SetNillableLID sets the "l_id" field if the given value is not nil.
func (mc *MiddlewareCreate) SetNillableLID(i *int) *MiddlewareCreate {
	if i != nil {
		mc.SetLID(*i)
	}
	return mc
}

// SetDID sets the "d_id" field.
func (mc *MiddlewareCreate) SetDID(i int) *MiddlewareCreate {
	mc.mutation.SetDID(i)
	return mc
}

// SetNillableDID sets the "d_id" field if the given value is not nil.
func (mc *MiddlewareCreate) SetNillableDID(i *int) *MiddlewareCreate {
	if i != nil {
		mc.SetDID(*i)
	}
	return mc
}

// SetSysID sets the "sys_id" field.
func (mc *MiddlewareCreate) SetSysID(i int) *MiddlewareCreate {
	mc.mutation.SetSysID(i)
	return mc
}

// Mutation returns the MiddlewareMutation object of the builder.
func (mc *MiddlewareCreate) Mutation() *MiddlewareMutation {
	return mc.mutation
}

// Save creates the Middleware in the database.
func (mc *MiddlewareCreate) Save(ctx context.Context) (*Middleware, error) {
	var (
		err  error
		node *Middleware
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiddlewareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Middleware)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MiddlewareMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MiddlewareCreate) SaveX(ctx context.Context) *Middleware {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MiddlewareCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MiddlewareCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MiddlewareCreate) check() error {
	if _, ok := mc.mutation.SysID(); !ok {
		return &ValidationError{Name: "sys_id", err: errors.New(`ent: missing required field "Middleware.sys_id"`)}
	}
	return nil
}

func (mc *MiddlewareCreate) sqlSave(ctx context.Context) (*Middleware, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MiddlewareCreate) createSpec() (*Middleware, *sqlgraph.CreateSpec) {
	var (
		_node = &Middleware{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: middleware.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: middleware.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.LID(); ok {
		_spec.SetField(middleware.FieldLID, field.TypeInt, value)
		_node.LID = value
	}
	if value, ok := mc.mutation.DID(); ok {
		_spec.SetField(middleware.FieldDID, field.TypeInt, value)
		_node.DID = value
	}
	if value, ok := mc.mutation.SysID(); ok {
		_spec.SetField(middleware.FieldSysID, field.TypeInt, value)
		_node.SysID = value
	}
	return _node, _spec
}

// MiddlewareCreateBulk is the builder for creating many Middleware entities in bulk.
type MiddlewareCreateBulk struct {
	config
	builders []*MiddlewareCreate
}

// Save creates the Middleware entities in the database.
func (mcb *MiddlewareCreateBulk) Save(ctx context.Context) ([]*Middleware, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Middleware, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MiddlewareMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MiddlewareCreateBulk) SaveX(ctx context.Context) []*Middleware {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MiddlewareCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MiddlewareCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
