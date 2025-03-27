package dto

import "github.com/ecomz/backend/auth-service/internal/model"

type UserResponse struct {
	ID     string       `json:"id"`
	Email  string       `json:"email"`
	Name   string       `json:"name"`
	RoleID int          `json:"role_id"`
	Role   RoleResponse `json:"role"`
}

type LoginAndRegisiterResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

func NewUserResponse(user *model.User, role *model.Role) UserResponse {
	return UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RoleID: user.RoleID,
		Role: RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		},
	}
}

func NewLoginResponse(user *model.User, role *model.Role, accessToken, refreshToken string) LoginAndRegisiterResponse {
	return LoginAndRegisiterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         NewUserResponse(user, role),
	}
}
