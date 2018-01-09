package main

import (
	"log"
	mpesa "github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = "GvzjNnYgNJtwgwfLBkZh65VPwfuKvs0V" // sandbox --> change to yours
	appSecret = "oOpJICRVlyrGSAkM"                 // sandbox --> change to yours
)

func main() {
	// These examples are taken from the mpesa-java-sdk examples
	// at https://github.com/safaricom/mpesa-java-sdk

	m, err := mpesa.New(appKey, appSecret, mpesa.DEV)
	if err != nil {
		panic(err)
	}

	response, err := m.B2BRequest(mpesa.B2B{})
	if err != nil {
		log.Println(err)
	}
	log.Println(response)

	c2b := mpesa.C2B{
		ShortCode:     "600576",
		CommandID:     "CustomerPayBillOnline",
		Amount:        "2",
		Msisdn:        "254708374149",
		BillRefNumber: "hkjhjkhjkh"}
	response, err = m.C2BSimulation(c2b)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)

	// Uncomment the code block you want to test
	//
	// b2c := mpesa.B2C{
	// 	InitiatorName:      "testapi",
	// 	SecurityCredential: "BVeDP3XWGFG+NCQri04jHp6c0rCajO1JAOccQ7Bsu/Mup3Rh2Gd9IHQEE0SeA1oBXAt/VBAL/cJP+VKU9qRF6voqCa0P1XG8pcv5hTZUcBkbbb8Qqvqn28+s/tBvsLXwsB4QaageFDDZgS6b6gbK1p7+UZ/hRYHL8WclTpYBrQGfhqKZxduh0bPWvK4rt+uqR3hdVlO0RdJSkcOVCVp+FxizPSk3nI6LFq14Jj2G0TwuQ4a13J/KVu5eeFG65gzE1NnIVouHKeBPz9b9xvove156aR16uxh4rBq5U6UAKC/kUhaJ0wOLTvb762CioudL87C6xaPVdTF4qcSD6jM4PA==",
	// 	CommandID:          "BusinessPayment",
	// 	Amount:             "22",
	// 	PartyA:             "600576",
	// 	PartyB:             "254708374149",
	// 	Remarks:            "This",
	// 	QueueTimeOutURL:    "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BConfirmation",
	// 	ResultURL:          "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation",
	// 	Occassion:          "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation"}
	// response, err = m.B2CRequest(b2c)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// b2b := mpesa.B2B{
	// 	Initiator:              "testapi",
	// 	AccountReference:       "his",
	// 	SecurityCredential:     "BVeDP3XWGFG+NCQri04jHp6c0rCajO1JAOccQ7Bsu/Mup3Rh2Gd9IHQEE0SeA1oBXAt/VBAL/cJP+VKU9qRF6voqCa0P1XG8pcv5hTZUcBkbbb8Qqvqn28+s/tBvsLXwsB4QaageFDDZgS6b6gbK1p7+UZ/hRYHL8WclTpYBrQGfhqKZxduh0bPWvK4rt+uqR3hdVlO0RdJSkcOVCVp+FxizPSk3nI6LFq14Jj2G0TwuQ4a13J/KVu5eeFG65gzE1NnIVouHKeBPz9b9xvove156aR16uxh4rBq5U6UAKC/kUhaJ0wOLTvb762CioudL87C6xaPVdTF4qcSD6jM4PA==",
	// 	CommandID:              "BusinessPayBill",
	// 	SenderIdentifierType:   "1",
	// 	ReceiverIdentifierType: "4",
	// 	Amount:                 22,
	// 	PartyA:                 "600576",
	// 	PartyB:                 "600000",
	// 	Remarks:                "This",
	// 	QueueTimeOutURL:        "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BConfirmation",
	// 	ResultURL:              "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation"}
	// response, err = m.B2BRequest(b2b)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// stkPush := mpesa.STKPush{
	// 	BusinessShortCode: "174379",
	// 	Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTcwODI0MTU1MDU1",
	// 	Timestamp:         "20170824155055",
	// 	TransactionType:   "CustomerPayBillOnline",
	// 	Amount:            "1",
	// 	PhoneNumber:       "254724513769",
	// 	PartyA:            "254724513769",
	// 	PartyB:            "174379",
	// 	CallBackURL:       "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation",
	// 	QueueTimeOutURL:   "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation",
	// 	AccountReference:  "sasas",
	// 	TransactionDesc:   "asdasd"}
	// response, err = m.STKPushSimulation(stkPush)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// reversal := mpesa.Reversal{
	// 	Initiator:              "testapi",
	// 	SecurityCredential:     "BVeDP3XWGFG+NCQri04jHp6c0rCajO1JAOccQ7Bsu/Mup3Rh2Gd9IHQEE0SeA1oBXAt/VBAL/cJP+VKU9qRF6voqCa0P1XG8pcv5hTZUcBkbbb8Qqvqn28+s/tBvsLXwsB4QaageFDDZgS6b6gbK1p7+UZ/hRYHL8WclTpYBrQGfhqKZxduh0bPWvK4rt+uqR3hdVlO0RdJSkcOVCVp+FxizPSk3nI6LFq14Jj2G0TwuQ4a13J/KVu5eeFG65gzE1NnIVouHKeBPz9b9xvove156aR16uxh4rBq5U6UAKC/kUhaJ0wOLTvb762CioudL87C6xaPVdTF4qcSD6jM4PA==",
	// 	CommandID:              "TransactionReversal",
	// 	TransactionID:          "2121",
	// 	Amount:                 "2",
	// 	ReceiverParty:          "22",
	// 	ReceiverIdentifierType: "4",
	// 	ResultURL:              "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation",
	// 	QueueTimeOutURL:        "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BConfirmation",
	// 	Remarks:                "Remarks",
	// 	Occassion:              "Ocassions"}
	// response, err = m.Reversal(reversal)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// registerURL := mpesa.RegisterURL{
	// 	ShortCode:       "600576",
	// 	ResponseType:    "Cancelled",
	// 	ConfirmationURL: "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation",
	// 	ValidationURL:   "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BValidation"}
	// response, err = m.RegisterURL(registerURL)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// balanceInquiry := mpesa.BalanceInquiry{
	// 	Initiator:          "testapi",
	// 	CommandID:          "AccountBalance",
	// 	SecurityCredential: "BVeDP3XWGFG+NCQri04jHp6c0rCajO1JAOccQ7Bsu/Mup3Rh2Gd9IHQEE0SeA1oBXAt/VBAL/cJP+VKU9qRF6voqCa0P1XG8pcv5hTZUcBkbbb8Qqvqn28+s/tBvsLXwsB4QaageFDDZgS6b6gbK1p7+UZ/hRYHL8WclTpYBrQGfhqKZxduh0bPWvK4rt+uqR3hdVlO0RdJSkcOVCVp+FxizPSk3nI6LFq14Jj2G0TwuQ4a13J/KVu5eeFG65gzE1NnIVouHKeBPz9b9xvove156aR16uxh4rBq5U6UAKC/kUhaJ0wOLTvb762CioudL87C6xaPVdTF4qcSD6jM4PA==",
	// 	PartyA:             "600576",
	// 	IdentifierType:     "4",
	// 	Remarks:            "These",
	// 	QueueTimeOutURL:    "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BConfirmation",
	// 	ResultURL:          "http://obscure-bayou-52273.herokuapp.com/api/Mpesa/C2BConfirmation"}
	// response, err = m.BalanceInquiry(balanceInquiry)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)

	// stkPushTS := mpesa.STKPush{
	// 	BusinessShortCode: "174379",
	// 	Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTcwODI0MTU1MDU1",
	// 	Timestamp:         "20170824155055",
	// 	CheckOutRequestID: "ws_CO_27102017101215530"}
	// response, err = m.STKPushTransactionStatus(stkPushTS)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(response)
}
