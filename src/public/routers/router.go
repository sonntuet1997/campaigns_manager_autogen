// Source code template generated by https://gitlab.com/technixo/tnx

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/golibs-starter/golib"
	golibgin "gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/golibs-starter/golib/web/actuator"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/controllers"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/docs"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/properties"
	"go.uber.org/fx"
)

// RegisterRoutersIn represents constructor params for fx
type RegisterRoutersIn struct {
	fx.In
	App                *golib.App
	Engine             *gin.Engine
	SwaggerProps       *properties.SwaggerProperties
	Actuator           *actuator.Endpoint
	CampaignController *controllers.CampaignController

	LockedSlotController *controllers.LockedSlotController
}

// RegisterHandlers register handlers
func RegisterHandlers(app *golib.App, engine *gin.Engine) {
	engine.Use(golibgin.InitContext())
	engine.Use(golibgin.WrapAll(app.Handlers())...)
}

// RegisterGinRouters register gin routes
func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))
	if p.SwaggerProps.Enabled {
		docs.SwaggerInfo.BasePath = p.App.Path()
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	v1 := p.Engine.Group(fmt.Sprintf("%s/v1", p.App.Path()))

	campaigns := v1.Group("campaigns")
	{
		campaigns.GET("", p.CampaignController.GetAllCampaign)
		campaigns.GET(":code", p.CampaignController.GetCampaign)
		campaigns.POST("", p.CampaignController.CreateCampaign)
		campaigns.PUT(":code", p.CampaignController.UpdateCampaign)
	}
	lockedSlots := v1.Group("locked-slots")
	{
		lockedSlots.GET("", p.LockedSlotController.GetAllLockedSlot)
		lockedSlots.GET(":code", p.LockedSlotController.GetLockedSlot)
		lockedSlots.POST("", p.LockedSlotController.CreateLockedSlot)
		lockedSlots.PUT(":code", p.LockedSlotController.UpdateLockedSlot)
	}
}
