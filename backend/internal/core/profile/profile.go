package profile

type Profile struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	Type   string `json:"type" db:"type"`
	Logo   string `json:"logo" db:"logo"`
	About  string `json:"about" db:"about"`
}
