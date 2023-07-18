package Modify

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var hash = "ac85e83127e49ec42487f272d9b9db8b"

func HashCheck(s string) bool {
	body, err := ioutil.ReadFile("banners/standard.txt")
	if err != nil {
		log.Fatal("unable to read!!", err)
	}

	if GetMD5Hash(string(body)) != hash {
		fmt.Println("wrong hash")
		os.Exit(0)
	}
	return true
}

func GetMD5Hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
