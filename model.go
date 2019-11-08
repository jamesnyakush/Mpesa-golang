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
	RecieverIdentifierType string
	Amount                 string
	PartyA                 string
	PartyB                 string
	Remarks                string
	AccountReference       string
	QueueTimeOutURL        string
	ResultURL              string
}

// Express is a model
type Express struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            string
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
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

// Pull is a model
type Pull struct {
	ShortCode  string
	StartDate  string
	EndDate    string
	PageNumber string
}

// RegisterURL is a model
type C2BRegisterURL struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}
