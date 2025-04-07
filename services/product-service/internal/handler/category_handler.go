package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ecomz/backend/libs/logger"
	"github.com/ecomz/backend/libs/utils"
	"github.com/ecomz/backend/product-service/internal/dto"
	"github.com/ecomz/backend/product-service/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	logger  logger.Logger
	service service.CategoryService
}

func NewCategoryHandler(logger logger.Logger, categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		logger:  logger,
		service: categoryService,
	}
}

func (ch *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// define req
	var req dto.CreateCategoryRequest

	// decode json
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// validate
	validationErr := utils.ValidateStruct(req)
	if validationErr != nil {
		utils.ValidationErrorResponse(w, validationErr)
		return
	}

	// call service
	category, err := ch.service.CreateCategory(&req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Category created successfully", category)
}

func (ch *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ch.service.GetAllCategories()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Categories fetched successfully", categories)
}

func (ch *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ch.logger.Error(err.Error(), zap.String("id", id))
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	category, err := ch.service.GetCategoryByID(idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Category fetched successfully", category)
}

func (ch *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ch.logger.Error(err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// define req
	var req dto.UpdateCategoryRequest

	// decode json
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// validate
	validationErr := utils.ValidateStruct(req)
	if validationErr != nil {
		utils.ValidationErrorResponse(w, validationErr)
		return
	}

	// call service
	err = ch.service.UpdateCategory(idInt, &req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Category updated successfully", nil)
}

func (ch *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ch.logger.Error(err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// call service
	err = ch.service.DeleteCategory(idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Category deleted successfully", nil)
}
