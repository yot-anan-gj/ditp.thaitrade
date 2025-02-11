package common_bindings

type EquotationDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountID  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	Status string `json:"status"`

	Channel      string `json:"chanel"`
	EquotationID string `json:"equotation_id"`
	ValidUntil   string `json:"valid_until"`
	Date         int64  `json:"quotation_date"`
	DueDate      int64  `json:"due_date"`
	Currency     string `json:"currency"`

	BuyerRequestID string `json:"buying_request_id"`

	ProductInSku        string  `json:"product_in_sku"`
	ProductVariantInSku string  `json:"product_variant_in_sku"`
	ProductQty          float64 `json:"product_qty"`
	ProductPrice        float64 `json:"product_price"`

	CurrencyExchangeRate    float64 `json:"currency_exchange_rate"`
	CurrencyExchangeRateUSD float64 `json:"currency_exchange_rate_usd"`

	SellerFname       string `json:"seller_fname"`
	SellerLname       string `json:"seller_lname"`
	SellerCompanyName string `json:"seller_company_name"`
	SellerPhoneNumber string `json:"seller_phone_number"`
	SellerAddress     string `json:"seller_address"`
	SellerCountry     string `json:"seller_country"`
	SellerProvince    string `json:"seller_province"`
	SellerState       string `json:"seller_state"`
	SellerPostalCode  string `json:"seller_postal_code"`

	BuyerAccountID   string `json:"buyer_account_id"`
	BuyerFname       string `json:"buyer_fname"`
	BuyerLname       string `json:"buyer_lname"`
	BuyerCompanyName string `json:"buyer_company_name"`
	BuyerPhoneNumber string `json:"buyer_phone_number"`
	BuyerAddress     string `json:"buyer_address"`
	BuyerCountry     string `json:"buyer_country"`
	BuyerProvince    string `json:"buyer_province"`
	BuyerState       string `json:"buyer_state"`
	BuyerPostalCode  string `json:"buyer_postal_code"`

	TotalAmount       float64 `json:"total_amount"`
	TermsAndCondition string  `json:"terms_and_condition"`
	Remark            string  `json:"remark"`

	ProductDetail []EquotationProductDetailDTO `json:"product_detail"`
}

type EquotationProductDetailDTO struct {
	RefCode string `json:"ref_code"`

	Items        string  `json:"items"`
	Description  string  `json:"description"`
	PricePerUnit float64 `json:"price_per_unit"`
	Qty          float64 `json:"qty"`
	Unit         string  `json:"unit"`
	TotalPrice   float64 `json:"total_price"`

	Status string `json:"product_detail_status"`
}

type EquotationRequestDTO struct {
	Channel            string                     `json:"Channel"`
	BuyingReqId        string                     `json:"BuyingReqId"`
	SellerCode         string                     `json:"SellerCode"`
	BuyerAccountId     string                     `json:"BuyerAccountId"`
	Lookingfor         string                     `json:"Lookingfor"`
	InternalSKU        string                     `json:"InternalSKU"`
	InternalProductSKU string                     `json:"InternalProductSKU`
	ProductName        string                     `json:"ProductName"`
	ProductVariant     string                     `json:"ProductVariant"`
	ProductImage       string                     `json:"ProductImage"`
	ProductQuantity    int64                      `json:"ProductQuantity"`
	ProductPrice       float64                    `json:"ProductPrice"`
	CurrencyExchange   string                     `json:"CurrencyExchange"`
	EquotationItems    []EquotationItemRequestDTO `json:"EquotationItems"`
}

type EquotationItemRequestDTO struct {
	Item        string  `json:"Item"`
	Description string  `json:"Description"`
	Price       float64 `json:"Price"`
	Quantity    int64   `json:"Quantity"`
	Unit        string  `json:"Unit"`
	Total       float64 `json:"Total"`
}

type EquotationResponseDTO struct {
	Success     bool                 `json:"Success"`
	MessageCode string               `json:"MessageCode"`
	Message     string               `json:"Message"`
	Data        EquotationRequestDTO `json:"Data"`
}
