package bootstrap

import (
	"gitlab.com/golibs-starter/golib"
	golibdata "gitlab.com/golibs-starter/golib-data"
	golibgin "gitlab.com/golibs-starter/golib-gin"
	"go.uber.org/fx"
)

// All register all constructors for fx container
func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),

		golib.EventOpt(),

		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golibgin.GinHttpServerOpt(),
		fx.Invoke(routers.RegisterHandlers),
		fx.Invoke(routers.RegisterGinRouters),

		// Provide all application properties

		// Provide datasource auto config
		golibdata.DatasourceOpt(),
		// Provide port's implements

		// Provide controllers, these controllers will be used
		// when register router was invoked

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		// golib.OnStopEventOpt() MUST be first
		golib.OnStopEventOpt(),
		golibgin.OnStopHttpServerOpt(),
	)
}
