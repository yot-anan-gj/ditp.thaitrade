package common_bindings

type SellerProfileRegisterDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	SellerCode string `json:"seller_code"`
	AccountID  string `json:"account_id"`

	CreateTime  int64  `json:"create_time"`
	CreateBy    string `json:"create_by"`
	CreateCode  string `json:"create_code"`
	UpdateTime  int64  `json:"update_time"`
	UpdateBy    string `json:"update_by"`
	UpdateCode  string `json:"update_code"`
	SubmitTime  int64  `json:"submit_time"`
	SubmitBy    string `json:"submit_by"`
	SubmitCode  string `json:"submit_code"`
	ApproveTime int64  `json:"approve_time"`
	ApproveBy   string `json:"approve_by"`
	ApproveCode string `json:"approve_code"`

	Status string `json:"status"`

	SellerNameEN     string   `json:"seller_name_en"`
	BusinessID       string   `json:"business_id"`
	DitpMemberEL     bool     `json:"ditp_member_el"`
	BusinessType     []string `json:"business_type"`
	CompanyType      string   `json:"company_type"` // company_type_code
	ExportMarket     []string `json:"export_market"`
	SellerDesc       string   `json:"seller_desc"`
	StoreUrl         string   `json:"store_url"`
	Website          string   `json:"website"`
	Email            string   `json:"email"`
	ReturnPolicy     string   `json:"return_policy"`
	GuaranteesPolicy string   `json:"guarantees_policy"`

	CategoryInterest []string `json:"category_interest"`

	SellerIdMagento   string   `json:"seller_id_magento"`
	SellerIdWCM       string   `json:"seller_id_wcm"`
	SellerIdGen       string   `json:"seller_id_gen"`
	SellerMemberOther []string `json:"seller_member_other"`

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

	Address      string `json:"address"`
	Country      string `json:"country"`
	Province     string `json:"province"`
	PostalCode   string `json:"postal_code"`
	Mobile       string `json:"mobile"`
	CompanyPhone string `json:"company_phone"`
	Fax          string `json:"fax"`
	MapLinkUrl   string `json:"map_link_url"`

	AddressImageName         string `json:"address_image_name"`
	AddressImageType         string `json:"address_image_type"`
	AddressImageRequestID    string `json:"address_image_request_id"`
	AddressImagePublicURL    string `json:"address_image_public_url"`
	AddressImageCDNURL       string `json:"address_image_cdn_url"`
	AddressImageContextName  string `json:"address_image_context_name"`
	AddressImageFolder       string `json:"address_image_folder"`
	AddressImageFileName     string `json:"address_image_file_name"`
	AddressImageFileLocation string `json:"address_image_file_location"`
	AddressImageFileSize     int64  `json:"address_image_file_size"`

	DocImportant     []SellerProfileRegisterDocImportantDTO     `json:"doc_important"`
	DocCertification []SellerProfileRegisterDocCertificationDTO `json:"doc_certification"`
	DocAward         []SellerProfileRegisterDocAwardDTO         `json:"doc_award"`

	TotalDocSize int64 `json:"total_doc_size"`

	QuestionnaireIsUsedExport              string   `json:"questionnaire_is_used_export"`              // "T" = YES, "F", NO, "" = null
	QuestionnaireUsedExportCountries       []string `json:"questionnaire_used_export_countries"`       // [country_code]
	QuestionnaireInterestedExportCountries []string `json:"questionnaire_interested_export_countries"` // [country_code]
	QuestionnaireIsReadyToExport           string   `json:"questionnaire_is_ready_to_export"`          // "T" = YES, "F", NO, "" = null
	QuestionnaireNotReadyToExportRemark    string   `json:"questionnaire_not_ready_to_export_remark"`
	QuestionnaireOtherOffer                string   `json:"questionnaire_other_offer"`

	ApplyToSook bool `json:"apply_to_sook"`

	SookPaymentPaypalId                string `json:"sook_payment_paypal_id"`
	SookPaymentPaypalFirstName         string `json:"sook_payment_paypal_first_name"`
	SookPaymentPaypalLastName          string `json:"sook_payment_paypal_last_name"`
	SookPaymentPaypalMerchantAccountId string `json:"sook_payment_paypal_merchant_account_id"`
	SookPaymentPaypalStatus            string `json:"sook_payment_paypal_status"`
	SookPaymentBBLMerchantAccountId    string `json:"sook_payment_bbl_merchant_account_id"`
	SookPaymentBBLStatus               string `json:"sook_payment_bbl_status"`

	SookShipmentThaipostStatus            string `json:"sook_shipment_thaipost_status"`
	SookShipmentDhlMerchantAccountId      string `json:"sook_shipment_dhl_merchant_account_id"`
	SookShipmentDhlStatus                 string `json:"sook_shipment_dhl_status"`
	SookShipmentFastshipMerchantAccountId string `json:"sook_shipment_fastship_merchant_account_id"`
	SookShipmentFastshipStatus            string `json:"sook_shipment_fastship_status"`

	CompanyProfileAddress []SellerProfileWarehourseAddressDTO `json:"company_profile_address"`

	TotalAnnualRevenue int64  `json:"total_annual_revenue"`
	TotalEmployees     int64  `json:"total_employees"`
	YearEstablished    int64  `json:"year_established"`
	VideoUrl1          string `json:"video_url1"`
	VideoUrl2          string `json:"video_url2"`
	VideoUrl3          string `json:"video_url3"`
}

type SellerProfileRegisterDocImportantDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	DocType  string                              `json:"doc_type"`
	DocImage []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocCertificationDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	DocType      string                              `json:"doc_type"`
	DocNameOther string                              `json:"doc_name_other"`
	DocImage     []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocAwardDTO struct {
	Version      int64  `json:"version"`
	VersionOld   int64  `json:"version_old"`
	RefCode      string `json:"ref_code"`
	StatusAction string `json:"status_action"`

	Year        string                              `json:"year"`
	AwardDetail string                              `json:"award_detail"`
	DocImage    []SellerProfileRegisterDocUploadDTO `json:"doc_image"`
}

type SellerProfileRegisterDocUploadDTO struct {
	StatusAction string `json:"status_action"`

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
}

type SellerProfileWarehourseAddressDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId  string `json:"account_id"`
	SellerCode string `json:"seller_code"`

	RefCode       string `json:"ref_code"`
	Address       string `json:"address"`
	Tel           string `json:"tel"`
	Country       string `json:"country"`
	ProvinceInput string `json:"province_input"`
	StateInput    string `json:"state_input"`
	PostalCode    string `json:"postal_code"`

	IsDefault bool `json:"is_default"`

	Company string `json:"company"`
}

type SellerSiteDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	SellerCode string `json:"seller_code"`
	AccountID  string `json:"account_id"`

	VideoUrl1 string `json:"video_url1"`
	VideoUrl2 string `json:"video_url2"`
	VideoUrl3 string `json:"video_url3"`

	MainBanners []SellerSiteDocUploadDTO `json:"main_banners"`
}

type SellerSiteDocUploadDTO struct {
	Status            string `json:"status"`
	Seq               int    `json:"seq"`
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
	Link              string `json:"link"`
}

type SubAccountDTO struct {
	RefCode        string `json:"ref_code"`
	SellerCode     string `json:"seller_code"`
	AccountID      string `json:"account_id"`
	BuyerAccountID string `json:"buyer_account_id"`
	UserType       string `json:"user_type"`
	Status         string `json:"status"`
}

type SookPaymentDTO struct {
	SellerCode                         string `json:"seller_code"`
	AccountID                          string `json:"account_id"`
	ApplyToSook                        bool   `json:"apply_to_sook"`
	PaymentMethodTypeCode              string `json:"payment_method_type"`
	SookPaymentPaypalId                string `json:"sook_payment_paypal_id"`
	SookPaymentPaypalFirstName         string `json:"sook_payment_paypal_first_name"`
	SookPaymentPaypalLastName          string `json:"sook_payment_paypal_last_name"`
	SookPaymentPaypalMerchantAccountId string `json:"sook_payment_paypal_merchant_account_id"`
	SookPaymentPaypalStatus            string `json:"sook_payment_paypal_status"`

	SookPaymentBblMerchantAccountId string `json:"sook_payment_bbl_merchant_account_id"`
	SookPaymentBblStatus            string `json:"sook_payment_bbl_status"`
}

type SookShipmentDTO struct {
	SellerCode                 string `json:"seller_code"`
	AccountID                  string `json:"account_id"`
	ShipmentMethodType         string `json:"shipment_method_type"`
	SookShipmentThaipostStatus string `json:"sook_shipment_thaipost_status"`

	SookShipmentDhlMerchantAccountId string `json:"sook_shipment_dhl_merchant_account_id"`
	SookShipmentDhlStatus            string `json:"sook_shipment_dhl_status"`

	SookShipmentFastshipMerchantAccountId string `json:"sook_shipment_fastship_merchant_account_id"`
	SookShipmentFastshipStatus            string `json:"sook_shipment_fastship_status"`
}
