package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"strconv"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{
		service: service,
	}
}

func (cC *CommentController) GetComments(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	comments, err := cC.service.GetComments(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(comments, "Get Comments Success"))
}

func (cC *CommentController) CreateComment(c *gin.Context) {
	commentRequest := params.CreateCommentRequest{}
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	createComment, err := cC.service.CreateComment(uint(userData["id"].(float64)), commentRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessCreateResponse(createComment, "Create Comment Success"))
}

func (cC *CommentController) UpdateComment(c *gin.Context) {
	commentRequest := params.UpdateCommentRequest{}
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	commentId, err := strconv.Atoi(c.Param("commentId"))

	updateComment, err := cC.service.UpdateComment(uint(commentId), commentRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(updateComment, "Update Comment Success"))
}

func (cC *CommentController) DeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	err = cC.service.DeleteComment(uint(commentId))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.DeleteSuccess("Your comment has been successfully deleted"))
}
