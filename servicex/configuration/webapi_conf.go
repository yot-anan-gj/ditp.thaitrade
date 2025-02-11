package configuration

import (
	"bytes"
	"fmt"
)

type WebAPIConfig struct {
	Port         int
	CORs         WebAPICORsConfig
	CSRF         WebAPICSRFConfig
	SessionStore WebAPISessionConfig
}

func (webAPI *WebAPIConfig) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("### WebAPI ### \n"))
	buffer.WriteString(fmt.Sprintf("\tPort: %d\n", webAPI.Port))
	buffer.WriteString(fmt.Sprintf("\tCORs:\n\t\t%s\n", webAPI.CORs.String()))
	buffer.WriteString(fmt.Sprintf("\tCSRF:\n\t\t%s\n", webAPI.CSRF.String()))
	buffer.WriteString(fmt.Sprintf("\tSessionStore:\n"))
	buffer.WriteString(fmt.Sprintf("\t\tRedisStores:\n"))
	for _, redisStore := range webAPI.SessionStore.RedisStores{
		buffer.WriteString(fmt.Sprintf("\t\t\t%s\n",redisStore.String()))
	}
	return buffer.String()

}
