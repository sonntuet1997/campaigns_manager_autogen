package tests

import (
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"net/http"
	"testing"
)

func TestGetCampaign_ShouldReturnSuccess(t *testing.T) {
	golibdataTestUtil.TruncateTables("campaigns")
	golibdataTestUtil.Insert([]*models.Campaign{
		{
			Code: "CODE_1",
			Name: "Name_1",
		},
		{
			Code: "CODE_2",
			Name: "Name_2",
		},
	})
	url := `/v1/campaigns/CODE_1`
	golibtest.NewRestAssured(t).
		When().
		Get(url).Then().Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.code", "CODE_1").
		Body("data.name", "Name_1")
}
