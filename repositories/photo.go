package repositories

import (
	"gorm.io/gorm"
	"mygram/models"
)

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) (*models.Photo, error)
	GetPhotos(userId uint, photos *[]models.Photo) (*[]models.Photo, error)
	UpdatePhoto(photoId uint, photo *models.Photo) (*models.Photo, error)
	DeletePhoto(id uint) error
	FindPhotoById(photoId uint) (*models.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (p *photoRepository) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	return photo, p.db.Create(photo).Error
}

func (p *photoRepository) GetPhotos(userId uint, photos *[]models.Photo) (*[]models.Photo, error) {
	err := p.db.Model(&models.Photo{}).Where("user_id=?", userId).Find(&photos).Error
	return photos, err
}

func (p *photoRepository) UpdatePhoto(photoId uint, updatePhoto *models.Photo) (*models.Photo, error) {
	photo := updatePhoto
	err := p.db.Model(&photo).Where("id=?", photoId).Updates(updatePhoto).Error
	return photo, err
}

func (p *photoRepository) DeletePhoto(photoId uint) error {
	err := p.db.Where("id=?", photoId).Delete(&models.Photo{}).Error
	return err
}

func (p *photoRepository) FindPhotoById(photoId uint) (*models.Photo, error) {
	photo := models.Photo{}
	err := p.db.Where("id=?", photoId).First(&photo).Error
	return &photo, err
}
