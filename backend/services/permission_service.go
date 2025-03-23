package services

import (
	"context"
	"backend/ent"
	"backend/ent/user"
)

type PermissionService struct {
	client *ent.Client
}

func NewPermissionService(client *ent.Client) *PermissionService {
	return &PermissionService{
		client: client,
	}
}

type PermissionsResponse struct {
	UserID int              `json:"userId"`
	Name   string           `json:"name"`
	Roles  []RoleWithAccess `json:"roles"`
}

type RoleWithAccess struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Models []ModelAccess `json:"models"`
}

type ModelAccess struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *PermissionService) GetUserPermissions(ctx context.Context, userID int) (*PermissionsResponse, error) {
	u, err := s.client.User.Query().
		Where(user.ID(userID)).
		WithUserRoles(func(q *ent.UserRoleQuery) {
			q.WithRole(func(q *ent.RoleQuery) {
				q.WithRoleOperations(func(q *ent.RoleOperationQuery) {
					q.WithOperations(func(q *ent.OperationQuery) {
						q.WithModel()
					})
				})
			})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	resp := &PermissionsResponse{
		UserID: u.ID,
		Name:   u.Name,
		Roles:  make([]RoleWithAccess, 0),
	}
	roleMap := make(map[int]*RoleWithAccess)
	for _, ur := range u.Edges.UserRoles {
		role := ur.Edges.Role
		if role == nil {
			continue
		}
		if _, exists := roleMap[role.ID]; !exists {
			roleMap[role.ID] = &RoleWithAccess{
				ID:     role.ID,
				Name:   role.Name,
				Models: make([]ModelAccess, 0),
			}
		}
		modelMap := make(map[int]*ModelAccess)
		for _, ro := range role.Edges.RoleOperations {
			op := ro.Edges.Operations
			if op == nil || op.Edges.Model == nil {
				continue
			}

			model := op.Edges.Model
			if _, exists := modelMap[model.ID]; !exists {
				modelMap[model.ID] = &ModelAccess{
					ID:         model.ID,
					Name:       model.Name,
					Operations: make([]Operation, 0),
				}
			}
			modelMap[model.ID].Operations = append(modelMap[model.ID].Operations, Operation{
				ID:   op.ID,
				Name: op.Name,
			})
		}
		for _, model := range modelMap {
			roleMap[role.ID].Models = append(roleMap[role.ID].Models, *model)
		}
	}
	for _, role := range roleMap {
		resp.Roles = append(resp.Roles, *role)
	}
	return resp, nil
}
