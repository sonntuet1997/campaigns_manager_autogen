// Code generated by cli. DO NOT EDIT.

package bootstrap

import (
	golibcron "gitlab.com/golibs-starter/golib-cron"
	golibmsg "gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/core/usecases"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/consumers"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/controllers"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/jobs"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/services"
	"go.uber.org/fx"
)

func GeneratedModule() fx.Option {
	return fx.Options(
		fx.Provide(mysql.NewCampaignRepositoryAdapter),
		fx.Provide(usecases.NewGetCampaignUsecase),
		fx.Provide(usecases.NewCreateCampaignUsecase),
		fx.Provide(usecases.NewUpdateCampaignUsecase),
		fx.Provide(services.NewCampaignService),
		fx.Provide(controllers.NewCampaignController),

		fx.Provide(mysql.NewLockedSlotRepositoryAdapter),
		fx.Provide(usecases.NewGetLockedSlotUsecase),
		fx.Provide(usecases.NewCreateLockedSlotUsecase),
		fx.Provide(usecases.NewUpdateLockedSlotUsecase),
		fx.Provide(services.NewLockedSlotService),
		fx.Provide(controllers.NewLockedSlotController),

		golibcron.ProvideJob(jobs.NewLockedSlotJob),

		golibmsg.ProvideConsumer(consumers.NewLockedSlotConsumer),
	)
}
