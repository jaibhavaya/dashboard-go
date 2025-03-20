package models

import (
	"database/sql"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByID(id string) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	return user, err
}

func (r *UserRepository) Create(user User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	return err
}
