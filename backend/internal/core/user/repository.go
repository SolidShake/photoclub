package user

import (
	"database/sql"

	"github.com/SolidShake/photoclub/db"
)

type Repository struct {
	db db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{db: db}
}

func (r Repository) CreateUser(email, nickname, password string) error {
	var id int
	var createdAt string
	query := `INSERT INTO users (email, nickname, password) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.Conn.QueryRow(query, email, nickname, password).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetUserById(id string) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE 1 = $1`
	row := r.db.Conn.QueryRow(query, id)
	switch err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, db.ErrNoMatch
	default:
		return user, err
	}
}

func (r Repository) GetUserByEmail(email string) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE email = $1;`
	row := r.db.Conn.QueryRow(query, email)
	switch err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, db.ErrNoMatch
	default:
		return user, err
	}
}

func (r Repository) GetUserByNickname(nickname string) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE nickname = $1;`
	row := r.db.Conn.QueryRow(query, nickname)
	switch err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, db.ErrNoMatch
	default:
		return user, err
	}
}
