package notification

import (
	"bytes"
	"net/http"

	"github.com/spf13/viper"
)

func sendPushNotiByPlayerID(playerIDs []string, data string, message string) {
	url := "https://onesignal.com/api/v1/notifications"
	viper.AutomaticEnv()
	appID := viper.GetString("OSAPPID")
	APIKEY := viper.GetString("OSAPIKEY")
	var jsonStr = []byte(`{
		"app_id": "` + appID + `",
		"include_player_ids": [` + genPlayerIDS(playerIDs) + `],
		"data": ` + data + `,
		"contents": {"en": "` + message + `"}
	  }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Basic "+APIKEY)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func sendPushNoti() {

}

func genPlayerIDS(arry []string) string {
	mySTR := ``
	for i, str := range arry {
		mySTR = mySTR + `'` + str + `'`
		if i < len(arry)-1 {
			mySTR = mySTR + `,`
		}
	}
	return mySTR
}
