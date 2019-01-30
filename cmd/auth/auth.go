package auth

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

type dataDog struct {
	DataDog authParams
}

type authParams struct {
	APIKey string
	APPKey string
}

func GetDDClient() (*datadog.Client, error) {
	conf, err := getParamsFromConfigFile()
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	client := datadog.NewClient(conf.DataDog.APIKey, conf.DataDog.APPKey)

	return client, nil
}

// 認証情報を設定ファイルから取得する
func getParamsFromConfigFile() (dataDog, error) {
	var conf dataDog

	if err := viper.Unmarshal(&conf); err != nil {
		return conf, err
	}

	return conf, nil
}
