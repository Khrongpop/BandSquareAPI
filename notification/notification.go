package notification

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khrongpop/BandSquareAPI/model"

	"github.com/spf13/viper"
)

func SendPushNotiByPlayerID(playerIDs []model.PlayerID, data string, message string) {
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
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Basic "+APIKEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func sendPushNoti() {

}

func genPlayerIDS(players []model.PlayerID) string {
	mySTR := ``
	for i, player := range players {
		mySTR = mySTR + `"` + player.PlayerID + `"`
		if i < len(players)-1 {
			mySTR = mySTR + `,`
		}
	}
	return mySTR
}
