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

func (r Repository) CreateUser(email, password string) error {
	var id int
	var createdAt string
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, created_at`
	err := r.db.Conn.QueryRow(query, email, password).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetUserById(id string) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE 1 = $1`
	row := r.db.Conn.QueryRow(query, id)
	switch err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, db.ErrNoMatch
	default:
		return user, err
	}
}

func (r Repository) GetUserByEmailAndPass(email, pass string) (User, error) {
	user := User{}
	query := `SELECT * FROM users WHERE email = $1, password = $2;`
	row := r.db.Conn.QueryRow(query, email, pass)
	switch err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, db.ErrNoMatch
	default:
		return user, err
	}
}
