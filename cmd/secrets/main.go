package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type GetSecretsResp struct {
	Data map[string]string `json:"data"`
}

func main() {
	resp := GetSecretsResp{}
	err := json.NewDecoder(os.Stdin).Decode(&resp)
	if err != nil {
		panic(err)
	}

	for k, v := range resp.Data {
		fmt.Println("\x1b[1;37m" + k + "\x1b[0m")
		dec, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dec))
	}
}
