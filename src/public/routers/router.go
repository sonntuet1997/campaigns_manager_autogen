// Source code template generated by https://gitlab.com/technixo/tnx

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/golibs-starter/golib"
	golibgin "gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/technixo/backend/campaigns_manager/public/properties"
	"gitlab.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

// RegisterRoutersIn represents constructor params for fx
type RegisterRoutersIn struct {
	fx.In
	App          *golib.App
	Engine       *gin.Engine
	SwaggerProps *properties.SwaggerProperties
	Actuator     *actuator.Endpoint
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
	_ = v1
}
