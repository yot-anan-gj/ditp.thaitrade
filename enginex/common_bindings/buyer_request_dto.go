package common_bindings

type BuyerRequestDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	BuyerRequestID string `json:"buying_request_id"`
	AccountID      string `json:"account_id"`

	Status string `json:"status"`
	Reason string `json:"reason"`

	LookingFor                string               `json:"looking_for"`
	ValidWithinCode           string               `json:"valid_within_code"`
	CategoryCode              string               `json:"category_code"`
	Spacification             string               `json:"spacification"`
	EstimatedOrderQuantity    float64              `json:"estimated_order_quantity"`
	PriceType                 string               `json:"price_type"`
	Unit                      string               `json:"unit"`
	UnitOther                 string               `json:"unit_other"`
	TargetPrice               float64              `json:"target_price"`
	Currency                  string               `json:"currency"`
	LocationOfDelivery        string               `json:"location_of_delivery"`
	CityState                 string               `json:"city_state"`
	ShippingPaymentConditions string               `json:"shipping_payment_conditions"`
	Documents                 []BuyerRequestDocDTO `json:"documents"`
}

type BuyerRequestDocDTO struct {
	RefCode string `json:"ref_code"`

	DocType      string                     `json:"doc_type"`
	DocNameOther string                     `json:"doc_name_other"`
	DocDetails   []BuyerRequestDocDetailDTO `json:"document_details"`
}

type BuyerRequestDocDetailDTO struct {
	ImageName         string `json:"image_name"`
	ImageType         string `json:"image_type"`
	ImageRequestID    string `json:"image_request_id"`
	ImagePublicURL    string `json:"image_public_url"`
	ImageCDNURL       string `json:"image_cdn_url"`
	ImageContextName  string `json:"image_context_name"`
	ImageFolder       string `json:"image_folder"`
	ImageFileName     string `json:"image_file_name"`
	ImageFileLocation string `json:"image_file_location"`
	ImageFileSize     int64  `json:"image_file_size"`

	Status string `json:"status"`
}

type BuyerRequestSellerDTO struct {
	VersionSeller    int64 `json:"version_seller"`
	VersionOldSeller int64 `json:"version_old_seller"`

	RefCodeSeller   string `json:"ref_code_seller"`
	BuyingRequestId string `json:"buying_request_id"`

	AccountID  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	Version    int64  `json:"version"`
	VersionOld int64  `json:"version_old"`
	RefCode    string `json:"ref_code"`
	Comment    string `json:"comment"`

	Documents []BuyerRequestDocDetailDTO `json:"documents"`
}

type BuyerRequestConversationDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	BuyingRequestId string `json:"buying_request_id"`
	RefCodeSeller   string `json:"ref_code_seller"`

	UserType   string `json:"user_type"`
	AccountID  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	RefCode string `json:"ref_code"`
	Comment string `json:"comment"`

	Documents []BuyerRequestDocDetailDTO `json:"documents"`
}
