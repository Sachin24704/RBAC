package handler

import (
	"backend/services"
	"encoding/json"
	"net/http"
)

type ModelHandler struct {
	modelService *services.ModelService
}

func NewModelHandler(srv *services.ModelService) *ModelHandler {
	return &ModelHandler{srv}
}

func (h *ModelHandler) ListModelsWithOperationsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	models, err := h.modelService.ListModelsWithOperations(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models)
}
