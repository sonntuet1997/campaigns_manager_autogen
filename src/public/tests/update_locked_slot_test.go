package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"gitlab.com/technixo/backend/campaigns-manager/public/resources"
	"net/http"
	"testing"
)

func TestUpdateLockedSlot_ShouldReturnSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("locked-slots")
	golibdataTestUtil.Insert([]*models.LockedSlot{
		{
			Code: "CODE_1",
			Name: "Name_1",
		},
		{
			Code: "CODE_2",
			Name: "Name_2",
		},
	})
	url := `/v1/locked-slots/CODE_1`
	requestBody := &resources.UpsertLockedSlotRequest{
		Code: "CODE_1",
		Name: "Name_11",
	}
	requestString, err := json.Marshal(requestBody)
	assert.NoError(t, err)
	golibtest.NewRestAssured(t).
		When().
		Put(url).Body(string(requestString)).Then().Status(http.StatusOK).
		Body("meta.code", 200)
	golibdataTestUtil.AssertDatabaseHas(t, "locked-slots", map[string]interface{}{
		"code": "CODE_1",
		"name": "Name_11",
	})
}
