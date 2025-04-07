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
)

type ProductHandler struct {
	logger  logger.Logger
	service service.ProductService
}

func NewProductHandler(logger logger.Logger, productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		logger:  logger,
		service: productService,
	}
}

func (ch *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// define req
	var req dto.CreateProductRequest

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
	product, err := ch.service.CreateProduct(&req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Product created successfully", product)
}

func (ch *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ch.service.GetAllProducts()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Products fetched successfully", products)
}

func (ch *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	product, err := ch.service.GetProductByID(idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Product fetched successfully", product)
}

func (ch *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// define req
	var req dto.UpdateProductRequest

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
	err = ch.service.UpdateProduct(idInt, &req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Product updated successfully", nil)
}

func (ch *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id from params
	id := mux.Vars(r)["id"]

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// call service
	err = ch.service.DeleteProduct(idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	utils.SuccessResponse(w, http.StatusOK, "Product deleted successfully", nil)
}
