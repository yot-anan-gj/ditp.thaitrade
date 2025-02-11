package common_bindings

type TransportServiceDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	TransportID            string `json:"transport_id"`
	ReferenceNo            string `json:"reference_no"`
	TransportStatus        string `json:"transport_status"`
	SellerCode             string `json:"seller_code"`
	SellerAccountID        string `json:"seller_account_id"`
	SellerNameEN           string `json:"seller_name_en"`
	SellerAddress          string `json:"seller_address"`
	SellerEmail            string `json:"seller_email"`
	SellerCountry          string `json:"seller_country"`
	CompanyID              string `json:"company_id"`
	SellerProvince         string `json:"seller_province"`
	SellerPostalCode       string `json:"seller_postal_code"`
	SellerMobileNo         string `json:"seller_mobile_no"`
	SellerTelNo            string `json:"seller_tel_no"`
	SellerContactFirstName string `json:"seller_contact_first_name"`
	SellerContactLastName  string `json:"seller_contact_last_name"`
	SellerContactPhone     string `json:"seller_contact_phone"`

	BuyerCompanyName   string  `json:"buyer_company_name"`
	BuyerContactPerson string  `json:"buyer_contact_person"`
	BuyerAddress       string  `json:"buyer_address"`
	BuyerEmail         string  `json:"buyer_email"`
	ProductName        string  `json:"product_name"`
	Qty                int     `json:"qty"`
	Amount             float64 `json:"amount"`
	Reason             string  `json:"reason"`
	Remark             string  `json:"remark"`
	CreateTime         int64   `json:"create_time"`
	CreateBy           string  `json:"create_by"`
	CreateCode         string  `json:"create_code"`
	UpdateTime         int64   `json:"update_time"`
	UpdateBy           string  `json:"update_by"`
	UpdateCode         string  `json:"update_code"`

	TransportCode  []string                  `json:"transport_service"`
	TransportItems []TransportServiceItemDTO `json:"transport_items"`
}

type TransportServiceItemDTO struct {
	Item                 string                       `json:"item"`
	RefCode              string                       `json:"ref_code"`
	TransportItemsStatus string                       `json:"transport_items_status"`
	Documents            []TransportServiceItemDocDTO `json:"documents"`
}

type TransportServiceItemDocDTO struct {
	DocumentsStatus   string `json:"documents_status"`
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
