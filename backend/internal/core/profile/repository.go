package profile

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

func (r Repository) GetProfile(userID string) (Profile, error) {
	profile := Profile{}
	query := `SELECT id, user_id, type, logo, about FROM user_profile WHERE user_id = $1`
	row := r.db.Conn.QueryRow(query, userID)
	switch err := row.Scan(&profile.ID, &profile.UserID, &profile.Type, &profile.Logo, &profile.About); err {
	case sql.ErrNoRows:
		return profile, db.ErrNoMatch
	default:
		return profile, err
	}
}

func (r Repository) UpdateProfile(userID, userType, logo, about string) error {
	var id int
	query := `
	INSERT INTO user_profile ("user_id", "type", "logo", "about")
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (user_id)
	DO UPDATE
	SET user_id = EXCLUDED.user_id, type = EXCLUDED.type, logo = EXCLUDED.logo, about = EXCLUDED.about
	RETURNING id`
	err := r.db.Conn.QueryRow(query, userID, userType, logo, about).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
