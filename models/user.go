package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int        `json:"id"`
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}

func NewUser() *User {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		uuid UUID NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE
	)`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return nil
	}

	return &User{
		UUID:      uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) Create(r User) error {
	query := `
	INSERT INTO users (uuid, name, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	return db.QueryRow(
		query,
		r.UUID,
		r.Name,
		r.Email,
		r.Password,
		r.CreatedAt,
		r.UpdatedAt,
	).Scan(&r.ID)
}

func (u *User) GetUserByID(id int) (*User, error) {
	r := &User{}
	query := `
	SELECT id, uuid, name, email, password, created_at, updated_at, deleted_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL`

	err := db.QueryRow(query, id).Scan(
		&r.ID,
		&r.UUID,
		&r.Name,
		&r.Email,
		&r.Password,
		&r.CreatedAt,
		&r.UpdatedAt,
		&r.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return r, nil
}
