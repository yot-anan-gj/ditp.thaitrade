package aws_dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBOptions func(*DynamoDB) error
type DynamoDBValidator func(*DynamoDB) error

var (
	USEastOhio            = "us-east-2"
	USEastNVirginia       = "us-east-1"
	USWestNCalifornia     = "us-west-1"
	USWestOregon          = "us-west-2"
	AsiaPacificHongKong   = "ap-east-1"
	AsiaPacificMumbai     = "ap-south-1"
	AsiaPacificOsakaLocal = "ap-northeast-3"
	AsiaPacificSeoul      = "ap-northeast-2"
	AsiaPacificSingapore  = "ap-southeast-1"
	AsiaPacificSydney     = "ap-southeast-2"
	AsiaPacificTokyo      = "ap-northeast-1"
	CanadaCentral         = "ca-central-1"
	ChinaBeijing          = "cn-north-1"
	ChinaNingxia          = "cn-northwest-1"
	EUFrankfurt           = "eu-central-1"
	EUIreland             = "eu-west-1"
	EULondon              = "eu-west-2"
	EUParis               = "eu-west-3"
	EUStockholm           = "eu-north-1"
	SouthAmericaSaoPaulo  = "sa-east-1"
	MiddleEastBahrain     = "me-south-1"
)

var supportRegion = map[string]bool{
	USEastOhio:            true,
	USEastNVirginia:       true,
	USWestNCalifornia:     true,
	USWestOregon:          true,
	AsiaPacificHongKong:   true,
	AsiaPacificMumbai:     true,
	AsiaPacificOsakaLocal: true,
	AsiaPacificSeoul:      true,
	AsiaPacificSingapore:  true,
	AsiaPacificSydney:     true,
	AsiaPacificTokyo:      true,
	CanadaCentral:         true,
	ChinaBeijing:          true,
	ChinaNingxia:          true,
	EUFrankfurt:           true,
	EUIreland:             true,
	EULondon:              true,
	EUParis:               true,
	EUStockholm:           true,
	SouthAmericaSaoPaulo:  true,
	MiddleEastBahrain:     true,
}


//tables attribute type
var (
	StringAttributeType = aws.String("S")
	NumberAttributeType = aws.String("N")
	BinaryAttributeType = aws.String("B")
)

//tables key schema type
var (
	PartitionKeyType = aws.String("HASH")
	SortKeyType = aws.String("RANGE")
)

//tables billing mode
var (
	BillingModeProvisioned = aws.String(dynamodb.BillingModeProvisioned)
	BillingModePayPerReqRequest = aws.String(dynamodb.BillingModePayPerRequest)
)

