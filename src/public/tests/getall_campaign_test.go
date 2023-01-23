package tests

import (
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"net/http"
	"testing"
)

func TestGetAllCampaign_ShouldReturnSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("campaigns")
	golibdataTestUtil.Insert([]*models.Campaign{
		{
			Code: "CODE_1",
		},
		{
			Code: "CODE_2",
		},
	})
	url := `/v1/campaigns`
	golibtest.NewRestAssured(t).
		When().
		Get(url).Then().Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.#", 2).
		Body("data.0.code", "CODE_1").
		Body("data.1.code", "CODE_2")
}
