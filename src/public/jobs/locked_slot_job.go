package jobs

import (
	"context"
	"github.com/robfig/cron/v3"
	"gitlab.com/golibs-starter/golib/log"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/public/services"
	"go.uber.org/fx"
)

type LockedSlotJob struct {
	LockedSlotService services.LockedSlotServiceable
}

type LockedSlotJobParams struct {
	fx.In
	LockedSlotService services.LockedSlotServiceable
}

func NewLockedSlotJob(params LockedSlotJobParams) cron.Job {
	return &LockedSlotJob{
		LockedSlotService: params.LockedSlotService,
	}
}

func (r *LockedSlotJob) Run() {
	log.Debugf("[LockedSlotJob] job start")
	entity := entities.LockedSlot{
		Code: "CODE_1",
		Name: "Name_1",
	}
	_, err := r.LockedSlotService.CreateLockedSlot(context.Background(), &entity)
	if err != nil {
		log.Errorf("[LockedSlotJob] create failed with error: %+v", err)
	}
	log.Debugf("[LockedSlotJob] job stop")
}
