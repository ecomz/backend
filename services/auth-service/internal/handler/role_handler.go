package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ecomz/backend/auth-service/internal/dto"
	"github.com/ecomz/backend/auth-service/internal/service"
	"github.com/ecomz/backend/libs/utils"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Roles retrieved successfully", roles)
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var roleRequest dto.RoleRequest

	if err := json.NewDecoder(r.Body).Decode(&roleRequest); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validationErrors := utils.ValidateStruct(roleRequest)
	if validationErrors != nil {
		utils.ValidationErrorResponse(w, validationErrors)
		return
	}

	if err := h.roleService.CreateRole(roleRequest.Name); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Role created successfully", nil)
}
