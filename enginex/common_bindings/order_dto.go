package common_bindings

type OrderDTO struct {
	AccountID  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	PaymentRequestNo string  `json:"payment_request_no"`
	PaymentChannel   string  `json:"payment_channel"`
	PaidAmount       float64 `json:"paid_amount"`
	PaidDate         float64 `json:"paid_date"`

	OrderNo        string  `json:"order_no"`
	ShippingMethod string  `json:"shipping_method"`
	TotalAmount    float64 `json:"total_amount"`
	TotalShipping  float64 `json:"total_shipping"`
	TotalBalance   float64 `json:"total_balance"`

	RequiredBilling  bool   `json:"required_billing"`
	IdCardBilling    string `json:"id_card_billing"`
	BuyerNameBilling string `json:"buyer_name_billing"`
	BillingAddress   string `json:"billing_address"`
	BillingCountry   string `json:"billing_country"`
	BillingState     string `json:"billing_state"`
	BillingCity      string `json:"billing_city"`
	BillingPostcode  string `json:"billing_postcode"`
	PhoneNoBilling   string `json:"phone_no_billing"`
	MobileNoBilling  string `json:"mobile_no_billing"`

	BuyerNameShipping string `json:"buyer_name_shipping"`
	ShippingAddress   string `json:"shipping_address"`
	ShippingCountry   string `json:"shipping_country"`
	ShippingState     string `json:"shipping_state"`
	ShippingCity      string `json:"shipping_city"`
	ShippingPostcode  string `json:"shipping_postcode"`
	PhoneNoShipping   string `json:"phone_no_shipping"`
	MobileNoShipping  string `json:"mobile_no_shipping"`

	DataTransactions []TransactionOrderDTO `json:"data_transactions"`
}

type TransactionOrderDTO struct {
	InternalSku string  `json:"internal_sku"`
	UnitPrice   float64 `json:"unit_price"`
	OrderQty    float64 `json:"order_qty"`
	PriceTotal  float64 `json:"price_total"`
}

type UpdatePaymentDTO struct {
	AccountID        string  `json:"account_id"`
	PaymentRequestNo string  `json:"payment_request_no"`
	PaidDate         float64 `json:"paid_date"`
	Status           string  `json:"status"`
}

type CompleteOrderDTO struct {
	AccountID   string  `json:"account_id"`
	OrderNo     string  `json:"order_no"`
	ReceiveDate float64 `json:"receive_date"`
}

type DeliveryOrderDTO struct {
	SellerCode   string  `json:"seller_code"`
	OrderNo      string  `json:"order_no"`
	TrackingNo   string  `json:"tracking_no"`
	DeliveryDate float64 `json:"delivery_date"`
}
