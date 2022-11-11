package profile

import "mime/multipart"

type updateProfileForm struct {
	Type  string                `form:"type" json:"type"`
	Logo  *multipart.FileHeader `form:"logo" json:"logo" swaggerignore:"true"`
	About string                `form:"about" json:"about"`
}
