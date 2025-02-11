package configuration

import (
	"fmt"
	"gitlab.com/ditp.thaitrade/enginex/queue/aws_sqs"
	"gitlab.com/ditp.thaitrade/enginex/util/cryptutil"
)


type SqsConfig struct {
	ContextName            string
	Region                 string
	QueueName              map[string]string
	DelaySeconds           string
	MessageRetentionPeriod string
	FifoQueue              string
	AccessKeyID            string
	SecretAccessKey        string
}


func (sqs SqsConfig) String() string {
	return fmt.Sprintf("ContextName: %s, " +
		"Region: %s, " +
		"QueueName: %s, " +
		"DelaySeconds: %s, " +
		"MessageRetentionPeriod: %s, " +
		"FifoQueue: %s, " +
		"AccessKeyID: %s, " +
		"SecretAccessKey: %s",
		sqs.ContextName,
		sqs.Region,
		sqs.QueueName,
		sqs.DelaySeconds,
		sqs.MessageRetentionPeriod,
		sqs.FifoQueue,
		sqs.AccessKeyID,
		sqs.SecretAccessKey)
}

func (sqs SqsConfig) GetAwsConfiguration(secretKey string) ([]aws_sqs.SqsOptions, error) {
	var options = make([]aws_sqs.SqsOptions, 0)

	accessKeyId := sqs.AccessKeyID

	//decrypt accessKeyID
	decryptAccessKeyID, err := cryptutil.DecryptString(accessKeyId, secretKey)
	if err != nil {
		return nil, err
	}


	secretAccessKey := sqs.SecretAccessKey

	//decrypt secretAccessKey
	decryptSecretAcessKey, err := cryptutil.DecryptString(secretAccessKey, secretKey)
	if err != nil {
		return nil, err
	}


	options = append(options, aws_sqs.SqsCredentialOpt(decryptAccessKeyID, decryptSecretAcessKey))

	queueName := sqs.QueueName
	options = append(options, aws_sqs.SqsQueueNameOpt(queueName))

	region := sqs.Region
	options = append(options, aws_sqs.SqsRegionOpt(region))

	delaySeconds := sqs.DelaySeconds
	options = append(options, aws_sqs.SqsDelaySecondsOpt(delaySeconds))

	fifoQueue := sqs.FifoQueue
	options = append(options, aws_sqs.SqsFifoQueueOpt(fifoQueue))

	messageRetentionPeriod := sqs.MessageRetentionPeriod
	options = append(options, aws_sqs.SqsMessageRetentionPeriodOpt(messageRetentionPeriod))

	return options, nil
}