package repositories

import (
	"database/sql"
	"github.com/HarvinRaj/goldshop/internal/models"
)

type Repository interface {
	SaveAll(*models.Users) (*models.Users, error)
	GetUserByEmail(*models.Users) (*models.Users, error)
	IsEmailExist(string) (bool, error)
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

func (u *UserRepository) SaveAll(user *models.Users) (*models.Users, error) {

	query := `INSERT INTO users (
		username,
		email,
		password_hash,
		first_name,
		last_name,
		role_id,
		is_active,
		created_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := u.db.Exec(query,
		user.UserName,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.RoleID,
		true,
		user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	userQuery := `SELECT user_id, username, email, first_name, last_name, role_id, is_active, created_at
				  FROM users
				  WHERE user_id = ?`

	var userDB models.Users

	err = u.db.QueryRow(userQuery, id).Scan(
		&userDB.UserID,
		&userDB.UserName,
		&userDB.Email,
		&userDB.FirstName,
		&userDB.LastName,
		&userDB.RoleID,
		&userDB.IsActive,
		&userDB.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &userDB, nil
}

func (u *UserRepository) GetUserByEmail(req *models.Users) (*models.Users, error) {

	var user models.Users

	query := "SELECT user_id, username, password_hash FROM users WHERE email = ?"

	err := u.db.QueryRow(query, req.Email).Scan(&user.UserID, &user.UserName, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
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

	query := "SELECT user_id, username, email, role_id, is_active FROM users"

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []*models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.RoleID, &user.IsActive); err != nil {
			return nil, err
		}

		userList = append(userList, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}
