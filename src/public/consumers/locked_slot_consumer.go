package consumers

import (
	"context"
	"encoding/json"
	"gitlab.com/golibs-starter/golib-message-bus/kafka/core"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/publisher/kafka"
	"gitlab.com/technixo/backend/campaigns-manager/core/entities"
	"gitlab.com/technixo/backend/campaigns-manager/core/events"
	"gitlab.com/technixo/backend/campaigns-manager/public/services"
	"time"
)

type LockedSlotConsumer struct {
	LockedSlotService services.LockedSlotServiceable
}

func NewLockedSlotConsumer(LockedSlotService services.LockedSlotServiceable) core.ConsumerHandler {
	return &LockedSlotConsumer{
		LockedSlotService: LockedSlotService,
	}
}

func (r *LockedSlotConsumer) HandlerFunc(msg *core.ConsumerMessage) {
	var eventMessage kafka.Event[events.LockedSlotEventMessage]
	if err := json.Unmarshal(msg.Value, &eventMessage); err != nil {
		log.Errorf("[LockedSlotConsumer] error when unmarshal event message %+v, detail: %+v", msg.Value, err)
		return
	}
	log.Infof("[LockedSlotConsumer] received event at %+v with payload %+v", time.Now().UnixMilli(), eventMessage)
	payload := eventMessage.Payload()
	_, err := r.LockedSlotService.CreateLockedSlot(context.Background(), &entities.LockedSlot{
		Code: payload.Code,
		Name: payload.Name,
	})
	if err != nil {
		log.Errorf("[LockedSlotConsumer] create failed with error: %+v", err)
	}
}

func (r *LockedSlotConsumer) Close() {
	//
}
