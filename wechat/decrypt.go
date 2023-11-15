package main

import "fmt"

//https://t-static.yeahkagame.com/hgame/egg_t_1.5/index.html#/login?a=3904&s=ck8it4th7p26g7rfci1g
func Decrypt() {
	miniProgram := getMiniProgram()
	a := miniProgram.GetAuth()
	sessionResult, err := a.Code2Session(jscode)
	if err != nil {
		panic(err)
	}
	sessionKey = sessionResult.SessionKey
	fmt.Println("sessionKey: ", sessionKey)

	e := miniProgram.GetEncryptor()
	data, err := e.Decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypt data: %+v\n", data)
}
