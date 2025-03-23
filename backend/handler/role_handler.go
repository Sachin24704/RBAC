package handler

import (
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RoleHandler struct {
    roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
    return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) AddOperationsToRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roleIDStr := vars["role_id"]
    roleID, err := strconv.Atoi(roleIDStr)
    if err != nil {
        http.Error(w, "Invalid role_id", http.StatusBadRequest)
        return
    }
    var req struct {
        OperationIDs []int `json:"operation_ids"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if err := h.roleService.AddOperationsToRole(r.Context(),roleID, req.OperationIDs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *RoleHandler) RemoveOperationsFromRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roleIDStr := vars["role_id"]

    roleID, err := strconv.Atoi(roleIDStr)
    if err != nil {
        http.Error(w, "Invalid role_id", http.StatusBadRequest)
        return
    }

    var req struct {
        OperationIDs []int `json:"operation_ids"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.roleService.RemoveOperationsFromRole(r.Context(), roleID, req.OperationIDs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
