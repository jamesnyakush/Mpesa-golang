package mpesa

type C2B struct {
	ShortCode     string
	CommandID     string
	Amount        string
	Msisdn        string
	BillRefNumber string
}

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

type B2B struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	ReceiverIdentifierType string
	Amount                 string
	PartyA                 string
	PartyB                 string
	Remarks                string
	AccountReference       string
	QueueTimeOutURL        string
	ResultURL              string
}

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
}
