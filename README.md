# MPESA Golang API Wrapper  [![CircleCI](https://circleci.com/gh/AndroidStudyOpenSource/mpesa-api-go.svg?style=shield)](https://circleci.com/gh/AndroidStudyOpenSource/mpesa-api-go)

The wrapper provides convenient access to the [Safaricom MPESA Daraja API](https://developer.safaricom.co.ke/apis-explorer) for applications written in server-side Golang. 

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

### C2B

### B2C

### B2B

### Account Balance

### Transaction Status

### Reversal

### Contributing

We’re glad you’re interested in MPESA Daraja Golang SDK, and we’d love to see where you take it. If you would like to contribute code to this project you can do so through GitHub by Forking the Repository and creating a Pull Request.

When submitting code, please make every effort to follow existing conventions and style in order to keep the code as readable as possible. We look forward to you submitting a Pull Request.

Thanks, and please do take it for a joyride!
