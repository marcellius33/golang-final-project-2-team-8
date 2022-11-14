package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(userId uint, createSocialMediaRequest params.AddSocialMediaRequest) (*params.AddSocialMediaResponse, error)
	GetSocialMedias(userId uint) (*[]params.GetSocialMediaResponse, error)
	UpdateSocialMedia(photoId uint, photoUpdateRequest params.AddSocialMediaRequest) (*params.UpdateSocialMediaResponse, error)
	DeleteSocialMedia(photoId uint) error
}

type socialMediaService struct {
	socialMediaR repositories.SocialMediaRepository
	userR        repositories.UserRepository
}

func NewSocialMediaService(socialMediaR repositories.SocialMediaRepository, userR repositories.UserRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaR: socialMediaR,
		userR:        userR,
	}
}

func (s *socialMediaService) CreateSocialMedia(userId uint, createSocialMediaRequest params.AddSocialMediaRequest) (*params.AddSocialMediaResponse, error) {
	newSocialMedia := models.SocialMedia{
		Name:           createSocialMediaRequest.Name,
		SocialMediaURL: createSocialMediaRequest.SocialMediaURL,
		UserID:         userId,
	}

	_, err := s.socialMediaR.CreateSocialMedia(&newSocialMedia)

	if err != nil {
		return &params.AddSocialMediaResponse{}, err
	}
	resp := params.ParseToAddSocialMediaResponse(&newSocialMedia)

	return &resp, nil
}

func (s *socialMediaService) GetSocialMedias(userId uint) (*[]params.GetSocialMediaResponse, error) {
	var socialMedias []models.SocialMedia
	_, err := s.socialMediaR.GetSocialMedias(userId, &socialMedias)

	if err != nil {
		return &[]params.GetSocialMediaResponse{}, err
	}
	user, _ := s.userR.FindUserByID(userId)
	resp := params.ParseToGetSocialMediasResponse(socialMedias, *user)

	return &resp, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaId uint, socialMediaUpdateRequest params.AddSocialMediaRequest) (*params.UpdateSocialMediaResponse, error) {
	socialMediaModel, err := s.socialMediaR.FindSocialMediaById(socialMediaId)
	if err != nil {
		return &params.UpdateSocialMediaResponse{}, err
	}
	socialMediaModel.SocialMediaURL = socialMediaUpdateRequest.SocialMediaURL
	socialMediaModel.Name = socialMediaUpdateRequest.Name

	_, err = s.socialMediaR.UpdateSocialMedia(socialMediaId, socialMediaModel)

	if err != nil {
		return &params.UpdateSocialMediaResponse{}, err
	}
	resp := params.ParseToUpdateSocialMediaResponse(socialMediaModel)

	return &resp, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) error {
	return s.socialMediaR.DeleteSocialMedia(socialMediaId)
}
