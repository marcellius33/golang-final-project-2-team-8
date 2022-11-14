package repositories

import (
	"gorm.io/gorm"
	"mygram/models"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	GetSocialMedias(userId uint, socialMedias *[]models.SocialMedia) (*[]models.SocialMedia, error)
	UpdateSocialMedia(socialMediaId uint, socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(id uint) error
	FindSocialMediaById(socialMediaId uint) (*models.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (s *socialMediaRepository) CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	return socialMedia, s.db.Create(socialMedia).Error
}

func (s *socialMediaRepository) GetSocialMedias(userId uint, socialMedias *[]models.SocialMedia) (*[]models.SocialMedia, error) {
	err := s.db.Model(&models.SocialMedia{}).Where("user_id=?", userId).Find(&socialMedias).Error
	return socialMedias, err
}

func (s *socialMediaRepository) UpdateSocialMedia(socialMediaId uint, updateSocialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	socialMedia := updateSocialMedia
	err := s.db.Model(&socialMedia).Where("id=?", socialMediaId).Updates(updateSocialMedia).Error
	return socialMedia, err
}

func (s *socialMediaRepository) DeleteSocialMedia(socialMediaId uint) error {
	err := s.db.Where("id=?", socialMediaId).Delete(&models.SocialMedia{}).Error
	return err
}

func (s *socialMediaRepository) FindSocialMediaById(socialMediaId uint) (*models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}
	err := s.db.Where("id=?", socialMediaId).First(&socialMedia).Error
	return &socialMedia, err
}
