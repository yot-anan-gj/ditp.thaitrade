package configuration

import "fmt"

type EmailServerConfig struct {
	SMTPs  []SMTPConfig
	AWSSES [] AWSSESConfig
}

type SMTPConfig struct {
	Name           string
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	User           string
	Password       string
}

func (smtpCfg SMTPConfig) String() string {
	return fmt.Sprintf("Name: %s, Server: %s, Port: %d, SenderEmail: %s, SenderIdentity: %s, User: %s, Password: %s",
		smtpCfg.Name,
		smtpCfg.Server,
		smtpCfg.Port,
		smtpCfg.SenderEmail,
		smtpCfg.SenderIdentity,
		smtpCfg.User,
		smtpCfg.Password)

}

type AWSSESConfig struct {
	Name               string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

func (awssesCfg AWSSESConfig) String() string {
	return fmt.Sprintf("Name: %s, AWSRegion: %s, AWSAccessKeyID: %s, AWSSecretAccessKey: %s",
		awssesCfg.Name,
		awssesCfg.AWSRegion,
		awssesCfg.AWSAccessKeyID,
		awssesCfg.AWSSecretAccessKey)
}
