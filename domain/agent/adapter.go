package agent

const (
	PRE_CHECK        = "/v2/precheck"
	TEMPLATE_PAYMENT = "/v2/templatepayment"
)

type AdapterInterface interface {
	TemplatePayment(adapterRequest *AdapterRequest) (*Response, error)
	PreCheck(adapterRequest *AdapterRequest) (*Response, error)
}

type AdapterRequest struct {
	Login      string  `url:"login" json:"login"`
	Password   string  `url:"pass" json:"password"`
	Type       string  `url:"type" json:"type"`
	Msisdn     string  `url:"msisdn,omitempty" json:"msisdn"`
	Amount     float64 `url:"pay_amount,omitempty" json:"amount"`
	Account    string  `url:"account,omitempty" json:"account"`
	ReceiptNum string  `url:"receipt_number,omitempty" json:"receipt_num"`
	PayDate    string  `url:"pay_date,omitempty" json:"pay_date"`
}

type AdapterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Receipt string `json:"receipt"`
	Balance string `json:"balance"`
}
