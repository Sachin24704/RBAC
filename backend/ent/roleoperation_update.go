// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/operation"
	"backend/ent/predicate"
	"backend/ent/role"
	"backend/ent/roleoperation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleOperationUpdate is the builder for updating RoleOperation entities.
type RoleOperationUpdate struct {
	config
	hooks    []Hook
	mutation *RoleOperationMutation
}

// Where appends a list predicates to the RoleOperationUpdate builder.
func (rou *RoleOperationUpdate) Where(ps ...predicate.RoleOperation) *RoleOperationUpdate {
	rou.mutation.Where(ps...)
	return rou
}

// SetRoleID sets the "role_id" field.
func (rou *RoleOperationUpdate) SetRoleID(i int) *RoleOperationUpdate {
	rou.mutation.SetRoleID(i)
	return rou
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rou *RoleOperationUpdate) SetNillableRoleID(i *int) *RoleOperationUpdate {
	if i != nil {
		rou.SetRoleID(*i)
	}
	return rou
}

// SetOperationID sets the "operation_id" field.
func (rou *RoleOperationUpdate) SetOperationID(i int) *RoleOperationUpdate {
	rou.mutation.SetOperationID(i)
	return rou
}

// SetNillableOperationID sets the "operation_id" field if the given value is not nil.
func (rou *RoleOperationUpdate) SetNillableOperationID(i *int) *RoleOperationUpdate {
	if i != nil {
		rou.SetOperationID(*i)
	}
	return rou
}

// SetRole sets the "role" edge to the Role entity.
func (rou *RoleOperationUpdate) SetRole(r *Role) *RoleOperationUpdate {
	return rou.SetRoleID(r.ID)
}

// SetOperationsID sets the "operations" edge to the Operation entity by ID.
func (rou *RoleOperationUpdate) SetOperationsID(id int) *RoleOperationUpdate {
	rou.mutation.SetOperationsID(id)
	return rou
}

// SetOperations sets the "operations" edge to the Operation entity.
func (rou *RoleOperationUpdate) SetOperations(o *Operation) *RoleOperationUpdate {
	return rou.SetOperationsID(o.ID)
}

// Mutation returns the RoleOperationMutation object of the builder.
func (rou *RoleOperationUpdate) Mutation() *RoleOperationMutation {
	return rou.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rou *RoleOperationUpdate) ClearRole() *RoleOperationUpdate {
	rou.mutation.ClearRole()
	return rou
}

// ClearOperations clears the "operations" edge to the Operation entity.
func (rou *RoleOperationUpdate) ClearOperations() *RoleOperationUpdate {
	rou.mutation.ClearOperations()
	return rou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rou *RoleOperationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rou.sqlSave, rou.mutation, rou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rou *RoleOperationUpdate) SaveX(ctx context.Context) int {
	affected, err := rou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rou *RoleOperationUpdate) Exec(ctx context.Context) error {
	_, err := rou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rou *RoleOperationUpdate) ExecX(ctx context.Context) {
	if err := rou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rou *RoleOperationUpdate) check() error {
	if rou.mutation.RoleCleared() && len(rou.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleOperation.role"`)
	}
	if rou.mutation.OperationsCleared() && len(rou.mutation.OperationsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleOperation.operations"`)
	}
	return nil
}

func (rou *RoleOperationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(roleoperation.Table, roleoperation.Columns, sqlgraph.NewFieldSpec(roleoperation.FieldID, field.TypeInt))
	if ps := rou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rou.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.RoleTable,
			Columns: []string{roleoperation.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rou.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.RoleTable,
			Columns: []string{roleoperation.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rou.mutation.OperationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.OperationsTable,
			Columns: []string{roleoperation.OperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rou.mutation.OperationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.OperationsTable,
			Columns: []string{roleoperation.OperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rou.mutation.done = true
	return n, nil
}

// RoleOperationUpdateOne is the builder for updating a single RoleOperation entity.
type RoleOperationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleOperationMutation
}

// SetRoleID sets the "role_id" field.
func (rouo *RoleOperationUpdateOne) SetRoleID(i int) *RoleOperationUpdateOne {
	rouo.mutation.SetRoleID(i)
	return rouo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rouo *RoleOperationUpdateOne) SetNillableRoleID(i *int) *RoleOperationUpdateOne {
	if i != nil {
		rouo.SetRoleID(*i)
	}
	return rouo
}

// SetOperationID sets the "operation_id" field.
func (rouo *RoleOperationUpdateOne) SetOperationID(i int) *RoleOperationUpdateOne {
	rouo.mutation.SetOperationID(i)
	return rouo
}

// SetNillableOperationID sets the "operation_id" field if the given value is not nil.
func (rouo *RoleOperationUpdateOne) SetNillableOperationID(i *int) *RoleOperationUpdateOne {
	if i != nil {
		rouo.SetOperationID(*i)
	}
	return rouo
}

// SetRole sets the "role" edge to the Role entity.
func (rouo *RoleOperationUpdateOne) SetRole(r *Role) *RoleOperationUpdateOne {
	return rouo.SetRoleID(r.ID)
}

// SetOperationsID sets the "operations" edge to the Operation entity by ID.
func (rouo *RoleOperationUpdateOne) SetOperationsID(id int) *RoleOperationUpdateOne {
	rouo.mutation.SetOperationsID(id)
	return rouo
}

// SetOperations sets the "operations" edge to the Operation entity.
func (rouo *RoleOperationUpdateOne) SetOperations(o *Operation) *RoleOperationUpdateOne {
	return rouo.SetOperationsID(o.ID)
}

// Mutation returns the RoleOperationMutation object of the builder.
func (rouo *RoleOperationUpdateOne) Mutation() *RoleOperationMutation {
	return rouo.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rouo *RoleOperationUpdateOne) ClearRole() *RoleOperationUpdateOne {
	rouo.mutation.ClearRole()
	return rouo
}

// ClearOperations clears the "operations" edge to the Operation entity.
func (rouo *RoleOperationUpdateOne) ClearOperations() *RoleOperationUpdateOne {
	rouo.mutation.ClearOperations()
	return rouo
}

// Where appends a list predicates to the RoleOperationUpdate builder.
func (rouo *RoleOperationUpdateOne) Where(ps ...predicate.RoleOperation) *RoleOperationUpdateOne {
	rouo.mutation.Where(ps...)
	return rouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rouo *RoleOperationUpdateOne) Select(field string, fields ...string) *RoleOperationUpdateOne {
	rouo.fields = append([]string{field}, fields...)
	return rouo
}

// Save executes the query and returns the updated RoleOperation entity.
func (rouo *RoleOperationUpdateOne) Save(ctx context.Context) (*RoleOperation, error) {
	return withHooks(ctx, rouo.sqlSave, rouo.mutation, rouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rouo *RoleOperationUpdateOne) SaveX(ctx context.Context) *RoleOperation {
	node, err := rouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rouo *RoleOperationUpdateOne) Exec(ctx context.Context) error {
	_, err := rouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rouo *RoleOperationUpdateOne) ExecX(ctx context.Context) {
	if err := rouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rouo *RoleOperationUpdateOne) check() error {
	if rouo.mutation.RoleCleared() && len(rouo.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleOperation.role"`)
	}
	if rouo.mutation.OperationsCleared() && len(rouo.mutation.OperationsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleOperation.operations"`)
	}
	return nil
}

func (rouo *RoleOperationUpdateOne) sqlSave(ctx context.Context) (_node *RoleOperation, err error) {
	if err := rouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(roleoperation.Table, roleoperation.Columns, sqlgraph.NewFieldSpec(roleoperation.FieldID, field.TypeInt))
	id, ok := rouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoleOperation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roleoperation.FieldID)
		for _, f := range fields {
			if !roleoperation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roleoperation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rouo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.RoleTable,
			Columns: []string{roleoperation.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rouo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.RoleTable,
			Columns: []string{roleoperation.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rouo.mutation.OperationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.OperationsTable,
			Columns: []string{roleoperation.OperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rouo.mutation.OperationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleoperation.OperationsTable,
			Columns: []string{roleoperation.OperationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RoleOperation{config: rouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rouo.mutation.done = true
	return _node, nil
}
