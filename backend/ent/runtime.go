// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/model"
	"backend/ent/operation"
	"backend/ent/role"
	"backend/ent/schema"
	"backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	modelFields := schema.Model{}.Fields()
	_ = modelFields
	// modelDescName is the schema descriptor for name field.
	modelDescName := modelFields[0].Descriptor()
	// model.NameValidator is a validator for the "name" field. It is called by the builders before save.
	model.NameValidator = modelDescName.Validators[0].(func(string) error)
	operationFields := schema.Operation{}.Fields()
	_ = operationFields
	// operationDescName is the schema descriptor for name field.
	operationDescName := operationFields[0].Descriptor()
	// operation.NameValidator is a validator for the "name" field. It is called by the builders before save.
	operation.NameValidator = operationDescName.Validators[0].(func(string) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
}
