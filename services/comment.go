package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type CommentService interface {
	CreateComment(userId uint, createCommentRequest params.CreateCommentRequest) (*params.CreateCommentResponse, error)
	GetComments(userId uint) (*[]params.GetCommentResponse, error)
	UpdateComment(commentId uint, commentUpdateRequest params.UpdateCommentRequest) (*params.UpdateCommentResponse, error)
	DeleteComment(commentId uint) error
}

type commentService struct {
	commentR repositories.CommentRepository
	userR    repositories.UserRepository
}

func NewCommentService(commentR repositories.CommentRepository, userR repositories.UserRepository) CommentService {
	return &commentService{
		commentR: commentR,
		userR:    userR,
	}
}

func (c *commentService) CreateComment(userId uint, createCommentRequest params.CreateCommentRequest) (*params.CreateCommentResponse, error) {
	newComment := models.Comment{
		UserID:  userId,
		PhotoID: createCommentRequest.PhotoId,
		Message: createCommentRequest.Message,
	}

	_, err := c.commentR.CreateComment(&newComment)

	if err != nil {
		return &params.CreateCommentResponse{}, err
	}
	resp := params.ParseToCreateCommentResponse(&newComment)

	return &resp, nil
}

func (c *commentService) GetComments(userId uint) (*[]params.GetCommentResponse, error) {
	var comments []models.Comment
	_, err := c.commentR.GetComments(userId, &comments)

	if err != nil {
		return &[]params.GetCommentResponse{}, err
	}
	user, _ := c.userR.FindUserByID(userId)
	resp := params.ParseToGetCommentsResponse(comments, *user)

	return &resp, nil
}

func (c *commentService) UpdateComment(commentId uint, commentUpdateRequest params.UpdateCommentRequest) (*params.UpdateCommentResponse, error) {
	commentModel, err := c.commentR.FindCommentById(commentId)
	if err != nil {
		return &params.UpdateCommentResponse{}, err
	}
	commentModel.Message = commentUpdateRequest.Message

	_, err = c.commentR.UpdateComment(commentId, commentModel)

	if err != nil {
		return &params.UpdateCommentResponse{}, err
	}
	resp := params.ParseToUpdateCommentResponse(commentModel)

	return &resp, nil
}

func (c *commentService) DeleteComment(commentId uint) error {
	return c.commentR.DeleteComment(commentId)
}
