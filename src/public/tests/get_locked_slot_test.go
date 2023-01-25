package tests

import (
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns_manager_autogen/adapter/repositories/mysql/models"
	"net/http"
	"testing"
)

func TestGetLockedSlot_ShouldReturnSuccess(t *testing.T) {
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
	golibtest.NewRestAssured(t).
		When().
		Get(url).Then().Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.code", "CODE_1").
		Body("data.name", "Name_1")
}
