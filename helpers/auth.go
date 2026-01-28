package helpers

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthHelper struct {
	DB *PGManager
}

func NewAuthHelper(db *PGManager) *AuthHelper {
	return &AuthHelper{DB: db}
}

func (h *AuthHelper) Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := h.DB.DB.Exec(query, email, string(hashed))
	return err
}

func (h *AuthHelper) Login(email, password string) (int64, error) {
	var id int64
	var hash string

	query := `SELECT id, password FROM users WHERE email=$1`
	err := h.DB.DB.QueryRow(query, email).Scan(&id, &hash)
	if err == sql.ErrNoRows {
		return 0, errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return 0, errors.New("invalid credentials")
	}

	return id, nil
}
