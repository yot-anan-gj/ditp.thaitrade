package configuration

import (
	"errors"
	"fmt"

	"github.com/yot-anan-gj/ditp.thaitrade/enginex/util/stringutil"
)

var (
	ErrSMTPServerConfigNameReq = errors.New("SMTP server name is require")
	ErrSMTPServerConfigEmpty   = func(name string) error { return fmt.Errorf("SMTP [%s] server config is empty", name) }
	ErrSMTPPPortEmpty          = func(name string) error { return fmt.Errorf("SMTP [%s] port  config is empty", name) }
	ErrSMTPUserEmpty           = func(name string) error { return fmt.Errorf("SMTP [%s] user config is empty", name) }
	ErrSMTPPasswordEmpty       = func(name string) error { return fmt.Errorf("SMTP [%s] password config is empty", name) }
	ErrSMTPSenderIdentityEmpty = func(name string) error { return fmt.Errorf("SMTP [%s] sender identity is empty", name) }
	ErrSMTPSenderEmailEmpty    = func(name string) error { return fmt.Errorf("SMTP [%s] sender email is empty", name) }
	ErrSMTPServerConfigNameDup = func(name string) error { return fmt.Errorf("SMTP name [%s] is duplicate", name) }

	ErrAWSSESConfigNameReq              = errors.New("AWS SES name is require")
	ErrAWSSESConfigRegionEmpty          = func(name string) error { return fmt.Errorf("AWS SES [%s] region empty", name) }
	ErrAWSSESConfigAccessKeyIDEmpty     = func(name string) error { return fmt.Errorf("AWS SES [%s] access key id empty", name) }
	ErrAWSSESConfigSecretAccessKeyEmpty = func(name string) error { return fmt.Errorf("AWS SES [%s] secret access key empty", name) }
	ErrAWSESConfigNameDup               = func(name string) error { return fmt.Errorf("AWS SES name [%s] is duplicate", name) }
)

func validEmailServersConfig(config *Configuration) error {
	if config == nil {
		return ErrorInvalidConfig
	}

	nameCount := make(map[string]int)

	for _, smtpCfg := range config.EmailServers.SMTPs {
		if nameCount[smtpCfg.Name] > 1 {
			return ErrSMTPServerConfigNameDup(smtpCfg.Name)
		}
		nameCount[smtpCfg.Name]++
		if stringutil.IsEmptyString(smtpCfg.Name) {
			return ErrSMTPServerConfigNameReq
		}

		if stringutil.IsEmptyString(smtpCfg.Server) {
			return ErrSMTPServerConfigEmpty(smtpCfg.Name)
		}

		if smtpCfg.Port <= 0 {
			return ErrSMTPPPortEmpty(smtpCfg.Name)
		}

		//if stringutil.IsEmptyString(smtpCfg.User) {
		//	return ErrSMTPUserEmpty(smtpCfg.Name)
		//}
		//
		//if stringutil.IsEmptyString(smtpCfg.Password) {
		//	return ErrSMTPPasswordEmpty(smtpCfg.Name)
		//}

		if stringutil.IsEmptyString(smtpCfg.SenderIdentity) {
			return ErrSMTPSenderIdentityEmpty(smtpCfg.Name)
		}

		if stringutil.IsEmptyString(smtpCfg.SenderEmail) {
			return ErrSMTPSenderEmailEmpty(smtpCfg.Name)
		}
	}

	for _, awsSES := range config.EmailServers.AWSSES {
		if stringutil.IsEmptyString(awsSES.Name) {
			return ErrAWSSESConfigNameReq
		}
		if nameCount[awsSES.Name] > 1 {
			return ErrAWSESConfigNameDup(awsSES.Name)
		}
		nameCount[awsSES.Name]++
		if stringutil.IsEmptyString(awsSES.AWSRegion) {
			return ErrAWSSESConfigRegionEmpty(awsSES.Name)
		}
		if stringutil.IsEmptyString(awsSES.AWSAccessKeyID) {
			return ErrAWSSESConfigAccessKeyIDEmpty(awsSES.Name)
		}

		if stringutil.IsEmptyString(awsSES.AWSSecretAccessKey) {
			return ErrAWSSESConfigSecretAccessKeyEmpty(awsSES.Name)
		}
	}

	return nil
}
