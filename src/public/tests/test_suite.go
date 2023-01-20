package tests

import (
	"context"
	"gitlab.com/golibs-starter/golib"
	golibdataTestUtil "gitlab.com/golibs-starter/golib-data/testutil"
	golibmsgTestUtil "gitlab.com/golibs-starter/golib-message-bus/testutil"
	golibmigrate "gitlab.com/golibs-starter/golib-migrate"
	golibtest "gitlab.com/golibs-starter/golib-test"
	golibkafkatestsuite "gitlab.com/golibs-starter/kafka-test-suite"
	"gitlab.com/technixo/backend/campaigns_manager/public/bootstrap"
	"go.uber.org/fx"
)

func init() {
	err := fx.New(
		bootstrap.All(),
		golib.ProvidePropsOption(golib.WithActiveProfiles([]string{"testing"})),
		golib.ProvidePropsOption(golib.WithPaths([]string{"../config/"})),
		golibmigrate.MigrationOpt(),
		golibdataTestUtil.EnableDatabaseTestUtilOpt(),
		golibtest.EnableWebTestUtil(),
		golibmsgTestUtil.EnableKafkaConsumerTestUtil(),
		golibkafkatestsuite.EnableKafkaProducerTestSuite(),
	).Start(context.Background())
	if err != nil {
		panic(err)
	}
}
