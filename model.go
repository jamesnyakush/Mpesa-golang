package mpesa

type authResponse struct {
	AccessToken string `json:"access_token"`
}

// C2B is a model
type C2B struct {
	ShortCode     string
	CommandID     string
	Amount        string
	Msisdn        string
	BillRefNumber string
}

// B2C is a model
type B2C struct {
	InitiatorName      string
	SecurityCredential string
	CommandID          string
	Amount             string
	PartyA             string
	PartyB             string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
	Occassion          string
}

// B2B is a model
type B2B struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	ReceiverIdentifierType string
	Amount                 float32
	PartyA                 string
	PartyB                 string
	Remarks                string
	AccountReference       string
	QueueTimeOutURL        string
	ResultURL              string
}

// STKPush is a model
type STKPush struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            string
	PhoneNumber       string
	PartyA            string
	PartyB            string
	CallBackURL       string
	AccountReference  string
	QueueTimeOutURL   string
	TransactionDesc   string
	CheckOutRequestID string
}

// Reversal is a model
type Reversal struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	TransactionID          string
	Amount                 string
	ReceiverParty          string
	ReceiverIdentifierType string
	QueueTimeOutURL        string
	ResultURL              string
	Remarks                string
	Occassion              string
}

// BalanceInquiry is a model
type BalanceInquiry struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	PartyA             string
	IdentifierType     string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
}

// RegisterURL is a model
type RegisterURL struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}
