package common_bindings

type BuyerServiceRequestDTO struct {
	AccountID string `json:"AccountID"`
}

type UserAccountGeneralDTO struct {
	Version                 int64    `json:"Version"`
	VersionOld              int64    `json:"VersionOld"`
	AccountID               string   `json:"AccountID"`
	NameEN                  string   `json:"NameEN"`
	SurnameEN               string   `json:"SurnameEN"`
	Tel                     string   `json:"Tel"`
	Email                   string   `json:"Email"`
	Country                 string   `json:"Country"`
	CountryState            string   `json:"country_state"`
	CompanyEmpTotal         int64    `json:"company_emp_total"`
	CompanyFile             string   `json:"company_file"`
	CompanyFileName         string   `json:"company_file_name"`
	CompanyFileShow         string   `json:"company_file_show"`
	CompanyFileShowUpload   string   `json:"company_file_show_upload"`
	CompanyFileShowCdn      string   `json:"company_file_show_cdn"`
	CompanyFileCdnUrlTemp   string   `json:"company_file_cdn_url_temp"`
	CompanyFileType         string   `json:"company_file_type"`
	CompanyFileRequestID    string   `json:"company_file_request_id"`
	CompanyFilePublicUrl    string   `json:"company_file_public_url"`
	CompanyFileCdnUrl       string   `json:"company_file_cdn_url"`
	CompanyFileContextName  string   `json:"company_file_context_name"`
	CompanyFileFolder       string   `json:"company_file_folder"`
	CompanyFileFileName     string   `json:"company_file_file_name"`
	CompanyFileFileLocation string   `json:"company_file_file_location"`
	CompanyFileFileSize     int64    `json:"company_file_file_size"`
	Gender                  string   `json:"Gender"`
	BusinessRole            string   `json:"BusinessRole"`
	BusinessID              string   `json:"BusinessID"`
	UserType                string   `json:"UserType"`
	IsSubscribe             bool     `json:"IsSubscribe"`
	IsVerifyByThaitrade     bool     `json:"IsVerifyByThaitrade"`
	IsUserThaitrade         bool     `json:"IsUserThaitrade"`
	ImageName               string   `json:"ImageName"`
	ImageType               string   `json:"ImageType"`
	ImageRequestID          string   `json:"ImageRequestID"`
	ImagePublicURL          string   `json:"ImagePublicURL"`
	ImageCDNURL             string   `json:"ImageCDNURL"`
	ImageContextName        string   `json:"ImageContextName"`
	ImageFolder             string   `json:"ImageFolder"`
	ImageFileName           string   `json:"ImageFileName"`
	ImageFileLocation       string   `json:"ImageFileLocation"`
	ImageFileSize           int64    `json:"ImageFileSize"`
	Channel                 string   `json:"Channel"`
	GoogleID                string   `json:"GoogleID"`
	GoogleFirstName         string   `json:"GoogleFirstName"`
	GoogleLastName          string   `json:"GoogleLastName"`
	GoogleFullName          string   `json:"GoogleFullName"`
	GoogleEmail             string   `json:"GoogleEmail"`
	GoogleAvatarURL         string   `json:"GoogleAvatarURL"`
	FacebookID              string   `json:"FacebookID"`
	FacebookFirstName       string   `json:"FacebookFirstName"`
	FacebookLastName        string   `json:"FacebookLastName"`
	FacebookFullName        string   `json:"FacebookFullName"`
	FacebookEmail           string   `json:"FacebookEmail"`
	FacebookAvatarURL       string   `json:"FacebookAvatarURL"`
	CategoryInterest        []string `json:"CategoryInterest"`
	TimeZone                string   `json:"TimeZone"`
	IsConfirmed             bool     `json:"IsConfirmed"`
}

type UserAccountGeneralResponseDTO struct {
	Success     bool                  `json:"Success"`
	MessageCode string                `json:"MessageCode"`
	Message     string                `json:"Message"`
	Data        UserAccountGeneralDTO `json:"Data"`
}

type SubAccountRequestDTO struct {
	Version    int64  `json:"Version"`
	VersionOld int64  `json:"VersionOld"`
	AccountID  string `json:"AccountID"`
	NameEN     string `json:"NameEN"`
	SurnameEN  string `json:"SurnameEN"`
	Email      string `json:"Email"`
	SellerCode string `json:"SellerCode"`
	UserType   string `json:"UserType"`
	Status     string `json:"Status"`
	CreateTime int64  `json:"CreateTime"`
}

type SubAccountResponseDTO struct {
	Success     bool                 `json:"success"`
	MessageCode string               `json:"messageCode"`
	Message     string               `json:"message"`
	Data        SubAccountRequestDTO `json:"data"`
}
