package params

import (
	"time"

	"mygram/models"
)

type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type CreatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseToCreatePhotoResponse(photo *models.Photo) CreatePhotoResponse {
	return CreatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ParseToUpdatePhotoResponse(photo *models.Photo) UpdatePhotoResponse {
	return UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
}

type GetPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}

func ParseToGetPhotosResponse(photos []models.Photo, u models.User) []GetPhotoResponse {
	var responses []GetPhotoResponse
	for _, photo := range photos {
		responses = append(responses, ParseToGetPhotoResponse(photo, u))
	}
	return responses
}

func ParseToGetPhotoResponse(p models.Photo, u models.User) GetPhotoResponse {
	return GetPhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		User: struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			Email:    u.Email,
			Username: u.Username,
		},
	}
}
