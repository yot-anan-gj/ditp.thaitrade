package configuration

import "fmt"

type GoogleTagManagerConfig struct {
	ContainerID string
}

func (gtm GoogleTagManagerConfig) String() string {
	return fmt.Sprintf("ContainerID: %s",
		gtm.ContainerID)
}
