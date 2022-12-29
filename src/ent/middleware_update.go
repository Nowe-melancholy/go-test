// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-test/ent/middleware"
	"go-test/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MiddlewareUpdate is the builder for updating Middleware entities.
type MiddlewareUpdate struct {
	config
	hooks    []Hook
	mutation *MiddlewareMutation
}

// Where appends a list predicates to the MiddlewareUpdate builder.
func (mu *MiddlewareUpdate) Where(ps ...predicate.Middleware) *MiddlewareUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetLID sets the "l_id" field.
func (mu *MiddlewareUpdate) SetLID(s string) *MiddlewareUpdate {
	mu.mutation.SetLID(s)
	return mu
}

// SetDID sets the "d_id" field.
func (mu *MiddlewareUpdate) SetDID(s string) *MiddlewareUpdate {
	mu.mutation.SetDID(s)
	return mu
}

// SetSysID sets the "sys_id" field.
func (mu *MiddlewareUpdate) SetSysID(s string) *MiddlewareUpdate {
	mu.mutation.SetSysID(s)
	return mu
}

// Mutation returns the MiddlewareMutation object of the builder.
func (mu *MiddlewareUpdate) Mutation() *MiddlewareMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MiddlewareUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiddlewareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MiddlewareUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MiddlewareUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MiddlewareUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MiddlewareUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   middleware.Table,
			Columns: middleware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: middleware.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.LID(); ok {
		_spec.SetField(middleware.FieldLID, field.TypeString, value)
	}
	if value, ok := mu.mutation.DID(); ok {
		_spec.SetField(middleware.FieldDID, field.TypeString, value)
	}
	if value, ok := mu.mutation.SysID(); ok {
		_spec.SetField(middleware.FieldSysID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{middleware.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MiddlewareUpdateOne is the builder for updating a single Middleware entity.
type MiddlewareUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MiddlewareMutation
}

// SetLID sets the "l_id" field.
func (muo *MiddlewareUpdateOne) SetLID(s string) *MiddlewareUpdateOne {
	muo.mutation.SetLID(s)
	return muo
}

// SetDID sets the "d_id" field.
func (muo *MiddlewareUpdateOne) SetDID(s string) *MiddlewareUpdateOne {
	muo.mutation.SetDID(s)
	return muo
}

// SetSysID sets the "sys_id" field.
func (muo *MiddlewareUpdateOne) SetSysID(s string) *MiddlewareUpdateOne {
	muo.mutation.SetSysID(s)
	return muo
}

// Mutation returns the MiddlewareMutation object of the builder.
func (muo *MiddlewareUpdateOne) Mutation() *MiddlewareMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MiddlewareUpdateOne) Select(field string, fields ...string) *MiddlewareUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Middleware entity.
func (muo *MiddlewareUpdateOne) Save(ctx context.Context) (*Middleware, error) {
	var (
		err  error
		node *Middleware
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiddlewareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (muo *MiddlewareUpdateOne) SaveX(ctx context.Context) *Middleware {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MiddlewareUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MiddlewareUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MiddlewareUpdateOne) sqlSave(ctx context.Context) (_node *Middleware, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   middleware.Table,
			Columns: middleware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: middleware.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Middleware.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, middleware.FieldID)
		for _, f := range fields {
			if !middleware.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != middleware.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.LID(); ok {
		_spec.SetField(middleware.FieldLID, field.TypeString, value)
	}
	if value, ok := muo.mutation.DID(); ok {
		_spec.SetField(middleware.FieldDID, field.TypeString, value)
	}
	if value, ok := muo.mutation.SysID(); ok {
		_spec.SetField(middleware.FieldSysID, field.TypeString, value)
	}
	_node = &Middleware{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{middleware.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}