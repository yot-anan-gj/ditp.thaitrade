package common_bindings

type BuyerRegisterDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	BusinessMatchingID string `json:"business_matching_id"`

	RefCode          string                   `json:"ref_code"`
	Status           string                   `json:"status"`
	BuyerAccountID   string                   `json:"buyer_account_id"`
	CompanyName      string                   `json:"company_name"`
	WebSite          string                   `json:"website"`
	AttendeeName     string                   `json:"attendee_name"`
	AttendeePosition string                   `json:"attendee_position"`
	BusinessType     []string                 `json:"business_type"`
	CategoryCode     string                   `json:"category_code"`
	ProductDetail    string                   `json:"product_detail"`
	HSCode           string                   `json:"hs_code"`
	Remark           string                   `json:"remark"`
	Documents        []BusinessMatchingDocDTO `json:"documents"`

	SelectSellers []BuyerRegisterSelectSeller `json:"select_sellers"`
}

type BuyerRegisterSelectSeller struct {
	SellerRegisterRefCode string `json:"seller_register_ref_code"`
}

type BusinessMatchingDocDTO struct {
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

type VerifyBuyerRegisterDTO struct {
	VersionEvent    int64 `json:"version_event"`
	VersionEventOld int64 `json:"version_event_old"`

	BusinessMatchingID string `json:"business_matching_id"`

	Action string `json:"action_from_web"`

	Buyers []VerifyBuyerRegisterDetailDTO `json:"buyers"`

	AccountID string `json:"account_id"`
}

type VerifyBuyerRegisterDetailDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	RefCode        string `json:"ref_code"`
	Selected       bool   `json:"selected"`
	BuyerAccountID string `json:"buyer_account_id"`
}

type SellerRegisterDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	BusinessMatchingID string `json:"business_matching_id"`

	RefCode          string                   `json:"ref_code"`
	Status           string                   `json:"status"`
	SellerCode       string                   `json:"seller_code"`
	SellerAccountID  string                   `json:"seller_account_id"`
	Email            string                   `json:"email"`
	AttendeeName     string                   `json:"attendee_name"`
	AttendeePosition string                   `json:"attendee_position"`
	CategoryCode     string                   `json:"category_code"`
	ProductDetail    string                   `json:"product_detail"`
	Remark           string                   `json:"remark"`
	Documents        []BusinessMatchingDocDTO `json:"documents"`

	SelectBuyers []SellerRegisterSelectBuyer `json:"select_buyers"`

	Mobile  string `json:"mobile"`
	Website string `json:"website"`

	SellerProducts []SellerRegisterSellerProduct `json:"seller_products"`
}

type SellerRegisterSellerProduct struct {
	ProductLink string `json:"product_link"`
	InternalSku string `json:"internal_sku"`
}

type SellerRegisterSelectBuyer struct {
	BuyerRegisterRefCode string `json:"buyer_register_ref_code"`
}

type SurveyFormDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	BusinessMatchingID string `json:"business_matching_id"`

	AccountID string `json:"buyer_account_id"`
	RefCode   string `json:"ref_code"`

	UserType      string  `json:"user_type"`
	CompanyName   string  `json:"company_name"`
	Email         string  `json:"email"`
	ContactPerson string  `json:"contact_person"`
	MobilePhone   string  `json:"mobile_phone"`
	InputDate     float64 `json:"input_date"`

	SurveyFormCompany []SurveyFormCompany `json:"survey_company"`

	SurveyFormSatisfied []SurveyFormSatisfied `json:"survey_satisfied"`

	Question1 string `json:"question1"`
	Answer1   string `json:"answer1"`
	Question2 string `json:"question2"`
	Answer2   string `json:"answer2"`
	Question3 string `json:"question3"`
	Answer3   string `json:"answer3"`
}

type SurveyFormCompany struct {
	ExporterCompanyName string `json:"exporter_company_name"`
	ProductDetail       string `json:"product_detail"`
	ExpectedOrderValue  string `json:"expected_order_value"`
	ActualOrderValue    string `json:"actual_order_value"`
}

type SurveyFormSatisfied struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type SendMailRegisterDTO struct {
	BusinessMatchingId  string   `json:"business_matching_id"`
	AccountId           string   `json:"account_id"`
	SendMailBuyers      []string `json:"send_mail_buyers"`
	SendEmailOther      string   `json:"send_email_other"`
	SendMailSubject     string   `json:"send_mail_subject"`
	SendMailBuyerDetail string   `json:"send_mail_buyer_detail"`
}
