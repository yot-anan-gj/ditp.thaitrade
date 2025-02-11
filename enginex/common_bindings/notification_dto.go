package common_bindings

type NotificationRequestDTO struct {
	AccountId  []string `json:"account_id"`
	SellerCode []string `json:"seller_code"`
	Code       string   `json:"code"`
	Title      string   `json:"title"`
	Link       string   `json:"link"`
	Icon       string   `json:"icon"`
	Type       string   `json:"type"`
	UserType   string   `json:"user_type"`
}

type NotificationResponseDTO struct {
	Success     bool   `json:"success"`
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
}
