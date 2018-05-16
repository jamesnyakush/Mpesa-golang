# MPESA Golang API Wrapper  [![CircleCI](https://circleci.com/gh/AndroidStudyOpenSource/mpesa-api-go.svg?style=shield)](https://circleci.com/gh/AndroidStudyOpenSource/mpesa-api-go)

The wrapper provides convenient access to the [Safaricom MPESA Daraja API](https://developer.safaricom.co.ke/apis-explorer) for applications written in server-side Golang. :rocket:

**This is Work in Progress!**

## Installing
You can install the package by running:

```
go get github.com/AndroidStudyOpenSource/mpesa-api-go
```

## Usage
The package needs to be configured with your **appKey** and **appSecret** which can be obtained from Safaricom.

```
const (
	appKey = "YOUR_APP_KEY"		    
	appSecret = "YOUR_APP_SECRET"	   
)
```

The following examples with show you how to make requests to the various api's available.

### MPESAExpress (Formerly STKPush)
This api allows you to do Lipa Na M-Pesa payment using STK Push.

### C2B
This api allows you to register C2B Callback URLs to Safaricom, and also Simulate a C2B Transaction in ```Sandbox```

### B2C
This api allows you to do M-Pesa Transaction from company to client.

### B2B
This api allows you to do M-Pesa Transaction from one company to another.

```golang
package main

import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {

	m, err := mpesa.New(appKey, appSecret, mpesa.PRODUCTION)
	if err != nil {
		panic(err)
	}

	response, err := m.B2BRequest(mpesa.B2B{
		Initiator:              "",
		SecurityCredential:     "",
		CommandID:              "",
		SenderIdentifierType:   "",
		RecieverIdentifierType: "",
		Remarks:                "",
		Amount:                 "",
		PartyA:                 "",
		PartyB:                 "",
		AccountReference:       "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(response)

}

```

### Account Balance
This api allows you to do balance inquiry.

### Transaction Status
This api allows you to check the status of transaction.

### Reversal
This api allows you to do a transaction reversal

### Contributing

We’re glad you’re interested in MPESA Daraja Golang SDK, and we’d love to see where you take it. If you would like to contribute code to this project you can do so through GitHub by Forking the Repository and creating a Pull Request.

When submitting code, please make every effort to follow existing conventions and style in order to keep the code as readable as possible. We look forward to you submitting a Pull Request.

Thanks, and please do take it for a joyride!

### License

```text
MIT License

Copyright (c) 2018 Android Study Open Source

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
