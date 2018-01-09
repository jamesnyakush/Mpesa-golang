package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Mpesa is mpesa
type Mpesa struct {
	AppKey    string
	AppSecret string
}

// New return a new Mpesa
func New(appKey, appSecret string) (*Mpesa, error) {
	return &Mpesa{appKey, appSecret}, nil
}

func (mpesa Mpesa) authenticate() (string, error) {
	appKeySecret := mpesa.AppKey + ":" + mpesa.AppSecret
	bytes := []byte(appKeySecret)

	encoded := base64.StdEncoding.EncodeToString(bytes)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials", strings.NewReader(encoded))
	if err != nil {
		return "", err
	}
	request.Header.Add("authorization", "Basic "+encoded)
	request.Header.Add("cache-control", "no-cache")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	object := make(map[string]interface{})
	json.NewDecoder(response.Body).Decode(&object)
	defer response.Body.Close()

	accessToken := object["access_token"].(string)
	log.Println(accessToken)
	return accessToken, nil
}

// C2BSimulation sends a new request
func (mpesa Mpesa) C2BSimulation(c2b C2B) (string, error) {
	var c2bs []C2B
	c2bs = append(c2bs, c2b)
	body, err := json.Marshal(c2bs)
	if err != nil {
		return "", err
	}

	url := "https://sandbox.safaricom.co.ke/safaricom/c2b/v1/simulate"
	return mpesa.newStringRequest(url, body, false)
}

// B2CRequest sends a new request
func (mpesa Mpesa) B2CRequest(b2c B2C) (string, error) {
	var b2cs []B2C
	b2cs = append(b2cs, b2c)

	body, err := json.Marshal(b2cs)
	if err != nil {
		return "", err
	}

	url := "https://sandbox.safaricom.co.ke/mpesa/b2c/v1/paymentrequest"
	return mpesa.newStringRequest(url, body, false)
}

// B2BRequest sends a new request
func (mpesa Mpesa) B2BRequest(b2b B2B) (string, error) {
	var b2bs []B2B
	b2bs = append(b2bs, b2b)

	body, err := json.Marshal(b2bs)
	if err != nil {
		return "", nil
	}

	url := "https://sandbox.safaricom.co.ke/safaricom/b2b/v1/paymentrequest"
	return mpesa.newStringRequest(url, body, false)
}

// STKPushSimulation sends an STK push?
func (mpesa Mpesa) STKPushSimulation(stkPush STKPush) (string, error) {
	var stkPushes []STKPush
	stkPushes = append(stkPushes, stkPush)

	body, err := json.Marshal(stkPushes)
	if err != nil {
		return "", nil
	}

	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	return mpesa.newStringRequest(url, body, false)
}

// STKPushTransactionStatus gets a status
func (mpesa Mpesa) STKPushTransactionStatus(stkPush STKPush) (string, error) {
	var stkPushes []STKPush
	stkPushes = append(stkPushes, stkPush)

	body, err := json.Marshal(stkPushes)
	if err != nil {
		return "", nil
	}

	url := "https://sandbox.safaricom.co.ke/mpesa/stkpushquery/v1/query"
	return mpesa.newStringRequest(url, body, true)
}

func (mpesa Mpesa) newStringRequest(url string, body []byte, noCache bool) (string, error) {
	auth, err := mpesa.authenticate()
	if err != nil {
		return "", nil
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	client := &http.Client{}
	request.Header.Set("content-type", "application/json")
	request.Header.Set("authorization", "Bearer "+auth)
	if noCache {
		request.Header.Set("cache-control", "no-cache")
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
