package params

import (
	"time"

	"mygram/models"
)

type AddSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

type AddSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func ParseToAddSocialMediaResponse(socmed *models.SocialMedia) AddSocialMediaResponse {
	return AddSocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		CreatedAt:      socmed.CreatedAt,
	}
}

type UpdateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ParseToUpdateSocialMediaResponse(socmed *models.SocialMedia) UpdateSocialMediaResponse {
	return UpdateSocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		UpdatedAt:      socmed.UpdatedAt,
	}
}

type GetSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           struct {
		ID              int    `json:"id"`
		Username        string `json:"username"`
		ProfileImageUrl string `json:"profile_image_url"`
	}
}

func ParseToGetSocialMediasResponse(socialMedias []models.SocialMedia, user models.User) []GetSocialMediaResponse {
	var responses []GetSocialMediaResponse
	for _, socialMedia := range socialMedias {
		responses = append(responses, ParseToGetSocialMediaResponse(socialMedia, user))
	}
	return responses
}

func ParseToGetSocialMediaResponse(socmed models.SocialMedia, user models.User) GetSocialMediaResponse {
	return GetSocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		CreatedAt:      socmed.CreatedAt,
		UpdatedAt:      socmed.UpdatedAt,
		User: struct {
			ID              int    `json:"id"`
			Username        string `json:"username"`
			ProfileImageUrl string `json:"profile_image_url"`
		}{
			ID:              int(user.ID),
			Username:        user.Username,
			ProfileImageUrl: socmed.SocialMediaURL,
		},
	}
}
