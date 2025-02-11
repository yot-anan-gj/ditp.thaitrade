package configuration

import (
	"fmt"
	"gitlab.com/ditp.thaitrade/enginex/database/nosql/aws_dynamodb"
	"gitlab.com/ditp.thaitrade/enginex/util/cryptutil"
)

type DynamoDBConfig struct {
	ContextName     string
	Region          string
	TableName       map[string]string
	AccessKeyID     string
	SecretAccessKey string
	BillingMode     string
}



func (dynamoDB DynamoDBConfig) String() string {
	return fmt.Sprintf("ContextName: %s, "+
		"TableName: %s,"+
		"Region: %s,"+
		"AccessKeyID: %s,"+
		"SecretAccessKey: %s,"+
		"BillingMode: %s",
		dynamoDB.ContextName,
		dynamoDB.TableName,
		dynamoDB.Region,
		dynamoDB.AccessKeyID,
		dynamoDB.SecretAccessKey,
		dynamoDB.BillingMode)
}

func (dynamoDB DynamoDBConfig) GetAwsConfiguration(secretKey string) ([]aws_dynamodb.DynamoDBOptions, error) {
	var options = make([]aws_dynamodb.DynamoDBOptions, 0)

	accessKeyId := dynamoDB.AccessKeyID

	//decrypt accessKeyID
	decryptAccessKeyID, err := cryptutil.DecryptString(accessKeyId, secretKey)
	if err != nil {
		return nil, err
	}


	secretAccessKey := dynamoDB.SecretAccessKey

	//decrypt secretAccessKey
	decryptSecretAcessKey, err := cryptutil.DecryptString(secretAccessKey, secretKey)
	if err != nil {
		return nil, err
	}

	options = append(options, aws_dynamodb.DynamoDBCredentialOpt(decryptAccessKeyID, decryptSecretAcessKey))


	region := dynamoDB.Region
	options = append(options, aws_dynamodb.DynamoDBRegionOpt(region))

	tableName := dynamoDB.TableName
	options = append(options, aws_dynamodb.DynamoDBTableNameOpt(tableName))

	return options, nil
}