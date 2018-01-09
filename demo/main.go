package main

import mpesa "github.com/AndroidStudyOpenSource/mpesa-api-go"
import "log"

const (
	appKey    = "GvzjNnYgNJtwgwfLBkZh65VPwfuKvs0V" // sandbox --> change to yours
	appSecret = "oOpJICRVlyrGSAkM"                 // sandbox --> change to yours
)

func main() {
	m, err := mpesa.New(appKey, appSecret)
	if err != nil {
		panic(err)
	}

	response, err := m.B2BRequest(mpesa.B2B{})
	if err != nil {
		log.Println(err)
	}

	log.Println(response)
}
