package service

import (
	"github.com/ecomz/backend/auth-service/internal/model"
	"github.com/ecomz/backend/auth-service/internal/repository"
	"go.uber.org/zap"
)

type RoleService interface {
	GetRoleByID(roleID int) (*model.Role, error)
	GetAllRoles() ([]*model.Role, error)
	CreateRole(name string) error
	DeleteRole(roleID int) error
}

type roleService struct {
	logger   *zap.Logger
	roleRepo repository.RoleRepository
}

func NewRoleService(logger *zap.Logger, roleRepo repository.RoleRepository) RoleService {
	return &roleService{logger: logger, roleRepo: roleRepo}
}

func (s *roleService) GetRoleByID(roleID int) (*model.Role, error) {
	s.logger.Info("Getting role by ID",
		zap.Int("role_id", roleID),
	)
	return s.roleRepo.GetRoleByID(roleID)
}

func (s *roleService) GetAllRoles() ([]*model.Role, error) {
	roles, err := s.roleRepo.GetAllRoles()
	if err != nil {
		s.logger.Error("error getting all roles", zap.Error(err))
		return nil, err
	}
	return roles, nil
}

func (s *roleService) CreateRole(name string) error {
	s.logger.Info("Creating role",
		zap.String("role_name", name),
	)
	if err := s.roleRepo.CreateRole(name); err != nil {
		s.logger.Error("error creating role", zap.Error(err))
		return err
	}

	return nil
}

func (s *roleService) DeleteRole(roleID int) error {
	s.logger.Info("Deleting role",
		zap.Int("role_id", roleID),
	)
	return s.roleRepo.DeleteRole(roleID)
}
