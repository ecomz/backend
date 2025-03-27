package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ecomz/backend/auth-service/internal/dto"
	"github.com/ecomz/backend/auth-service/internal/model"
	"github.com/ecomz/backend/auth-service/internal/service"
	"github.com/ecomz/backend/libs/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validationErrors := utils.ValidateStruct(loginRequest)
	if validationErrors != nil {
		utils.ValidationErrorResponse(w, validationErrors)
		return
	}

	result, err := h.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Login successful", result)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validationErrors := utils.ValidateStruct(registerRequest)
	if validationErrors != nil {
		utils.ValidationErrorResponse(w, validationErrors)
		return
	}

	user := model.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		RoleID:   registerRequest.RoleID,
	}
	res, err := h.userService.Register(&user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "User registered successfully", res)
}

func (h *UserHandler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	// get token from header
	authHeader := r.Header.Get("Authorization")
	accessToken := strings.Split(authHeader, "Bearer ")
	if len(accessToken) != 2 {
		utils.ErrorResponse(w, http.StatusUnauthorized, "invalid access token")
		return
	}

	// get user from token
	user, err := h.userService.CurrentUser(accessToken[1])
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}
