package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dedinirtadinata/bpjs-helper/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetSignature(isecret string, imessage string) string {
	secret := []byte(isecret)
	message := []byte(imessage)

	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))

	// to base64
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func GetTimestamp() string {
	currentUnix := time.Now()
	layout := "01/02/2006 3:04:05 PM"
	lastTime, err := time.Parse(layout, "01/01/1970 0:00:00 AM")
	if err != nil {
		fmt.Println(err)
	}

	timestamp := strconv.FormatInt(currentUnix.Unix()-lastTime.Unix(), 10)
	return timestamp
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PostData(endpoint string, postModel interface{}, model interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.userkey")
	host := viper.GetString("bpjs.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	jsonReq, err := json.Marshal(postModel)

	client := &http.Client{}

	req, err := http.NewRequest("POST", host+"/"+endpoint, bytes.NewBuffer(jsonReq))
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	reqbodyString := string(jsonReq)
	fmt.Println("Request: " + reqbodyString)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("Response: " + bodyString)
	var result models.BPJSResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == "200" {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		decrypt, gagal := AESCBCDecrypt(result.Response, key)
		//
		if gagal != nil {
			log.Println(gagal)
			return models.ResponseModel{Status: "OK", Data: result.Response}
		}

		bodyString := string(decrypt)
		fmt.Println("Response" + bodyString)

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		return models.ResponseModel{Status: "OK", Data: model}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func UpdateData(endpoint string, postModel interface{}, model interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.userkey")
	host := viper.GetString("bpjs.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	jsonReq, err := json.Marshal(postModel)

	client := &http.Client{}

	req, err := http.NewRequest("PUT", host+"/"+endpoint, bytes.NewBuffer(jsonReq))
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	reqbodyString := string(jsonReq)
	fmt.Println("Request: " + reqbodyString)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)

	var result models.BPJSResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == "200" {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		decrypt, gagal := AESCBCDecrypt(result.Response, key)
		//
		if gagal != nil {
			log.Println(gagal)
			return models.ResponseModel{Status: "OK", Data: result.Response}
		}

		bodyString := string(decrypt)
		fmt.Println("Response" + bodyString)

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		return models.ResponseModel{Status: "OK", Data: model}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func DeleteData(endpoint string, postModel interface{}, model interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.userkey")
	host := viper.GetString("bpjs.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	jsonReq, err := json.Marshal(postModel)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", host+"/"+endpoint, bytes.NewBuffer(jsonReq))
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	reqbodyString := string(jsonReq)
	fmt.Println("Request: " + reqbodyString)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)

	var result models.BPJSResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == "200" {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		decrypt, gagal := AESCBCDecrypt(result.Response, key)
		//
		if gagal != nil {
			log.Println(gagal)
			return models.ResponseModel{Status: "OK", Data: result.Response}
		}

		bodyString := string(decrypt)
		fmt.Println("Response " + bodyString)

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		return models.ResponseModel{Status: "OK", Data: model}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func GetData(endpoint string, model interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.userkey")
	host := viper.GetString("bpjs.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"/"+endpoint, nil)
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)

	var result models.BPJSResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == "200" {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		decrypt, gagal := AESCBCDecrypt(result.Response, key)
		//
		if gagal != nil {
			log.Println(gagal)
			return models.ResponseModel{Status: "NOT_OK", Message: "Gagal mendekripsikan data"}
		}

		bodyString := string(decrypt)
		fmt.Println(bodyString)

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		return models.ResponseModel{Status: "OK", Data: model}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func Get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func PostDataAplicare(endpoint string, postModel interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.antrean.userkey")
	host := viper.GetString("bpjs.aplicares.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	jsonReq, err := json.Marshal(postModel)

	client := &http.Client{}

	req, err := http.NewRequest("POST", host+"/"+endpoint, bytes.NewBuffer(jsonReq))
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	var result models.BPJSResponse2
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == 1 {
		return models.ResponseModel{Status: "OK", Message: result.Metadata.Message}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func PostDataAntrean(endpoint string, postModel interface{}) models.ResponseModel {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.antrean.userkey")
	host := viper.GetString("bpjs.antrean.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	jsonReq, err := json.Marshal(postModel)

	bodyString := string(jsonReq)
	fmt.Println("Request to BPJS :" + bodyString)

	client := &http.Client{}

	req, err := http.NewRequest("POST", host+"/"+endpoint, bytes.NewBuffer(jsonReq))
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString = string(bodyBytes)
	fmt.Println("Response BPJS :" + bodyString)

	var result models.BPJSResponse2
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == 200 {
		return models.ResponseModel{Status: "OK"}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}

func GetDataBPJS(codeSuccess int, service string, endpoint string, model interface{}) models.ResponseModel {
	var userkey, host string
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")

	if service == "antrean" {
		userkey = viper.GetString("bpjs.antrean.userkey")
		host = viper.GetString("bpjs.antrean.host")
	} else if service == "aplicares" {
		userkey = viper.GetString("bpjs.userkey")
		host = viper.GetString("bpjs.aplicares.host")
	} else {
		userkey = viper.GetString("bpjs.userkey")
		host = viper.GetString("bpjs.host")
	}

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"/"+endpoint, nil)
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var result models.BPJSResponse2
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		//fmt.Println("Can not unmarshal JSON")
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)
	}

	if result.Metadata.Code == codeSuccess {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		fmt.Println(result.Response)
		decrypt, gagal := AESCBCDecrypt(result.Response, key)

		//
		if gagal != nil {
			return models.ResponseModel{Status: "NOT_OK", Message: gagal.Error()}
		}

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
			return models.ResponseModel{Status: "NOT_OK", Message: "Gagal mendekripsikan data"}
		}

		return models.ResponseModel{Status: "OK", Data: model}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}
func GetDataBPJSAplicares(codeSuccess int, service string, endpoint string, model interface{}) models.ResponseModel {
	var userkey, host string
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")

	if service == "antrean" {
		userkey = viper.GetString("bpjs.antrean.userkey")
		host = viper.GetString("bpjs.antrean.host")
	} else if service == "aplicares" {
		userkey = viper.GetString("bpjs.userkey")
		host = viper.GetString("bpjs.aplicares.host")
	} else {
		userkey = viper.GetString("bpjs.userkey")
		host = viper.GetString("bpjs.host")
	}

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"/"+endpoint, nil)
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return models.ResponseModel{Status: "NOT_OK", Message: err.Error()}
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var result models.BPJSResponse3
	result.Response = model
	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		//fmt.Println("Can not unmarshal JSON")
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)
	}

	if result.Metadata.Code == codeSuccess {
		return models.ResponseModel{Status: "OK", Data: result.Response}
	} else {
		//data gagal di dapat
		return models.ResponseModel{Status: "NOT_OK", Message: result.Metadata.Message}
	}
}
func GetDataAntrean(endpoint string, model interface{}) (interface{}, error) {
	consId := viper.GetString("bpjs.consid")
	secret := viper.GetString("bpjs.secret")
	userkey := viper.GetString("bpjs.antrean.userkey")
	host := viper.GetString("bpjs.antrean.host")

	timestamp := GetTimestamp()
	salt := consId + "&" + timestamp

	signature := GetSignature(secret, salt)

	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"/"+endpoint, nil)
	// ...
	req.Header.Add("X-cons-id", consId)
	req.Header.Add("X-timestamp", timestamp)
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", userkey)
	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)

	var result models.BPJSResponse2

	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to go struct pointer
		logrus.Error(err)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if result.Metadata.Code == 1 {
		//data ada dan lakukan dekripsi
		key := consId + secret + timestamp
		decrypt, gagal := AESCBCDecrypt(result.Response, key)

		if gagal != nil {
			return nil, gagal
		}

		if err := json.Unmarshal([]byte(decrypt), &model); err != nil { // Parse []byte to go struct pointer
			return nil, err
		}

		return model, nil
	} else {
		return nil, errors.New(result.Metadata.Message)
	}

}
func Post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}
