package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	appId         = "25dzX3vbYqktVxyX"
	deviceId      = "URJuHA0FgRACAW8eGUh/zJ+2"
	userId        = "31a1407588d94e47961094xxxxxxxxx"
	nonce         = 0
	publicKey     = ""
	signatureData = ""
)

func randomString(l int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		rand.NewSource(time.Now().UnixNano())
		bytes[i] = byte(randInt(1, 2^256-1))
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func AddHeader(req *http.Request) *http.Request {
	req.Header.Set("authorization", "Bearer xxxxxxxxxxxxxxx")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Add("origin", "https://aliyundrive.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("x-canary", "client=web,app=adrive,version=v3.17.0")
	req.Header.Add("x-device-id", deviceId)
	req.Header.Add("x-signature", signatureData)
	return req
}

func CreateSession() error {
	api := "https://api.aliyundrive.com/users/v1/users/device/create_session"
	form := fmt.Sprintf(`{"deviceName": "Edge浏览器","modelName": "Windows网页版","pubKey":"%s"}`, publicKey)
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(form))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req = AddHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(body), "true") {
		return nil
	}
	return errors.New(string(body))
}

func AliDownload(fileId string) error {
	api := "https://api.aliyundrive.com/v2/file/get_download_url"
	form := fmt.Sprintf(`{"expire_sec": 14400, "drive_id": "409666480","file_id": "636767a8287a5f3ae94b4bd480bf0e56dc294029"}`)
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(form))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req = AddHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(body), "url") {
		log.Println(string(body))
		return nil
	}
	return errors.New(string(body))
}

func InitAliKey() error {
	max := 32
	key := randomString(max)
	data := fmt.Sprintf("%s:%s:%s:%d", appId, deviceId, userId, nonce)
	var privKey = secp256k1.PrivKey(key)
	pubKey := privKey.PubKey()
	publicKey = "04" + hex.EncodeToString(pubKey.Bytes())
	signature, err := privKey.Sign([]byte(data))
	if err != nil {
		return err
	}
	signatureData = hex.EncodeToString(signature) + "01"
	return nil
}

func main() {
	err := InitAliKey()
	if err != nil {
		log.Println(err)
	}
	err = CreateSession()
	if err != nil {
		log.Println(err)
	}
	AliDownload("")
}
