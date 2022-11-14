package repositories

import (
	"gorm.io/gorm"
	"mygram/models"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetComments(userId uint, Comments *[]models.Comment) (*[]models.Comment, error)
	UpdateComment(CommentId uint, Comment *models.Comment) (*models.Comment, error)
	DeleteComment(id uint) error
	FindCommentById(CommentId uint) (*models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) CreateComment(Comment *models.Comment) (*models.Comment, error) {
	return Comment, c.db.Create(Comment).Error
}

func (c *commentRepository) GetComments(userId uint, Comments *[]models.Comment) (*[]models.Comment, error) {
	err := c.db.Preload("Photo").Model(&models.Comment{}).Where("user_id=?", userId).Find(&Comments).Error
	return Comments, err
}

func (c *commentRepository) UpdateComment(CommentId uint, updateComment *models.Comment) (*models.Comment, error) {
	Comment := updateComment
	err := c.db.Model(&Comment).Where("id=?", CommentId).Updates(updateComment).Error
	return Comment, err
}

func (c *commentRepository) DeleteComment(CommentId uint) error {
	err := c.db.Where("id=?", CommentId).Delete(&models.Comment{}).Error
	return err
}

func (c *commentRepository) FindCommentById(CommentId uint) (*models.Comment, error) {
	Comment := models.Comment{}
	err := c.db.Where("id=?", CommentId).First(&Comment).Error
	return &Comment, err
}
