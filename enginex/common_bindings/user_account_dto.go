package common_bindings

type UserAccountDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId               string `json:"account_id"`
	NameEN                  string `json:"name_en"`
	SurnameEN               string `json:"surname_en"`
	Tel                     string `json:"tel"`
	Email                   string `json:"email"`
	Country                 string `json:"country"`
	CountryState            string `json:"country_state"`
	CompanyEmpTotal         int64  `json:"company_emp_total"`
	CompanyFile             string `json:"company_file"`
	CompanyFileName         string `json:"company_file_name"`
	CompanyFileShow         string `json:"company_file_show"`
	CompanyFileShowUpload   string `json:"company_file_show_upload"`
	CompanyFileShowCdn      string `json:"company_file_show_cdn"`
	CompanyFileCdnUrlTemp   string `json:"company_file_cdn_url_temp"`
	CompanyFileType         string `json:"company_file_type"`
	CompanyFileRequestID    string `json:"company_file_request_id"`
	CompanyFilePublicUrl    string `json:"company_file_public_url"`
	CompanyFileCdnUrl       string `json:"company_file_cdn_url"`
	CompanyFileContextName  string `json:"company_file_context_name"`
	CompanyFileFolder       string `json:"company_file_folder"`
	CompanyFileFileName     string `json:"company_file_file_name"`
	CompanyFileFileLocation string `json:"company_file_file_location"`
	CompanyFileFileSize     int64  `json:"company_file_file_size"`
	Gender                  string `json:"gender_id"`

	BusinessRole string `json:"business_role"`
	BusinessID   string `json:"business_id"`
	UserType     string `json:"user_type"`

	IsSubscribe         bool `json:"is_subscribe"`
	IsVerifyByThaitrade bool `json:"is_verify_by_thaitrade"`
	IsUserThaitrade     bool `json:"is_user_thaitrade"`

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

	Channel string `json:"channel"`

	GoogleID        string `json:"google_id"`
	GoogleFirstName string `json:"google_first_name"`
	GoogleLastName  string `json:"google_last_name"`
	GoogleFullName  string `json:"google_full_name"`
	GoogleEmail     string `json:"google_email"`
	GoogleAvatarURL string `json:"google_avatar_url"`

	FacebookID        string `json:"facebook_id"`
	FacebookFirstName string `json:"facebook_first_name"`
	FacebookLastName  string `json:"facebook_last_name"`
	FacebookFullName  string `json:"facebook_full_name"`
	FacebookEmail     string `json:"facebook_email"`
	FacebookAvatarURL string `json:"facebook_avatar_url"`

	AppleID string `json:"apple_id"`

	CategoryInterest []string `json:"category_interest"`
	TimeZone         string   `json:"time_zone"`

	IsConfirmed    bool `json:"is_confirmed"`
	UserTransfer   bool `json:"user_transfer"`
	UserTransferID int  `json:"user_transfer_id"`

	CompanyName     string `json:"company_name"`
	VerifyTime      int64  `json:"verify_time"`
	VerifyAccountID string `json:"verify_account_id"`
}

type UserAccountConfirmEmailDTO struct {
	AccountId   string `json:"account_id"`
	Email       string `json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	ConfirmDate int64  `json:"confirm_date"`
}

type UserAccountAddressDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	AccountId string `json:"account_id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	RefCode       string `json:"ref_code"`
	Address       string `json:"address"`
	Tel           string `json:"tel"`
	Country       string `json:"country"`
	ProvinceInput string `json:"province_input"`
	StateInput    string `json:"state_input"`
	PostalCode    string `json:"postal_code"`

	IsBilling  bool `json:"is_billing"`
	IsShipping bool `json:"is_shipping"`
}

type ReviewDTO struct {
	RefCode string `json:"ref_code"`

	SellerCode string `json:"seller_code"`
	AccountID  string `json:"account_id"`

	ReviewScore   string `json:"review_score"`
	ReviewComment string `json:"review_comment"`
	ReviewType    string `json:"review_type"`

	ProductSKU string `json:"product_sku"`
	VariantSKU string `json:"variant_sku"`

	Status string `json:"status"`
}
