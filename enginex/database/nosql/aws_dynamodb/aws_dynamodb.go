package aws_dynamodb

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	log "github.com/yot-anan-gj/ditp.thaitrade/enginex/echo_logrus"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var dynamoDBValidators = []DynamoDBValidator{
	dynamoDBValidAccessKeyID, dynamoDBValidSecretAccessKey, dynamoDBValidRegion,
}

type DynamoDB struct {
	accessKeyID     string
	secretAccessKey string
	credentials     *credentials.Credentials
	region          string
	tableName       map[string]string
	Timeout         int
	endPoint        string
}

type DynamoDBs map[string]*DynamoDB

var (
	ErrorDynamoDBTableNameReq       = errors.New("table name is require")
	ErrorDynamoDBAccessKeyIDReq     = errors.New("access key id is require")
	ErrorDynamoDBSecretAccessKeyReq = errors.New("secret access key is require")
	ErrorDynamoDBEndPointReq        = errors.New("endpoint is require")
	ErrorDynamoDBInvalidRegion      = func(region string) error { return fmt.Errorf("not support region %s", region) }
)

func dynamoDBValidAccessKeyID(inst *DynamoDB) error {
	if stringutil.IsEmptyString(inst.accessKeyID) {
		return ErrorDynamoDBAccessKeyIDReq
	}
	return nil
}

func dynamoDBValidSecretAccessKey(inst *DynamoDB) error {
	if stringutil.IsEmptyString(inst.secretAccessKey) {
		return ErrorDynamoDBSecretAccessKeyReq
	}
	return nil
}

func dynamoDBValidRegion(inst *DynamoDB) error {
	if !supportRegion[inst.region] {
		return ErrorDynamoDBInvalidRegion(inst.region)
	}
	return nil
}

func DynamoDBCredentialOpt(accessKeyID string, secretAccessKey string) DynamoDBOptions {
	return func(inst *DynamoDB) error {
		if stringutil.IsEmptyString(accessKeyID) {
			return ErrorDynamoDBAccessKeyIDReq
		}
		if stringutil.IsEmptyString(secretAccessKey) {
			return ErrorDynamoDBSecretAccessKeyReq
		}
		inst.accessKeyID = accessKeyID
		inst.secretAccessKey = secretAccessKey
		return nil
	}

}

func DynamoDBTableNameOpt(tableName map[string]string) DynamoDBOptions {
	return func(inst *DynamoDB) error {
		if len(tableName) == 0 {
			return ErrorDynamoDBTableNameReq
		}
		inst.tableName = tableName
		return nil
	}
}

func DynamoDBTimeoutOpt(timeout int) DynamoDBOptions {
	return func(inst *DynamoDB) error {
		if timeout <= 0 {
			return errors.New("time out not less than 0")
		}
		inst.Timeout = timeout
		return nil
	}
}

func DynamoDBRegionOpt(region string) DynamoDBOptions {
	return func(inst *DynamoDB) error {
		if !supportRegion[region] {
			return ErrorDynamoDBInvalidRegion(region)
		}
		inst.region = region
		return nil
	}
}

func DynamoDBEndPointOpt(endPoint string) DynamoDBOptions {
	return func(inst *DynamoDB) error {
		if len(endPoint) == 0 {
			return ErrorDynamoDBEndPointReq
		}
		inst.endPoint = endPoint
		return nil
	}
}

func New(opts ...DynamoDBOptions) (*DynamoDB, error) {
	inst := &DynamoDB{
		Timeout: 30,
	}

	for _, setter := range opts {
		if err := setter(inst); err != nil {
			log.Warn(err)
			return nil, err
		}
	}
	//validate dynamoDB
	for _, validator := range dynamoDBValidators {
		err := validator(inst)
		if err != nil {
			log.Warn(err)
			return nil, err
		}
	}

	creds := credentials.NewStaticCredentials(
		inst.accessKeyID,
		inst.secretAccessKey, "")

	inst.credentials = creds
	return inst, nil
}

func (inst *DynamoDB) session() (*session.Session, error) {
	return session.NewSession(
		&aws.Config{
			Credentials: inst.credentials,
			Region:      aws.String(inst.region),
			Endpoint: func() *string {
				if len(inst.endPoint) == 0 {
					return nil
				}
				return aws.String(inst.endPoint)
			}(),
		},
	)
}

func (inst *DynamoDB) client() (*dynamodb.DynamoDB, error) {
	round := 0
reconnectLoop:
	for {
		sess, err := inst.session()
		if err != nil && round >= inst.Timeout {
			return nil, err
		} else if err != nil && round < inst.Timeout {
			log.Logger().Warnf("dynamoDB create session fail %s", err)
			time.Sleep(1 * time.Second)
			round++
			log.Logger().Warnf("dynamoDB re-create session %d time", round)
			continue reconnectLoop
		}

		dynamoDB := dynamodb.New(sess)
		if dynamoDB == nil && round >= inst.Timeout {
			return nil, fmt.Errorf("dynamoDB create client fail")
		} else if dynamoDB == nil && round < inst.Timeout {
			log.Logger().Warnf("dynamoDB create client fail")
			time.Sleep(1 * time.Second)
			round++
			log.Logger().Warnf("dynamoDB recreate client %d time", round)
			continue reconnectLoop
		}
		return dynamoDB, nil

	}
}

func (inst *DynamoDB) Connection() (*dynamodb.DynamoDB, error) {
	return inst.client()
}

func (inst *DynamoDB) DescribeTable(tableName string) (*dynamodb.DescribeTableOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	req := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	table, err := client.DescribeTable(req)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return table, nil

}

func (inst *DynamoDB) PingTable(tableName string) (bool, error) {
	client, err := inst.client()
	if err != nil {
		log.Warn(err)
		return false, err
	}

	req := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	_, err = client.DescribeTable(req)
	if err != nil {
		log.Warn(err)
		return false, err
	}
	return true, nil

}

func (inst *DynamoDB) CreateTable(table *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	if table == nil {
		log.Warn("table for create Table in dynamoDB is require")
		return nil, errors.New("table for create Table in dynamoDB is require")
	}

	client, err := inst.client()
	if err != nil {
		return nil, err
	}

	tableOutput, err := client.CreateTable(table)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return tableOutput, nil

}

func (inst *DynamoDB) DeleteTable(name string) (*dynamodb.DeleteTableOutput, error) {

	if stringutil.IsEmptyString(name) {
		log.Warn("table name for delete in dynamoDB is require")
		return nil, errors.New("table name for delete in dynamoDB is require")
	}
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(name),
	}
	client, err := inst.client()
	if err != nil {
		return nil, err
	}
	output, err := client.DeleteTable(input)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return output, nil
}

func (inst *DynamoDB) ListTables() ([]string, error) {
	client, err := inst.client()
	if err != nil {
		return nil, err
	}
	input := &dynamodb.ListTablesInput{}
	tableNames := make([]string, 0)

	for {
		result, err := client.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					log.Logger().Errorf("dynamoDB list table fail %s, %s",
						dynamodb.ErrCodeInternalServerError,
						aerr.Error())
				default:
					log.Logger().Errorf("dynamoDB list table fail %s", aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				log.Logger().Errorf("dynamoDB list table fail %s", err.Error())
			}
		}

		for _, n := range result.TableNames {
			tableNames = append(tableNames, *n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}

	}
	return tableNames, nil
}

func (inst *DynamoDB) ScanItem(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.Scan(input)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst getitem %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.GetItem(input)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst getitem %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) QueryItem(params *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.Query(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst query item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) PutItem(params *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.PutItem(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst put item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) UpdateItem(params *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.UpdateItem(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst update item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) DeleteItem(params *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.DeleteItem(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst delete item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) TransactWriteItems(params *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.TransactWriteItems(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB inst transaction item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) TransactGetItems(params *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	client, err := inst.client()
	if err != nil {
		log.Logger().Errorf("dynamoDB inst table fail %s", err.Error())
		return nil, err
	}

	result, err := client.TransactGetItems(params)
	if err != nil {
		log.Logger().Errorf("dynamoDB get transaction item %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (inst *DynamoDB) GetTableName(key string) string {
	return inst.tableName[key]
}
