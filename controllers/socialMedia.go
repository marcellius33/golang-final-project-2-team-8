package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"strconv"
)

type SocialMediaController struct {
	service services.SocialMediaService
}

func NewSocialMediaController(service services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		service: service,
	}
}

func (s *SocialMediaController) GetSocialMedias(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	socialMedias, err := s.service.GetSocialMedias(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(socialMedias, "Get Social Medias Success"))
}

func (s *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	socialMediaRequest := params.AddSocialMediaRequest{}
	if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	createSocialMedia, err := s.service.CreateSocialMedia(uint(userData["id"].(float64)), socialMediaRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessCreateResponse(createSocialMedia, "Create Social Media Success"))
}

func (s *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	socialMediaRequest := params.AddSocialMediaRequest{}
	if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

	updateSocialMedia, err := s.service.UpdateSocialMedia(uint(socialMediaId), socialMediaRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(updateSocialMedia, "Update Social Media Success"))
}

func (s *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	err = s.service.DeleteSocialMedia(uint(socialMediaId))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.DeleteSuccess("Your social media has been successfully deleted"))
}
