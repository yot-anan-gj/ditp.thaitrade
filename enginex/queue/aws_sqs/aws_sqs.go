package aws_sqs

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/yot-anan-gj/ditp.thaitrade/enginex/echo_logrus"
	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var sqsValidators = []SqsValidator{
	sqsValidAccessKeyID, sqsValidSecretAccessKey, sqsValidRegion,
}

type AwsSqs struct {
	accessKeyID            string
	secretAccessKey        string
	credentials            *credentials.Credentials
	region                 string
	queueName              map[string]string
	delaySeconds           string
	messageRetentionPeriod string
	fifoQueue              string
	Timeout                int
}

type AwsSqss map[string]*AwsSqs

var (
	ErrorSqsQueueNameReq              = errors.New("queue name is require")
	ErrorSqsDelaySecondsReq           = errors.New("delay seconds is require")
	ErrorSqsMessageRetentionPeriodReq = errors.New("message retention period is require")
	ErrorSqsFifoQueueReq              = errors.New("fifo queue is require")
	ErrorSqsAccessKeyIDReq            = errors.New("access key id is require")
	ErrorSqsSecretAccessKeyReq        = errors.New("secret access key is require")
	ErrorSqsInvalidRegion             = func(region string) error { return fmt.Errorf("not support region %s", region) }
)

func sqsValidAccessKeyID(sqs *AwsSqs) error {
	if stringutil.IsEmptyString(sqs.accessKeyID) {
		return ErrorSqsAccessKeyIDReq
	}
	return nil
}

func sqsValidSecretAccessKey(sqs *AwsSqs) error {
	if stringutil.IsEmptyString(sqs.secretAccessKey) {
		return ErrorSqsSecretAccessKeyReq
	}
	return nil
}

func sqsValidRegion(sqs *AwsSqs) error {
	if !supportRegion[sqs.region] {
		return ErrorSqsInvalidRegion(sqs.region)
	}
	return nil
}

func SqsTimeoutOpt(timeout int) SqsOptions {
	return func(sqs *AwsSqs) error {
		if timeout <= 0 {
			return errors.New("time out not less than 0")
		}
		sqs.Timeout = timeout
		return nil
	}
}

func SqsRegionOpt(region string) SqsOptions {
	return func(sqs *AwsSqs) error {
		if !supportRegion[region] {
			return ErrorSqsInvalidRegion(region)
		}
		sqs.region = region
		return nil
	}
}

func SqsQueueNameOpt(queueName map[string]string) SqsOptions {
	return func(sqs *AwsSqs) error {
		if len(queueName) == 0 {
			return ErrorSqsQueueNameReq
		}
		sqs.queueName = queueName
		return nil
	}
}

func SqsDelaySecondsOpt(delaySeconds string) SqsOptions {
	return func(sqs *AwsSqs) error {
		if stringutil.IsEmptyString(delaySeconds) {
			return ErrorSqsDelaySecondsReq
		}
		sqs.delaySeconds = delaySeconds
		return nil
	}
}

func SqsMessageRetentionPeriodOpt(messageRetentionPeriod string) SqsOptions {
	return func(sqs *AwsSqs) error {
		if stringutil.IsEmptyString(messageRetentionPeriod) {
			return ErrorSqsMessageRetentionPeriodReq
		}
		sqs.messageRetentionPeriod = messageRetentionPeriod
		return nil
	}
}

func SqsFifoQueueOpt(fifoQueue string) SqsOptions {
	return func(sqs *AwsSqs) error {
		if stringutil.IsEmptyString(fifoQueue) {
			return ErrorSqsFifoQueueReq
		}
		sqs.fifoQueue = fifoQueue
		return nil
	}
}

func SqsCredentialOpt(accessKeyID string, secretAccessKey string) SqsOptions {
	return func(inst *AwsSqs) error {
		if stringutil.IsEmptyString(accessKeyID) {
			return ErrorSqsAccessKeyIDReq
		}
		if stringutil.IsEmptyString(secretAccessKey) {
			return ErrorSqsSecretAccessKeyReq
		}
		inst.accessKeyID = accessKeyID
		inst.secretAccessKey = secretAccessKey
		return nil
	}
}

func New(opts ...SqsOptions) (*AwsSqs, error) {
	sqs := &AwsSqs{
		Timeout: 30,
	}

	for _, setter := range opts {
		if err := setter(sqs); err != nil {
			log.Warn(err)
			return nil, err
		}
	}
	//validate sqs
	for _, validator := range sqsValidators {
		err := validator(sqs)
		if err != nil {
			log.Warn(err)
			return nil, err
		}
	}

	creds := credentials.NewStaticCredentials(
		sqs.accessKeyID,
		sqs.secretAccessKey, "")

	sqs.credentials = creds
	return sqs, nil
}

func (awsSqs *AwsSqs) session() (*session.Session, error) {
	return session.NewSession(
		&aws.Config{
			Credentials: awsSqs.credentials,
			Region:      aws.String(awsSqs.region),
		},
	)
}

func (awsSqs *AwsSqs) client() (*sqs.SQS, error) {
	round := 0
reconnectLoop:
	for {
		sess, err := awsSqs.session()
		if err != nil && round >= awsSqs.Timeout {
			return nil, err
		} else if err != nil && round < awsSqs.Timeout {
			log.Warnf("sqs create session fail %s", err)
			time.Sleep(1 * time.Second)
			round++
			log.Warnf("sqs re-create session %d time", round)
			continue reconnectLoop
		}

		s := sqs.New(sess)
		if s == nil && round >= awsSqs.Timeout {
			return nil, fmt.Errorf("sqs create client fail")
		} else if s == nil && round < awsSqs.Timeout {
			log.Warnf("sqs create client fail")
			time.Sleep(1 * time.Second)
			round++
			log.Warnf("sqs recreate client %d time", round)
			continue reconnectLoop
		}
		return s, nil
	}
}

func (awsSqs *AwsSqs) Connection() (*sqs.SQS, error) {
	return awsSqs.client()
}

func (awsSqs *AwsSqs) GetQueueName(key string) string {
	return awsSqs.queueName[key]
}

func (awsSqs *AwsSqs) GetDelaySeconds() string {
	return awsSqs.delaySeconds
}

func (awsSqs *AwsSqs) GetMessageRetentionPeriod() string {
	return awsSqs.messageRetentionPeriod
}

func (awsSqs *AwsSqs) GetFifoQueue() string {
	return awsSqs.fifoQueue
}
