package configuration

import (
	"bytes"
	"errors"
	"fmt"
)

type Configuration struct {
	WebAPI                   WebAPIConfig
	GRPCServer               GRPCServerConfig
	HealthPort               int
	K8SZeroDownTimeThreshold int //>= periodSeconds * failureThreshold
	//default > 10
	GracefulShutdownAPITimeout int
	ExternalAuthentications    []ExternalAuthenticationConfig
	Log                        LogConfig
	SecretKey                  string
	Databases                  []DBConfig
	RedisCaches                []RedisCacheConfig
	EmailServers               EmailServerConfig
	Parameters                 map[string]interface{}
	DynamoDBs                   []DynamoDBConfig
	Sqss                        []SqsConfig
}

func (config *Configuration) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(config.WebAPI.String() + "\n")
	buffer.WriteString(config.GRPCServer.String() + "\n")
	buffer.WriteString(fmt.Sprintf("HealthPort: %d\n", config.HealthPort))
	buffer.WriteString(fmt.Sprintf("K8SZeroDownTimeThreshold: %d\n", config.K8SZeroDownTimeThreshold))
	buffer.WriteString(fmt.Sprintf("GracefulShutdownAPITimeout: %d\n", config.GracefulShutdownAPITimeout))
	buffer.WriteString("### ExternalAuthentications ###\n")
	for _, extAuthen := range config.ExternalAuthentications {
		buffer.WriteString(fmt.Sprintf("\t%s\n", extAuthen.String()))
	}
	buffer.WriteString("### Log ###\n")
	buffer.WriteString("\t" + config.Log.String() + "\n")
	buffer.WriteString(fmt.Sprintf("\nSecretKey: %s\n", config.SecretKey))
	buffer.WriteString("### Databases ###\n")
	for _, dbCfg := range config.Databases {
		buffer.WriteString(fmt.Sprintf("\t%s\n", dbCfg.String()))
	}
	buffer.WriteString("### DynamoDBs ###\n")
	for _, dynamoCfg := range config.DynamoDBs {
		buffer.WriteString(fmt.Sprintf("\t%s\n", dynamoCfg.String()))
	}
	buffer.WriteString("### Sqss ###\n")
	for _, sqsCfg := range config.Sqss {
		buffer.WriteString(fmt.Sprintf("\t%s\n", sqsCfg.String()))
	}
	buffer.WriteString("### RedisCaches ###\n")
	for _, cache := range config.RedisCaches {
		buffer.WriteString(fmt.Sprintf("\t%s\n", cache.String()))
	}

	buffer.WriteString("### Email Servers ###\n")
	if len(config.EmailServers.SMTPs) > 0 {
		buffer.WriteString("\tSMTPs:\n")
		for _, smtp := range config.EmailServers.SMTPs {
			buffer.WriteString(fmt.Sprintf("\t\t%s\n", smtp.String()))

		}
	}
	if len(config.EmailServers.AWSSES) > 0 {
		buffer.WriteString("\tAWSSES:\n")
		for _, awsses := range config.EmailServers.AWSSES {
			buffer.WriteString(fmt.Sprintf("\t\t%s\n", awsses.String()))

		}
	}

	if len(config.Parameters) > 0 {
		buffer.WriteString("### Parameters ###\n")
		for paramKey, paramVal := range config.Parameters {
			buffer.WriteString(fmt.Sprintf("\t%s = %v\n", paramKey, paramVal))
		}
	}

	return buffer.String()
}

var (
	ErrNotfoundParamValue       = func(key string) error { return fmt.Errorf("parameter value from key %s not found", key) }
	ErrorKeyIsReq               = errors.New("key is require")
	ErrNotfoundSMTPServerName   = func(key string) error { return fmt.Errorf("SMTP Server configuratin %s not found", key) }
	ErrNotfoundAWSSESServerName = func(key string) error { return fmt.Errorf("AWS SES configuratin %s not found", key) }
)

func (config *Configuration) GetParamsStr(key string) (string, error) {
	if len(config.Parameters) > 0 {
		if valueStr, ok := config.Parameters[key].(string); ok {
			return valueStr, nil
		} else {
			return "", ErrNotfoundParamValue(key)
		}
	} else {
		return "", ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParamsBool(key string) (bool, error) {
	if len(config.Parameters) > 0 {
		if valueBool, ok := config.Parameters[key].(bool); ok {
			return valueBool, nil
		} else {
			return false, ErrNotfoundParamValue(key)
		}
	} else {
		return false, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParamsInt64(key string) (int64, error) {
	if len(config.Parameters) > 0 {
		if valueInt64, ok := config.Parameters[key].(int64); ok {
			return valueInt64, nil
		} else {
			return -1, ErrNotfoundParamValue(key)
		}
	} else {
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParamsFloat64(key string) (float64, error) {
	if len(config.Parameters) > 0 {
		if valueFloat64, ok := config.Parameters[key].(float64); ok {
			return valueFloat64, nil
		} else {
			return -1, ErrNotfoundParamValue(key)
		}
	} else {
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParamsInt(key string) (int, error) {
	if len(config.Parameters) > 0 {
		if valueInt, ok := config.Parameters[key].(int); ok {
			return valueInt, nil
		} else {
			return -1, ErrNotfoundParamValue(key)
		}
	} else {
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParamsFloat32(key string) (float32, error) {
	if len(config.Parameters) > 0 {
		if valueFloat32, ok := config.Parameters[key].(float32); ok {
			return valueFloat32, nil
		} else {
			return -1, ErrNotfoundParamValue(key)
		}
	} else {
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetParams(key string) (interface{}, error) {
	if len(config.Parameters) > 0 {
		if valueInterface, ok := config.Parameters[key]; ok {
			return valueInterface, nil
		} else {
			return false, ErrNotfoundParamValue(key)
		}
	} else {
		return false, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration) GetSMTPEmailServerConfig(name string) (SMTPConfig, error) {
	if len(config.EmailServers.SMTPs) > 0 {
		for _, smtpCfg := range config.EmailServers.SMTPs {
			if smtpCfg.Name == name {
				return smtpCfg, nil
			}
		}
	}

	return SMTPConfig{}, ErrNotfoundSMTPServerName(name)
}

func (config *Configuration) GetSMTPAWSSESServerConfig(name string) (AWSSESConfig, error) {
	if len(config.EmailServers.AWSSES) > 0 {
		for _, awscfg := range config.EmailServers.AWSSES {
			if awscfg.Name == name {
				return awscfg, nil
			}
		}
	}

	return AWSSESConfig{}, ErrNotfoundAWSSESServerName(name)
}
