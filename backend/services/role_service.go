package services

import (
	"backend/ent"
	"backend/ent/roleoperation"
	"context"
	"fmt"
)

type RoleService struct {
    client *ent.Client
}

func NewRoleService(client *ent.Client) *RoleService {
    return &RoleService{client: client}
}

func (s *RoleService) AddOperationsToRole(ctx context.Context,roleID int, operationIDs []int) error {
    existingOps, err := s.client.RoleOperation.Query().
        Where(roleoperation.RoleID(roleID)).
        Select(roleoperation.FieldOperationID).
        Ints(ctx)
    if err != nil {
        return fmt.Errorf("failed fetching existing operations: %v", err)
    }
    existingOpsMap := make(map[int]struct{}, len(existingOps))
    for _, opID := range existingOps {
        existingOpsMap[opID] = struct{}{}
    }
    var bulk []*ent.RoleOperationCreate
    for _, opID := range operationIDs {
        if _, exists := existingOpsMap[opID]; !exists {
            bulk = append(bulk, s.client.RoleOperation.Create().
                SetRoleID(roleID).
                SetOperationID(opID))
        }
    }
    if len(bulk) > 0 {
        if _, err := s.client.RoleOperation.CreateBulk(bulk...).Save(ctx); err != nil {
            return fmt.Errorf("failed adding new operations: %v", err)
        }
    }

    return nil
}

func (s *RoleService) RemoveOperationsFromRole(ctx context.Context, roleID int, operationIDs []int) error {
    _, err := s.client.RoleOperation.Delete().
        Where(
            roleoperation.RoleID(roleID),
            roleoperation.OperationIDIn(operationIDs...),
        ).Exec(ctx)
    if err != nil {
        return fmt.Errorf("failed removing operations: %v", err)
    }

    return nil
}
