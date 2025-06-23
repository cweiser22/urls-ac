package repository

import (
	"fmt"
	"github.com/cweiser22/urls-ac/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository creates a new UserRepository with the provided database connection.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type CreateUserDTO struct {
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

type UpdateUserDTO struct {
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

// InsertUser creates a new user and returns a populated User.
func (r *UserRepository) InsertUser(dto *CreateUserDTO) (*models.User, error) {
	query := `
		INSERT INTO accounts (email, password_hash)
		VALUES (:email, :password_hash)
		RETURNING id
	`
	rows, err := r.DB.NamedQuery(query, dto)
	if err != nil {
		return nil, fmt.Errorf("insert user: %w", err)
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan user id: %w", err)
		}
	}

	return &models.User{
		ID:           id,
		Email:        dto.Email,
		PasswordHash: dto.PasswordHash,
	}, nil
}

// UpdateUser updates an existing user by ID.
func (r *UserRepository) UpdateUser(id int, dto *UpdateUserDTO) error {
	query := `
		UPDATE accounts
		SET email = :email, password_hash = :password_hash
		WHERE id = :id
	`

	params := map[string]interface{}{
		"id":            id,
		"email":         dto.Email,
		"password_hash": dto.PasswordHash,
	}

	_, err := r.DB.NamedExec(query, params)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user by ID.
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}

// GetUserByID retrieves a user by ID.
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash FROM accounts WHERE id = $1`
	if err := r.DB.Get(&user, query, id); err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email.
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash FROM accounts WHERE email = $1`
	if err := r.DB.Get(&user, query, email); err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &user, nil
}
