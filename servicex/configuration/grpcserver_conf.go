package configuration

import (
	"bytes"
	"fmt"
)

type GRPCServerConfig struct{
	Port int
}

func (gserv *GRPCServerConfig)String() string{
	buffer := bytes.Buffer{}
	buffer.WriteString("### GRPC Server ###\n")
	buffer.WriteString(fmt.Sprintf("\tPort: %d\n",gserv.Port))
	return buffer.String()
}
