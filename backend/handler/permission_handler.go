package handler

import (
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PermissionHandler struct {
	permService *services.PermissionService
}

func NewPermissionHandler(permService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permService: permService,
	}
}

func (h *PermissionHandler) GetUserPermissions(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	permissions, err := h.permService.GetUserPermissions(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(permissions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
