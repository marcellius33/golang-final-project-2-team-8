package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"strconv"
)

type PhotoController struct {
	service services.PhotoService
}

func NewPhotoController(service services.PhotoService) *PhotoController {
	return &PhotoController{
		service: service,
	}
}

func (p *PhotoController) GetPhotos(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	photos, err := p.service.GetPhotos(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(photos, "Get Photos Success"))
}

func (p *PhotoController) CreatePhoto(c *gin.Context) {
	photoRequest := params.CreatePhotoRequest{}
	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	createPhoto, err := p.service.CreatePhoto(uint(userData["id"].(float64)), photoRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessCreateResponse(createPhoto, "Create Photo Success"))
}

func (p *PhotoController) UpdatePhoto(c *gin.Context) {
	photoRequest := params.CreatePhotoRequest{}
	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	photoId, err := strconv.Atoi(c.Param("photoId"))

	updatePhoto, err := p.service.UpdatePhoto(uint(photoId), photoRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(updatePhoto, "Update Photo Success"))
}

func (p *PhotoController) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	err = p.service.DeletePhoto(uint(photoId))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.DeleteSuccess("Your photo has been successfully deleted"))
}
