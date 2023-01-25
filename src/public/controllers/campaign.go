package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/golibs-starter/golib/web/response"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/constants"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources/mappers"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/services"
	"net/http"
	"strings"
)

type CampaignController struct {
	CampaignService services.CampaignServiceable
}

func NewCampaignController(
	CampaignService services.CampaignServiceable,
) *CampaignController {
	return &CampaignController{
		CampaignService: CampaignService,
	}
}

// GetAllCampaign ...
//
//	@Summary		Get List campaign
//	@Description	Return list of campaign
//	@Tags			campaign
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]resources.Campaign}
//	@Failure		500	{object}	response.Response
//	@Router			/v1/campaigns [get]
func (r *CampaignController) GetAllCampaign(ctx *gin.Context) {
	var request resources.GetCampaignRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			return
		}
		log.Warn(ctx, "[CampaignController][GetAllCampaign] cannot bind request error: [%s]", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		return
	}
	result, err := r.CampaignService.GetAllCampaign(ctx, mappers.ToCampaignFilter(&request))
	if err != nil {
		log.Error(ctx, "[CampaignController][GetAllCampaign] get all campaign failed with error: %v", err)
		response.WriteError(ctx.Writer, baseEx.NotFound)
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// GetCampaign ...
//
//	@Summary		Get campaign
//	@Description	Return campaign
//	@Tags			campaign
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			code	path		string	true	"code"
//	@Success		200		{object}	response.Response{data=resources.Campaign}
//	@Failure		500		{object}	response.Response
//	@Router			/v1/campaigns/{code} [get]
func (r *CampaignController) GetCampaign(ctx *gin.Context) {
	code := strings.TrimSpace(ctx.Param("code"))
	result, err := r.CampaignService.GetCampaign(ctx, code)
	if err != nil {
		if errors.Is(err, constants.ErrOrmRecordNotFound) {
			log.Warn(ctx, "[CampaignController][GetCampaign] campaign not found with code %+v", code)
			response.WriteError(ctx.Writer, baseEx.NotFound)
			return
		}
		log.Error(ctx, "[CampaignController][GetCampaign] get campaign %+v failed with err: %+v", code, err)
		response.WriteError(ctx.Writer, baseEx.SystemError)
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// CreateCampaign ...
//
//	@Summary		Create Campaign
//	@Description	Use this API to create campaign
//	@Tags			campaign
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resources.UpsertCampaignRequest	true	"Request body""
//	@Success		200		{object}	response.Response{data=resources.Campaign}
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Router			/v1/campaigns [post]
func (r *CampaignController) CreateCampaign(ctx *gin.Context) {
	var request resources.UpsertCampaignRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			ctx.Abort()
			return
		}
		log.Warn(ctx, "[CampaignController][CreateCampaign] cannot bind request with error: %s", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		ctx.Abort()
		return
	}
	result, err := r.CampaignService.CreateCampaign(ctx, mappers.ToCampaignEntity(&request))
	if err != nil {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, err.Error()))
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// UpdateCampaign ...
//
//	@Summary		Update Campaign
//	@Description	Use this API to update campaign
//	@Tags			campaign
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string									true	"campaign code"
//	@Param			request	body		resources.UpsertCampaignRequest	true	"Request body""
//	@Success		200		{object}	response.Response{data=resources.Campaign}
//	@Failure		404		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/v1/campaigns/{code} [put]
func (r *CampaignController) UpdateCampaign(ctx *gin.Context) {
	code := strings.TrimSpace(ctx.Param("code"))
	var request resources.UpsertCampaignRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			ctx.Abort()
			return
		}
		log.Warn(ctx, "[CampaignController][UpdateCampaign] cannot bind request with error: %s", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		ctx.Abort()
		return
	}
	if code != request.Code {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, "code not matched"))
		ctx.Abort()
		return
	}
	result, err := r.CampaignService.UpdateCampaign(ctx, mappers.ToCampaignEntity(&request))
	if err != nil {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, err.Error()))
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}
