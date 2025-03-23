// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/model"
	"backend/ent/operation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ModelCreate is the builder for creating a Model entity.
type ModelCreate struct {
	config
	mutation *ModelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *ModelCreate) SetName(s string) *ModelCreate {
	mc.mutation.SetName(s)
	return mc
}

// AddOperationIDs adds the "operations" edge to the Operation entity by IDs.
func (mc *ModelCreate) AddOperationIDs(ids ...int) *ModelCreate {
	mc.mutation.AddOperationIDs(ids...)
	return mc
}

// AddOperations adds the "operations" edges to the Operation entity.
func (mc *ModelCreate) AddOperations(o ...*Operation) *ModelCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mc.AddOperationIDs(ids...)
}

// Mutation returns the ModelMutation object of the builder.
func (mc *ModelCreate) Mutation() *ModelMutation {
	return mc.mutation
}

// Save creates the Model in the database.
func (mc *ModelCreate) Save(ctx context.Context) (*Model, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModelCreate) SaveX(ctx context.Context) *Model {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModelCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModelCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModelCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Model.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := model.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Model.name": %w`, err)}
		}
	}
	return nil
}

func (mc *ModelCreate) sqlSave(ctx context.Context) (*Model, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *ModelCreate) createSpec() (*Model, *sqlgraph.CreateSpec) {
	var (
		_node = &Model{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(model.Table, sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(model.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := mc.mutation.OperationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   model.OperationsTable,
			Columns: []string{model.OperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ModelCreateBulk is the builder for creating many Model entities in bulk.
type ModelCreateBulk struct {
	config
	err      error
	builders []*ModelCreate
}

// Save creates the Model entities in the database.
func (mcb *ModelCreateBulk) Save(ctx context.Context) ([]*Model, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Model, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModelMutation)
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
func (mcb *ModelCreateBulk) SaveX(ctx context.Context) []*Model {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModelCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModelCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
