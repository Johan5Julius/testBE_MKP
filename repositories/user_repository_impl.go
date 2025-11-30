package repositories

import (
	"database/sql"
	"testMKP/models"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) Create(user models.UserCreateRequest) (models.User, error) {
	query := "INSERT INTO users(name, email) VALUES ($1, $2) RETURNING id, name, email"

	var result models.User
	err := u.DB.QueryRow(query, user.Name, user.Email).Scan(&result.ID, &result.Name, &result.Email)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (u *UserRepositoryImpl) FindAll() ([]models.User, error) {
	query := "SELECT id, name, email FROM users"

	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindByID(id int) (models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"

	var user models.User
	err := u.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Update(user models.UserUpdateRequest) (models.User, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email"

	var result models.User
	err := u.DB.QueryRow(query, user.ID, user.Name, user.Email).Scan(&result.ID, &result.Name, &result.Email)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (u *UserRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := u.DB.Exec(query, id)
	return err
}
