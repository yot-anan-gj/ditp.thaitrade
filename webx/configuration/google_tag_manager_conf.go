package configuration

import (
	"bytes"
	"fmt"
)

type WebAppConfig struct {
	Statics                  []map[string]string
	Port                     int
	HealthPort               int
	K8SZeroDownTimeThreshold int //>= periodSeconds * failureThreshold
	//default > 10
	GracefulShutdownTimeout int
	HomeURL                 string
	CORs                    WebAppCORsConfig
	CSRF                    WebAppCSRFConfig
	SessionStore            WebAppSessionConfig
	GoogleTagManager        GoogleTagManagerConfig
	Databases               []WebAppDBConfig
}

func (wc *WebAppConfig) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("### WebServer ### \n"))
	buffer.WriteString("\tResources:\n")
	for _, mapStatic := range wc.Statics {
		for key, val := range mapStatic {
			buffer.WriteString(fmt.Sprintf("\t\t- %s <-- %s\n", key, val))
		}
	}
	buffer.WriteString(fmt.Sprintf("\tPort: %d\n", wc.Port))
	buffer.WriteString(fmt.Sprintf("\tHealthPort: %d\n", wc.HealthPort))
	buffer.WriteString(fmt.Sprintf("\tK8SZeroDownTimeThreshold: %d\n", wc.K8SZeroDownTimeThreshold))
	buffer.WriteString(fmt.Sprintf("\tGracefulShutdownTimeout: %d\n", wc.GracefulShutdownTimeout))
	buffer.WriteString(fmt.Sprintf("\tHomeURL: %s\n", wc.HomeURL))
	buffer.WriteString(fmt.Sprintf("\tCORs:\n\t\t%s\n", wc.CORs.String()))
	buffer.WriteString(fmt.Sprintf("\tCSRF:\n\t\t%s\n", wc.CSRF.String()))
	buffer.WriteString(fmt.Sprintf("\tGoogleTagManager:\n\t\t%s\n", wc.GoogleTagManager.String()))
	buffer.WriteString(fmt.Sprintf("\tSessionStore:\n"))
	buffer.WriteString(fmt.Sprintf("\t\tRedisStores:\n"))
	for _, redisStore := range wc.SessionStore.RedisStores {
		buffer.WriteString(fmt.Sprintf("\t\t\t%s\n", redisStore.String()))
	}
	buffer.WriteString(fmt.Sprint("\tDatabases: \n"))
	for _, database := range wc.Databases {
		buffer.WriteString("\t\t" + database.String() + "\n")
	}
	return buffer.String()

}
