package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func ConvertHex(HexInput string) string {

	//Decode Hex to string. Log errors to console
	DecodedHex, err := hex.DecodeString(HexInput)
	if err != nil {
		log.Fatal(err)
	}

	//Encode string to Base64. Return data
	EncodedBase64 := base64.StdEncoding.EncodeToString(DecodedHex)
	return EncodedBase64
}
