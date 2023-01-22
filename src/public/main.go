package main

import (
	"gitlab.com/technixo/backend/campaigns-manager/public/bootstrap"
	"go.uber.org/fx"
)

//	@title		API
//	@version	1.0.0

// @host		apex-qc.nonprodposi.com
// @BasePath	/campaigns-manager
func main() {
	fx.New(
		bootstrap.All(),
	).Run()
}
