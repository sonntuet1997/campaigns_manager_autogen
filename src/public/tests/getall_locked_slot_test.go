package tests

import (
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibtest "gitlab.com/golibs-starter/golib-test"
	"gitlab.com/technixo/backend/campaigns-manager/adapter/repositories/mysql/models"
	"net/http"
	"testing"
)

func TestGetAllLockedSlot_ShouldReturnSuccess(t *testing.T) {
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
	url := `/v1/locked-slots`
	golibtest.NewRestAssured(t).
		When().
		Get(url).Then().Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.#", 2).
		Body("data.0.code", "CODE_1").
		Body("data.0.name", "Name_1").
		Body("data.1.code", "CODE_2").
		Body("data.1.name", "Name_2")
}
