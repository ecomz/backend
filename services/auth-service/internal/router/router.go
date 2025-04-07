package router

import (
	"net/http"

	"github.com/ecomz/backend/auth-service/internal/handler"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handler.UserHandler, roleHandler *handler.RoleHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	auth := api.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/role", roleHandler.GetAllRoles).Methods(http.MethodGet)
	auth.HandleFunc("/role", roleHandler.CreateRole).Methods(http.MethodPost)

	auth.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	auth.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	auth.HandleFunc("/current-user", userHandler.CurrentUser).Methods(http.MethodGet)

	return r
}
