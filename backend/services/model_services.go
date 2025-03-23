package services

import (
	"backend/ent"
	"context"
)

type ModelService struct{
	client *ent.Client
}

func NewModelService(client *ent.Client) *ModelService {
	return &ModelService{client}
}

// update operations for a model for a role
func(s *ModelService) ListModelsWithOperations(ctx context.Context) ([]*ent.Model, error) {
	models, err := s.client.Model.Query().WithOperations().All(ctx)
	if err != nil {
		return nil, err
	}
	return models, nil
}
