package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/public/resources"
	"net/http"
	"testing"
)

func TestCreateCampaign_ShouldReturnSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("campaigns")
	url := `/v1/campaigns`
	requestBody := &resources.UpsertCampaignRequest{
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
	golibdataTestUtil.AssertDatabaseHas(t, "campaigns", map[string]interface{}{
		"code": "CODE_1",
		"name": "Name_1",
	})
}
