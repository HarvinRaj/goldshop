package repositories

import (
	"database/sql"
	"github.com/HarvinRaj/goldshop/internal/models"
)

type Repository interface {
	Save(user *models.Users) error
	GetUserByEmail(user *models.Users) error
	IsEmailExist(email string) (bool, error)
	GetAllUsersList() ([]*models.Users, error)
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

func (u *UserRepository) GetAllUsersList() ([]*models.Users, error) {

	query := "SELECT id, name, email FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []*models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		userList = append(userList, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}
