package profile

import (
	"errors"
	"log"
	"net/http"

	"github.com/SolidShake/photoclub/db"
	"github.com/SolidShake/photoclub/internal/api/auth"
	coreProfile "github.com/SolidShake/photoclub/internal/core/profile"
	"github.com/SolidShake/photoclub/pkg/validation"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *coreProfile.Service
}

func NewHandler(service *coreProfile.Service) *Handler {
	return &Handler{service: service}
}

// Profile godoc
// @Summary      Profile
// @Description  get user profile
// @Tags         Profile
// @Produce      json
// @Success      200  {object}  profileResponse
// @Failure      500
// @Security     ApiKeyAuth
// @Router       /user/profile [get]
func (h Handler) UserProfileGetHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	profile, err := h.service.GetProfile(claims[auth.IdentityKey].(string))
	if err == db.ErrNoMatch {
		c.Status(http.StatusOK)
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, errors.New("internal error"))
		return
	}

	response := profileResponse{
		Type:  profile.Type,
		Logo:  profile.Logo,
		About: profile.About,
	}

	c.JSON(http.StatusOK, response)
}

// Profile godoc
// @Summary      Profile
// @Description  get user profile
// @Tags         Profile
// @Produce      json
// @Param        request body updateProfileForm true "update profile form"
// @Success      200
// @Failure      400
// @Security     ApiKeyAuth
// @Router       /user/profile [put]
func (h Handler) UserProfileUpdateHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	var updateProfileVals updateProfileForm
	errs := validation.Validate(c, &updateProfileVals)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	err := h.service.UpdateProfile(
		claims[auth.IdentityKey].(string),
		updateProfileVals.Type,
		updateProfileVals.Logo,
		updateProfileVals.About,
	)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.New("internal error"))
		return
	}

	c.Status(http.StatusOK)
}
