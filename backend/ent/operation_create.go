// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/model"
	"backend/ent/operation"
	"backend/ent/roleoperation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OperationCreate is the builder for creating a Operation entity.
type OperationCreate struct {
	config
	mutation *OperationMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (oc *OperationCreate) SetName(s string) *OperationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetModelID sets the "model_id" field.
func (oc *OperationCreate) SetModelID(i int) *OperationCreate {
	oc.mutation.SetModelID(i)
	return oc
}

// SetModel sets the "model" edge to the Model entity.
func (oc *OperationCreate) SetModel(m *Model) *OperationCreate {
	return oc.SetModelID(m.ID)
}

// AddRoleOperationIDs adds the "role_operations" edge to the RoleOperation entity by IDs.
func (oc *OperationCreate) AddRoleOperationIDs(ids ...int) *OperationCreate {
	oc.mutation.AddRoleOperationIDs(ids...)
	return oc
}

// AddRoleOperations adds the "role_operations" edges to the RoleOperation entity.
func (oc *OperationCreate) AddRoleOperations(r ...*RoleOperation) *OperationCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return oc.AddRoleOperationIDs(ids...)
}

// Mutation returns the OperationMutation object of the builder.
func (oc *OperationCreate) Mutation() *OperationMutation {
	return oc.mutation
}

// Save creates the Operation in the database.
func (oc *OperationCreate) Save(ctx context.Context) (*Operation, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OperationCreate) SaveX(ctx context.Context) *Operation {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OperationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OperationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OperationCreate) check() error {
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Operation.name"`)}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := operation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Operation.name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ModelID(); !ok {
		return &ValidationError{Name: "model_id", err: errors.New(`ent: missing required field "Operation.model_id"`)}
	}
	if len(oc.mutation.ModelIDs()) == 0 {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required edge "Operation.model"`)}
	}
	return nil
}

func (oc *OperationCreate) sqlSave(ctx context.Context) (*Operation, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OperationCreate) createSpec() (*Operation, *sqlgraph.CreateSpec) {
	var (
		_node = &Operation{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(operation.Table, sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(operation.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := oc.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operation.ModelTable,
			Columns: []string{operation.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ModelID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.RoleOperationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operation.RoleOperationsTable,
			Columns: []string{operation.RoleOperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roleoperation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OperationCreateBulk is the builder for creating many Operation entities in bulk.
type OperationCreateBulk struct {
	config
	err      error
	builders []*OperationCreate
}

// Save creates the Operation entities in the database.
func (ocb *OperationCreateBulk) Save(ctx context.Context) ([]*Operation, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Operation, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OperationCreateBulk) SaveX(ctx context.Context) []*Operation {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OperationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OperationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
