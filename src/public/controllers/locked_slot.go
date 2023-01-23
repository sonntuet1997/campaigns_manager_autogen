package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/golibs-starter/golib/web/response"
	"gitlab.com/technixo/backend/campaigns-manager/core/constants"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources/mappers"
	"gitlab.com/technixo/backend/campaigns-manager/public/services"
	"net/http"
	"strings"
)

type LockedSlotController struct {
	LockedSlotService services.LockedSlotServiceable
}

func NewLockedSlotController(
	LockedSlotService services.LockedSlotServiceable,
) *LockedSlotController {
	return &LockedSlotController{
		LockedSlotService: LockedSlotService,
	}
}

// GetAllLockedSlot ...
//
//	@Summary		Get List locked slot
//	@Description	Return list of locked slot
//	@Tags			locked slot
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]resources.LockedSlot}
//	@Failure		500	{object}	response.Response
//	@Router			/v1/locked-slots [get]
func (r *LockedSlotController) GetAllLockedSlot(ctx *gin.Context) {
	var request resources.GetLockedSlotRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			return
		}
		log.Warn(ctx, "[LockedSlotController][GetAllLockedSlot] cannot bind request error: [%s]", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		return
	}
	result, err := r.LockedSlotService.GetAllLockedSlot(ctx, mappers.ToLockedSlotFilter(&request))
	if err != nil {
		log.Error(ctx, "[LockedSlotController][GetAllLockedSlot] get all locked slot failed with error: %v", err)
		response.WriteError(ctx.Writer, baseEx.NotFound)
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// GetLockedSlot ...
//
//	@Summary		Get locked slot
//	@Description	Return locked slot
//	@Tags			locked slot
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			code	path		string	true	"code"
//	@Success		200		{object}	response.Response{data=resources.LockedSlot}
//	@Failure		500		{object}	response.Response
//	@Router			/v1/locked-slots/{code} [get]
func (r *LockedSlotController) GetLockedSlot(ctx *gin.Context) {
	code := strings.TrimSpace(ctx.Param("code"))
	result, err := r.LockedSlotService.GetLockedSlot(ctx, code)
	if err != nil {
		if errors.Is(err, constants.ErrOrmRecordNotFound) {
			log.Warn(ctx, "[LockedSlotController][GetLockedSlot] locked slot not found with code %+v", code)
			response.WriteError(ctx.Writer, baseEx.NotFound)
			return
		}
		log.Error(ctx, "[LockedSlotController][GetLockedSlot] get locked slot %+v failed with err: %+v", code, err)
		response.WriteError(ctx.Writer, baseEx.SystemError)
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// CreateLockedSlot ...
//
//	@Summary		Create LockedSlot
//	@Description	Use this API to create locked slot
//	@Tags			locked slot
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resources.UpsertLockedSlotRequest	true	"Request body""
//	@Success		200		{object}	response.Response{data=resources.LockedSlot}
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Router			/v1/locked-slots [post]
func (r *LockedSlotController) CreateLockedSlot(ctx *gin.Context) {
	var request resources.UpsertLockedSlotRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			ctx.Abort()
			return
		}
		log.Warn(ctx, "[LockedSlotController][CreateLockedSlot] cannot bind request with error: %s", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		ctx.Abort()
		return
	}
	result, err := r.LockedSlotService.CreateLockedSlot(ctx, mappers.ToLockedSlotEntity(&request))
	if err != nil {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, err.Error()))
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}

// UpdateLockedSlot ...
//
//	@Summary		Update LockedSlot
//	@Description	Use this API to update locked slot
//	@Tags			locked slot
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string									true	"campaign code"
//	@Param			request	body		resources.UpsertLockedSlotRequest	true	"Request body""
//	@Success		200		{object}	response.Response{data=resources.LockedSlot}
//	@Failure		404		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/v1/locked-slots/{code} [put]
func (r *LockedSlotController) UpdateLockedSlot(ctx *gin.Context) {
	code := strings.TrimSpace(ctx.Param("code"))
	var request resources.UpsertLockedSlotRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, validationErrors[0].Error()))
			ctx.Abort()
			return
		}
		log.Warn(ctx, "[LockedSlotController][UpdateLockedSlot] cannot bind request with error: %s", err)
		response.WriteError(ctx.Writer, baseEx.BadRequest)
		ctx.Abort()
		return
	}
	if code != request.Code {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, "code not matched"))
		ctx.Abort()
		return
	}
	result, err := r.LockedSlotService.UpdateLockedSlot(ctx, mappers.ToLockedSlotEntity(&request))
	if err != nil {
		response.WriteError(ctx.Writer, baseEx.New(http.StatusBadRequest, err.Error()))
		return
	}
	response.Write(ctx.Writer, response.Ok(result))
}
