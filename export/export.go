package main

import (
	"C"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	rootUrl    = "https://app.candyhouse.co/api/sesame2/"
	cmd_unlock = 83
	cmd_lock   = 82
	src        = []byte("by Felica")
	history    = base64.StdEncoding.EncodeToString(src)
)

type RequestBody struct {
	Cmd     int    `json:"cmd"`
	History string `json:"history"`
	Sign    string `json:"sign"`
}

type ResponseBody struct {
	BatteryPercentage int     `json:"batteryPercentage"`
	BatteryVoltage    float64 `json:"batteryVoltage"`
	Position          int     `json:"position"`
	CHSesame2Status   string  `json:"CHSesame2Status"`
	Timestamp         int     `json:"timestamp"`
	Wm2State          bool    `json:"wm2State"`
}

func main() {
}

//export executeSesame3
func executeSesame3(signPtr, apiPtr, uuidPtr *C.char) {
	// sign,api,uuidはpython側から入力されるSIGN,API_TOKEN.UUIDに一致する
	sign := C.GoString(signPtr)
	api := C.GoString(apiPtr)
	uuid := C.GoString(uuidPtr)
	// fetchStatusでは鍵の状態を読みこんでいる
	// fetchStatusの戻り値はバイト列
	statusResponse := fetchStatus(api, uuid)
	// 鍵の状態のみをkey_statusとして取り出す
	key_status := string(statusResponse.CHSesame2Status)
	//executeLockで施錠を、executeUnlockで解錠を行う
	if isUnlocked(key_status) {
		fmt.Println("Key is " + key_status + ". Locking ...")
		executeResponse := executeLock(sign, api, uuid)
		fmt.Println(executeResponse)
	} else {
		fmt.Println("Key is " + key_status + ". Unlocking ...")
		executeResponse := executeUnlock(sign, api, uuid)
		fmt.Println(executeResponse)
	}

}

func isUnlocked(status string) bool {
	return status == "unlocked"
}

//export fetchStatus
func fetchStatus(api, uuid string) ResponseBody {
	// candyhouse公式(https://doc.candyhouse.co/ja/SesameAPI)に記載されているurlを準備する
	fetchUrl := rootUrl + uuid
	req, err := http.NewRequest("GET", fetchUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	// headerにpython側から受け取ったAPI_TOKENを渡す
	req.Header.Set("x-api-key", api)
	// リクエストの実行
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var statusResponse ResponseBody
	if err := json.Unmarshal(respbody, &statusResponse); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}
	return statusResponse
}

//export executeUnlock
func executeUnlock(sign, api, uuid string) string {
	// candyhouse公式(https://doc.candyhouse.co/ja/SesameAPI)に記載されているurlを準備する
	cmdUrl := rootUrl + uuid + "/cmd"
	// リクエスト構造体の初期化
	requestBody := RequestBody{
		Cmd:     cmd_unlock,
		History: history,
		Sign:    sign,
	}
	// リクエスト構造体をjson化してPOSTのbodyに追加する
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", cmdUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	// headerにpython側から受け取ったAPI_TOKENを渡す
	req.Header.Set("x-api-key", api)
	// リクエストの実行
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	return "Unlock command was executed."
}

//export executeLock
func executeLock(sign, api, uuid string) string {
	// candyhouse公式(https://doc.candyhouse.co/ja/SesameAPI)に記載されているurlを準備する
	cmdUrl := rootUrl + uuid + "/cmd"
	// リクエスト構造体の初期化
	requestBody := RequestBody{
		Cmd:     cmd_lock,
		History: history,
		Sign:    sign,
	}
	// リクエスト構造体をjson化してPOSTのbodyに追加する
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", cmdUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	// headerにpython側から受け取ったAPI_TOKENを渡す
	req.Header.Set("x-api-key", api)
	// リクエストの実行
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	return "Lock command was executed."
}
