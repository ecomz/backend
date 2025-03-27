package repository

import (
	"github.com/ecomz/backend/auth-service/internal/model"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db    *sqlx.DB
	redis *redis.Pool
}

func NewUserRepository(db *sqlx.DB, redis *redis.Pool) UserRepository {
	return &userRepository{
		db:    db,
		redis: redis,
	}
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (name, email, password, role_id)
		VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.RoleID).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
