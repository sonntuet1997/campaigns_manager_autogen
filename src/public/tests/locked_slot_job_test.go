package tests

import (
	golibcrontestsuite "gitlab.com/golibs-starter/golib-cron/testsuite"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	"testing"
)

func TestRunLockedSlotJob_ShouldSuccess(t *testing.T) {
	t.Run("given normal run should success", func(t *testing.T) {
		golibdataTestUtil.TruncateTables("locked-slots")
		golibcrontestsuite.RunJob("LockedSlotJob")
		golibdataTestUtil.AssertDatabaseHas(t, "locked-slots", map[string]interface{}{
			"code": "CODE_1",
			"name": "Name_1",
		})
	})
}
