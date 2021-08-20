package main

import (
	"C"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var rootUrl = "https://app.candyhouse.co/api/sesame2/"

//export fetch
func fetch(api, uuid string) []byte {
	fetchUrl := rootUrl + uuid
	req, _ := http.NewRequest("GET", fetchUrl, nil)
	req.Header.Set("x-api-key", api)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

type RequestBody struct {
	Cmd     int    `json:"cmd"`
	History string `json:"history"`
	Sign    string `json:"sign"`
}

//export unlock
func unlock(sign, api, uuid string) int {
	src := []byte("By IC")
	enchistory := base64.StdEncoding.EncodeToString(src)
	requestBody := RequestBody{
		Cmd:     83,
		History: enchistory,
		Sign:    sign,
	}

	jsonString, _ := json.Marshal(requestBody)
	fmt.Printf("[+] %s\n", string(jsonString))
	cmdUrl := rootUrl + uuid + "/cmd"
	req, err := http.NewRequest("POST", cmdUrl, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}
	req.Header.Set("x-api-key", api)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}
	fmt.Printf("%#v", string(byteArray))
	return 0
}

//export lock
func lock(sign, api, uuid string) int {
	src := []byte("By IC")
	enchistory := base64.StdEncoding.EncodeToString(src)
	requestBody := RequestBody{
		Cmd:     82,
		History: enchistory,
		Sign:    sign,
	}

	jsonString, _ := json.Marshal(requestBody)
	fmt.Printf("[+] %s\n", string(jsonString))

	cmdUrl := rootUrl + uuid + "/cmd"
	req, err := http.NewRequest("POST", cmdUrl, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}
	req.Header.Set("x-api-key", api)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	fmt.Printf("%#v", string(byteArray))
	return 0
}

type AutoGenerated struct {
	BatteryPercentage int     `json:"batteryPercentage"`
	BatteryVoltage    float64 `json:"batteryVoltage"`
	Position          int     `json:"position"`
	CHSesame2Status   string  `json:"CHSesame2Status"`
	Timestamp         int     `json:"timestamp"`
	Wm2State          bool    `json:"wm2State"`
}

//export operation
func operation(sign, api, uuid string) {
	//op := unlock
	Unlock := unlock
	Lock := lock
	//op(sign)
	jsonStr := fetch(api, uuid)
	//fmt.Print(jsonStr)

	jsonBytes := jsonStr //([]byte)(jsonStr)
	data := new(AutoGenerated)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	key_status := string(data.CHSesame2Status)
	//fmt.Println(key_status)
	if key_status == "unlocked" {
		fmt.Print("Key is " + key_status + ". Locking ...")
		Lock(sign, api, uuid)
	} else {
		fmt.Print("Key is " + key_status + ". Unlocking ...")
		Unlock(sign, api, uuid)
	}

}

func main() {
}
