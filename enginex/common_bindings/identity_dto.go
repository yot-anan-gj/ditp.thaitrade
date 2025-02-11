package common_bindings

type BuyerToSellerRequestDTO struct {
	AccountID    string `json:"AccountID"`
	UserType     string `json:"UserType"`
	BusinessRole string `json:"BusinessRole"`
	SellerCode   string `json:"SellerCode"`
}

type BuyerToSellerResponseDTO struct {
	Success     bool   `json:"success"`
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
}
