package repositories

import (
	"database/sql"
	"github.com/HarvinRaj/goldshop/internal/models"
)

type Repository interface {
	Save(user *models.Users) error
	GetUserByEmail(user *models.Users) error
	IsEmailExist(email string) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Save(user *models.Users) error {

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	_, err := u.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(*models.Users) error {
	return nil
}

func (u *UserRepository) IsEmailExist(email string) (bool, error) {

	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"

	err := u.db.QueryRow(query, email).Scan(&exist)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return exist, nil
}
