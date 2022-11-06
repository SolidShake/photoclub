package profile

type updateProfileForm struct {
	Type  string `form:"type" json:"type"`
	Logo  string `form:"logo" json:"logo"`
	About string `form:"about" json:"about"`
}
