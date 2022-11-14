package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type PhotoService interface {
	CreatePhoto(userId uint, createPhotoRequest params.CreatePhotoRequest) (*params.CreatePhotoResponse, error)
	GetPhotos(userId uint) (*[]params.GetPhotoResponse, error)
	UpdatePhoto(photoId uint, photoUpdateRequest params.CreatePhotoRequest) (*params.UpdatePhotoResponse, error)
	DeletePhoto(photoId uint) error
}

type photoService struct {
	photoR repositories.PhotoRepository
	userR  repositories.UserRepository
}

func NewPhotoService(photoR repositories.PhotoRepository, userR repositories.UserRepository) PhotoService {
	return &photoService{
		photoR: photoR,
		userR:  userR,
	}
}

func (p *photoService) CreatePhoto(userId uint, createPhotoRequest params.CreatePhotoRequest) (*params.CreatePhotoResponse, error) {
	newPhoto := models.Photo{
		Title:    createPhotoRequest.Title,
		Caption:  createPhotoRequest.Caption,
		PhotoURL: createPhotoRequest.PhotoURL,
		UserID:   userId,
	}

	_, err := p.photoR.CreatePhoto(&newPhoto)

	if err != nil {
		return &params.CreatePhotoResponse{}, err
	}
	resp := params.ParseToCreatePhotoResponse(&newPhoto)

	return &resp, nil
}

func (p *photoService) GetPhotos(userId uint) (*[]params.GetPhotoResponse, error) {
	var photos []models.Photo
	_, err := p.photoR.GetPhotos(userId, &photos)

	if err != nil {
		return &[]params.GetPhotoResponse{}, err
	}
	user, _ := p.userR.FindUserByID(userId)
	resp := params.ParseToGetPhotosResponse(photos, *user)

	return &resp, nil
}

func (p *photoService) UpdatePhoto(photoId uint, photoUpdateRequest params.CreatePhotoRequest) (*params.UpdatePhotoResponse, error) {
	photoModel, err := p.photoR.FindPhotoById(photoId)
	if err != nil {
		return &params.UpdatePhotoResponse{}, err
	}
	photoModel.Title = photoUpdateRequest.Title
	photoModel.Caption = photoUpdateRequest.Caption
	photoModel.PhotoURL = photoUpdateRequest.PhotoURL

	_, err = p.photoR.UpdatePhoto(photoId, photoModel)

	if err != nil {
		return &params.UpdatePhotoResponse{}, err
	}
	resp := params.ParseToUpdatePhotoResponse(photoModel)

	return &resp, nil
}

func (p *photoService) DeletePhoto(photoId uint) error {
	return p.photoR.DeletePhoto(photoId)
}
