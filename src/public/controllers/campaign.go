package controllers

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/golibs-starter/golib/web/response"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources/mappers"
	"gitlab.com/technixo/backend/campaigns-manager/public/services"
	"net/http"
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
