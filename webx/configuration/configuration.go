package configuration

import (
	"bytes"
	"errors"
	"fmt"
)

//Configuration : structure contain all configuration
type Configuration struct {
	WebApp     WebAppConfig
	Log        LogConfig
	SecretKey  string
	Parameters map[string]interface{}
}

func (config *Configuration) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(config.WebApp.String() + "\n")
	buffer.WriteString("### Log ###\n")
	buffer.WriteString("\t" + config.Log.String() + "\n")
	buffer.WriteString(fmt.Sprintf("\nSecretKey: %s\n", config.SecretKey))
	if len(config.Parameters) > 0 {
		buffer.WriteString("### Parameters ###\n")
		for paramKey, paramVal := range config.Parameters {
			buffer.WriteString(fmt.Sprintf("\t%s = %v\n",paramKey, paramVal))
		}
	}

	return buffer.String()
}

var (
	ErrNotfoundParamValue = func(key string) error{return fmt.Errorf("parameter vaulue from key %s not found", key)}
	ErrorKeyIsReq = errors.New("key is require")
)

func (config *Configuration)GetParamsStr(key string)(string, error){
	if len(config.Parameters) > 0 {
		if valueStr, ok := config.Parameters[key].(string); ok{
			return valueStr, nil
		}else{
			return "", ErrNotfoundParamValue(key)
		}
	}else{
		return "", ErrNotfoundParamValue(key)
	}
}

func (config *Configuration)GetParamsBool(key string)(bool, error){
	if len(config.Parameters) > 0 {
		if valueBool, ok := config.Parameters[key].(bool); ok{
			return valueBool, nil
		}else{
			return false, ErrNotfoundParamValue(key)
		}
	}else{
		return false, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration)GetParamsInt64(key string)(int64, error){
	if len(config.Parameters) > 0 {
		if valueInt64, ok := config.Parameters[key].(int64); ok{
			return valueInt64, nil
		}else{
			return -1, ErrNotfoundParamValue(key)
		}
	}else{
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration)GetParamsFloat64(key string)(float64, error){
	if len(config.Parameters) > 0 {
		if valueFloat64, ok := config.Parameters[key].(float64); ok{
			return valueFloat64, nil
		}else{
			return -1, ErrNotfoundParamValue(key)
		}
	}else{
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration)GetParamsInt(key string)(int, error){
	if len(config.Parameters) > 0 {
		if valueInt, ok := config.Parameters[key].(int); ok{
			return valueInt, nil
		}else{
			return -1, ErrNotfoundParamValue(key)
		}
	}else{
		return -1, ErrNotfoundParamValue(key)
	}
}

func (config *Configuration)GetParamsFloat32(key string)(float32, error){
	if len(config.Parameters) > 0 {
		if valueFloat32, ok := config.Parameters[key].(float32); ok{
			return valueFloat32, nil
		}else{
			return -1, ErrNotfoundParamValue(key)
		}
	}else{
		return -1, ErrNotfoundParamValue(key)
	}
}


func (config *Configuration)GetParams(key string)(interface{}, error){
	if len(config.Parameters) > 0 {
		if valueInterface , ok := config.Parameters[key]; ok{
			return valueInterface, nil
		}else{
			return false, ErrNotfoundParamValue(key)
		}
	}else{
		return false, ErrNotfoundParamValue(key)
	}
}
