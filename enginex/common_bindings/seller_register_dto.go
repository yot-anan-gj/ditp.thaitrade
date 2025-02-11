package common_bindings

type CompanyProfileSellerRequestDTO struct {
	SellerCode []string `json:"seller_code"`
}

type SellerCompanyProfileRequestDTO struct {
	AccountId string `json:"account_id"`
}

type CompanyProfileSellerResponseDTO struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

type SellerAndAccountIdFromSubAccountRequestDTO struct {
	InsertSeller bool   `json:"InsertSeller"`
	AccountID    string `json:"AccountID"`
	SellerCode   string `json:"SellerCode"`
}

type SellerServiceResponseDTO struct {
	Success     bool   `json:"success"`
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
}
