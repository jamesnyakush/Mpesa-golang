package mpesa

// Mpesa is mpesa
type Mpesa struct {
	AppKey    string
	AppSecret string
}

// NewMpesa return a new Mpesa
func NewMpesa(appKey, appSecret string) (*Mpesa, error) {
	return &Mpesa{appKey, appSecret}, nil
}
