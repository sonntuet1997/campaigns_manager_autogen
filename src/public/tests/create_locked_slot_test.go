package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"net/http"
	"testing"
)

func TestCreateLockedSlot_ShouldReturnSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("locked-slots")
	url := `/v1/locked-slots`
	requestBody := &resources.UpsertLockedSlotRequest{
		Code: "CODE_1",
		Name: "Name_1",
	}
	requestString, err := json.Marshal(requestBody)
	assert.NoError(t, err)
	golibtest.NewRestAssured(t).
		When().
		Post(url).Body(string(requestString)).Then().Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.code", "CODE_1").
		Body("data.name", "Name_1")
	golibdataTestUtil.AssertDatabaseHas(t, "locked-slots", map[string]interface{}{
		"code": "CODE_1",
		"name": "Name_1",
	})
}
