package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	Timezone string `env:"TIMEZONE,required"`

	// MongoDB config
	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName             string `env:"MONGODB_NAME" envDefault:"go_project"`
	MongoDBStaffTableName   string `env:"MONGODB_STAFF_TABLE_NAME" envDefault:"staff_test"`
	MongoDBCompanyTableName string `env:"MONGODB_COMPANY_TABLE_NAME" envDefault:"company_tes"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
