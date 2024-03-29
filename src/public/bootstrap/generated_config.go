// Code generated by cli. DO NOT EDIT.

package bootstrap

import (
	"github.com/go-resty/resty/v2"
	"gitlab.com/golibs-starter/golib"
	golibcron "gitlab.com/golibs-starter/golib-cron"
	golibdata "gitlab.com/golibs-starter/golib-data"
	golibgin "gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/properties"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/routers"
	"go.uber.org/fx"
	"net/http"
)

func GeneratedConfig() fx.Option {
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
		golib.ProvideProps(properties.NewSwaggerProperties),
		golibmsg.KafkaCommonOpt(),
		golibdata.DatasourceOpt(),
		golibmsg.KafkaConsumerOpt(),
		golibcron.EnableCron(),
		golibdata.RedisOpt(),
		fx.Provide(func() *http.Client { return &http.Client{} }),
		fx.Provide(func(client *http.Client) *resty.Client {
			return resty.NewWithClient(client)
		}),
		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		// golib.OnStopEventOpt() MUST be first
		golib.OnStopEventOpt(),
		golibgin.OnStopHttpServerOpt(),
		golibmsg.OnStopConsumerOpt(),
	)
}
