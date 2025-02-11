package common_bindings

type UserProfile struct {
	AccountID                  string
	NameEN                     string
	SurnameEN                  string
	Tel                        string
	Email                      string
	Country                    string
	CountryState               string
	CompanyEmpTotal            int64
	CompanyFile                string
	CompanyFileName            string
	CompanyFileShow            string
	CompanyFileShowUpload      string
	CompanyFileShowCdn         string
	CompanyFileCdnUrlTemp      string
	CompanyFileType            string
	CompanyFileRequestID       string
	CompanyFilePublicUrl       string
	CompanyFileCdnUrl          string
	CompanyFileContextName     string
	CompanyFileFolder          string
	CompanyFileFileName        string
	CompanyFileFileLocation    string
	CompanyFileFileSize        int64
	Gender                     string
	BusinessRole               string
	UserType                   string
	Password                   string
	BusinessID                 string
	IsSubscribe                bool
	IsVerifyByThaitrade        bool
	IsUserThaitrade            bool
	IsConfirmed                bool
	HasFavoriteProduct         bool
	Channel                    string
	UserTransfer               bool
	UserTransferChangePassword bool
}

type CompanyProfile struct {
	SellerCode       string
	AccountID        string
	Status           string
	SellerNameEN     string
	BusinessId       string
	DitpMemberEl     bool
	BusinessType     []string
	ExportMarket     []string
	SellerDesc       string
	StoreUrl         string
	Website          string
	Email            string
	CategoryInterest []string
	Country          string
	Mobile           string
	CompanyPhone     string
}
