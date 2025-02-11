package common_bindings

type ProductDTO struct {
	Version    int64 `json:"version"`
	VersionOld int64 `json:"version_old"`

	Status      string `json:"status"`
	PassApprove bool   `json:"pass_approve"`

	InternalSku string `json:"internal_sku"`
	Enable      bool   `json:"enable"`
	AccountID   string `json:"account_id"`
	SellerCode  string `json:"seller_code"`

	CategoryLv1      string `json:"category_lv1"`
	CategoryLv2      string `json:"category_lv2"`
	CategoryLv3      string `json:"category_lv3"`
	AttributeSetCode string `json:"attribute_set_code"`

	ProductRejectReason []ProductRejectReason `json:"product_reject_reason"`
	ProductName         []ProductName         `json:"product_name"`
	ProductKeyword      []ProductKeyword      `json:"product_keyword"`
	ProductImage        []ProductImage        `json:"product_image"`
	ProductVideo        []ProductVideo        `json:"product_video"`
	ProductIntro        []ProductIntro        `json:"product_intro"`
	ProductDetail       []ProductDetail       `json:"product_detail"`
	ProductOptional     []ProductOptional     `json:"product_optional"`
	Attribute           []Attribute           `json:"attribute"`

	AskForPrice      bool               `json:"ask_for_price"`
	PriceType        string             `json:"price_type"`
	UnitPrice        float64            `json:"unit_price"`
	IsVariant        bool               `json:"is_variant"`
	UnitOfMeasure    string             `json:"unit_of_measure"`
	UnitOther        string             `json:"unit_other"`
	RangePrice       string             `json:"range_price"`
	TierPrice        []TierPrice        `json:"tier_price"`
	VariantAttribute []VariantAttribute `json:"variant_attribute"`
	VariantSKU       []VariantSku       `json:"variant_sku"`

	Width  float64 `json:"width"`
	Length float64 `json:"length"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`

	ProductIDMagento int `json:"product_id_magento"`

	SookApply             bool            `json:"sook_apply"`               // true,
	SookPriceType         string          `json:"sook_price_type"`          // "U|T",
	SookUnitPrice         float64         `json:"sook_unit_price"`          // 300.00,
	SookUnitOfMeasure     string          `json:"sook_unit_of_measure"`     // "",
	SookUnitOther         string          `json:"sook_unit_other"`          // "",
	SookMinOrderMain      int64           `json:"sook_min_order_main"`      // 1,
	SookRangePrice        string          `json:"sook_range_price"`         // "",
	HaveSookSpacialPrice  bool            `json:"have_sook_spacial_price"`  // true,
	SookSpacialPrice      float64         `json:"sook_spacial_price"`       // 300.00,
	SookSpacialPriceStart int64           `json:"sook_spacial_price_start"` // 1596167705000,
	SookSpacialPriceEnd   int64           `json:"sook_spacial_price_end"`   // 1596167705000,
	SookStock             int64           `json:"sook_stock"`               // true,
	SookTierPrice         []SookTierPrice `json:"sook_tier_price"`

	SookWeight                     float64 `json:"sook_weight"`
	SookWidth                      float64 `json:"sook_width"`
	SookLength                     float64 `json:"sook_length"`
	SookHeight                     float64 `json:"sook_height"`
	SookMaxQuantitiesPerOnePackage float64 `json:"sook_max_quantities_per_one_package"`

	SookBrand string `json:"sook_brand"`
}

type ProductRejectReason struct {
	SubjectLang []LangModel `json:"subject_lang"`
	Reason      string      `json:"reason"`
}

type LangModel struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductName struct {
	RefCode string `json:"ref_code"`

	Code string `json:"language_code"`
	Name string `json:"product_name"`

	Status string `json:"product_name_status"`
}

type ProductKeyword struct {
	RefCode string `json:"ref_code"`

	Keyword  string `json:"keyword"`
	CreateBy string `json:"create_by"`

	Status string `json:"product_keyword_status"`
}

type ProductImage struct {
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

	Status string `json:"product_image_status"`
}

type ProductVideo struct {
	VideoName         string `json:"video_name"`
	VideoType         string `json:"video_type"`
	VideoRequestID    string `json:"video_request_id"`
	VideoPublicURL    string `json:"video_public_url"`
	VideoCDNURL       string `json:"video_cdn_url"`
	VideoContextName  string `json:"video_context_name"`
	VideoFolder       string `json:"video_folder"`
	VideoFileName     string `json:"video_file_name"`
	VideoFileLocation string `json:"video_file_location"`
	VideoFileSize     int64  `json:"video_file_size"`

	Status string `json:"product_video_status"`
}

type ProductIntro struct {
	RefCode string `json:"ref_code"`

	Code    string `json:"language_code"`
	Content string `json:"product_intro"`

	Status string `json:"product_intro_status"`
}

type ProductDetail struct {
	RefCode string `json:"ref_code"`

	Code    string `json:"language_code"`
	Content string `json:"product_detail"`

	Status string `json:"product_detail_status"`
}

type ProductOptional struct {
	RefCode string `json:"ref_code"`

	Code    string `json:"language_code"`
	Content string `json:"product_optional"`

	Status string `json:"product_optional_status"`
}

type Attribute struct {
	RefCode string `json:"ref_code"`

	AttrCode  string `json:"attr_code"`
	ValueType string `json:"value_type"`
	AnsText   string `json:"ans_text"`
	AnsList   string `json:"ans_list"`

	Status string `json:"attribute_status"`
}

type TierPrice struct {
	RefCode string `json:"ref_code"`

	Min          int     `json:"min_order_detail"`
	Max          int     `json:"max_order_detail"`
	RegularPrice float64 `json:"regular_price"`

	Status string `json:"tier_price_status"`
}

type VariantAttribute struct {
	RefCode string `json:"ref_code"`

	AttrCode       string   `json:"attr_code"`
	AttrDetailCode []string `json:"attr_detail_code"`

	Status string `json:"attribute_status"`
}

type VariantSku struct {
	RefCode string `json:"ref_code"`

	VariantAttr1       string `json:"variant_attr1"`
	VariantAttr2       string `json:"variant_attr2"`
	InternalProductSKU string `json:"internal_product_sku"`

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

	UnitPrice  float64     `json:"unit_price"`
	RangePrice string      `json:"range_price"`
	TierPrice  []TierPrice `json:"tier_price"`

	Status string `json:"variant_sku_status"`
}

type SookTierPrice struct {
	RefCode           string  `json:"ref_code"`            //:"",
	MinOrderDetail    int64   `json:"min_order_detail"`    //: 1,
	MaxOrderDetail    int64   `json:"max_order_detail"`    //: 4,
	RegularPrice      float64 `json:"regular_price"`       //: 300.00,
	SpacialPrice      float64 `json:"spacial_price"`       //: 300.00,
	SpacialPriceStart int64   `json:"spacial_price_start"` //: 1596167705000,
	SpacialPriceEnd   int64   `json:"spacial_price_end"`   //: 1596167705000
	TierPriceStatus   string  `json:"tier_price_status"`   //: C|U|D
}
