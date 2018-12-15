package main

import (
	"encoding/json"
	"fmt"
)

var (
	gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
	accessKeyId     = "LTAIWxWZXPLYT5px"
	accessKeySecret = "tO11eWbHMR1zlzc3M7OvbEXqs74jPf"
	phoneNumbers    = "15088131357"
	signName        = "刘佛添"
	templateCode    = "SMS_139810284"
	templateParam   = "{\"code\":\"1234\"}"
)

func main() {
	smsClient := aliyunsmsclient.New(gatewayUrl)
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	fmt.Println("Got raw response from server:", string(result.RawResponse))
	if err != nil {
		panic("Failed to send Message: " + err.Error())
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	if result.IsSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
	} else {
		fmt.Println("Failed to send a SMS:", resultJson)
	}
}
