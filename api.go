package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Env is the environment type
type Env string

const (
	// DEV is the development env tag
	SANDBOX = iota
	// PRODUCTION is the production env tag
	PRODUCTION
)

// Mpesa is mpesa
type Mpesa struct {
	ConsumerKey    string
	ConsumerSecret string
	Env            int
}

// New return a new Mpesa
func New(appKey, appSecret string, env int) (Mpesa, error) {
	return Mpesa{appKey, appSecret, env}, nil
}

//Generate Daraja Access Token
func (m Mpesa) authenticate() (string, error) {
	b := []byte(m.ConsumerKey + ":" + m.ConsumerSecret)
	encoded := base64.StdEncoding.EncodeToString(b)

	url := m.baseURL() + "oauth/v1/generate?grant_type=client_credentials"
	request, err := http.NewRequest(http.MethodGet, url, strings.NewReader(encoded))
	if err != nil {
		return "", err
	}
	request.Header.Add("authorization", "Basic "+encoded)
	request.Header.Add("cache-control", "no-cache")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	var authResponse authResponse
	json.NewDecoder(response.Body).Decode(&authResponse)
	defer response.Body.Close()

	accessToken := authResponse.AccessToken
	log.Println("Received access_token: ", accessToken)
	return accessToken, nil
}

// STKPushSimulation sends an STK push?
func (m Mpesa) STKPushSimulation(stkPush STKPush) (string, error) {
	var stkPushes []STKPush
	stkPushes = append(stkPushes, stkPush)

	body, err := json.Marshal(stkPushes)
	if err != nil {
		return "", nil
	}
	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := m.baseURL() + "mpesa/stkpush/v1/processrequest"
	return m.newStringRequest(url, body, headers)
}

// STKPushTransactionStatus gets a status
func (m Mpesa) STKPushTransactionStatus(stkPush STKPush) (string, error) {
	var stkPushes []STKPush
	stkPushes = append(stkPushes, stkPush)

	body, err := json.Marshal(stkPushes)
	if err != nil {
		return "", nil
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth

	url := m.baseURL() + "mpesa/stkpushquery/v1/query"
	return m.newStringRequest(url, body, headers)
}

// RegisterURL requests
func (m Mpesa) RegisterURL(registerURL RegisterURL) (string, error) {
	var registerURLs []RegisterURL
	registerURLs = append(registerURLs, registerURL)

	body, err := json.Marshal(registerURLs)
	if err != nil {
		return "", err
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["Cache-Control"] = "no-cache"

	url := m.baseURL() + "mpesa/c2b/v1/registerurl"
	return m.newStringRequest(url, body, headers)
}

// C2BSimulation sends a new request
func (m Mpesa) C2BSimulation(c2b C2B) (string, error) {
	var c2bs []C2B
	c2bs = append(c2bs, c2b)
	body, err := json.Marshal(c2bs)
	if err != nil {
		return "", err
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := m.baseURL() + "mpesa/c2b/v1/simulate"
	return m.newStringRequest(url, body, headers)
}

// B2CRequest sends a new request
func (m Mpesa) B2CRequest(b2c B2C) (string, error) {
	var b2cs []B2C
	b2cs = append(b2cs, b2c)

	log.Println(b2cs)

	body, err := json.Marshal(b2cs)
	if err != nil {
		return "", err
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := m.baseURL() + "mpesa/b2c/v1/paymentrequest"
	return m.newStringRequest(url, body, headers)
}

// B2BRequest sends a new request
func (m Mpesa) B2BRequest(b2b B2B) (string, error) {
	var b2bs []B2B
	b2bs = append(b2bs, b2b)

	body, err := json.Marshal(b2bs)
	if err != nil {
		return "", nil
	}
	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := m.baseURL() + "safaricom/b2b/v1/paymentrequest"
	return m.newStringRequest(url, body, headers)
}

// Reversal requests a reversal?
func (m Mpesa) Reversal(reversal Reversal) (string, error) {
	var reversals []Reversal
	reversals = append(reversals, reversal)

	body, err := json.Marshal(reversals)
	if err != nil {
		return "", err
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := m.baseURL() + "safaricom/reversal/v1/request"
	return m.newStringRequest(url, body, headers)
}

// BalanceInquiry sends a balance inquiry
func (m Mpesa) BalanceInquiry(balanceInquiry BalanceInquiry) (string, error) {
	var inquiries []BalanceInquiry
	inquiries = append(inquiries, balanceInquiry)

	auth, err := m.authenticate()
	if err != nil {
		return "", nil
	}

	body, err := json.Marshal(inquiries)
	if err != nil {
		return "", err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"
	headers["postman-token"] = "2aa448be-7d56-a796-065f-b378ede8b136"

	url := m.baseURL() + "safaricom/accountbalance/v1/query"
	return m.newStringRequest(url, body, headers)
}

func (m Mpesa) newStringRequest(url string, body []byte, headers map[string]string) (string, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	log.Println("Response received")
	return string(respBody), nil
}

func (m Mpesa) baseURL() string {
	if m.Env == PRODUCTION {
		return "https://api.safaricom.co.ke/"
	}
	return "https://sandbox.safaricom.co.ke/"
}
