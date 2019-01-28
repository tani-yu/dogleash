package auth

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

type authParams struct {
	APIKey string
	APPKey string
}

func GetDDClient() (*datadog.Client, error) {
	ap, err := getParamsFromConfigFile()
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	client := datadog.NewClient(ap.APIKey, ap.APPKey)

	return client, nil
}

// 認証情報を設定ファイルから取得する
func getParamsFromConfigFile() (authParams, error) {
	var ap authParams

	// key := "profiles." + cmd.ProfileName()
	// log.Printf(key + "\n")
	// fmt.Println(viper.AllSettings())

	// log.Printf("%s", viper.Get(key))
	//	if viper.IsSet(key) == false {
	//	return ap, fmt.Errorf("Profile Not Found: " + cmd.ProfileName())
	//}

	if err := viper.Unmarshal(&ap); err != nil {
		return ap, err
	}

	return ap, nil
}
