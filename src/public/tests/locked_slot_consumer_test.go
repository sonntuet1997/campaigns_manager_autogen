package tests

import (
	"github.com/ecodia/golang-awaitility/awaitility"
	"github.com/stretchr/testify/assert"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibmsgTestUtil "gitlab.com/golibs-starter/golib-message-bus/testutil"
	"testing"
	"time"
)

func TestLockedSlotConsumer_ShouldSuccess(t *testing.T) {
	t.Run("given normal run when receiving 1 message should success", func(t *testing.T) {
		golibdataTestUtil.TruncateTables("locked-slots")
		eventMessage := `{
  "id": "9e01594c-8346-11ed-803d-1aa9442a23c3",
  "event": "LockedSlotEvent",
  "source": "not_used",
  "service_code": "",
  "timestamp": 1671857484019,
  "payload": {
    "code": "CODE_1",
    "name": "Name_1"
  }
}`
		golibmsgTestUtil.HandleMessage("LockedSlotConsumer", []byte(eventMessage))
		err := awaitility.Await(time.Second, 5*time.Second, func() bool {
			return golibdataTestUtil.CountWithQuery("locked-slots", map[string]interface{}{
				"code": "CODE_1",
				"name": "Name_1",
			}) == 1
		})
		assert.Nil(t, err)
	})
}
