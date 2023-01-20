package main

import (
	golibmsg "gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/technixo/backend/campaigns_manager/public/bootstrap"

	"go.uber.org/fx"
)

//	@title		API
//	@version	1.0.0

// @host		apex-qc.nonprodposi.com
// @BasePath	/campaigns-manager
func main() {
	fx.New(
		bootstrap.AutoConfig(),
		bootstrap.All(),
		golibmsg.KafkaProducerOpt(),
	).Run()
}
