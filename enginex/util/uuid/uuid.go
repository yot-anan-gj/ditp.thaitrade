package uuid

import "github.com/satori/go.uuid"

func UUIDv4() string{
	id := uuid.NewV4()
	return id.String()
}

func UUIDv4Raw()[uuid.Size]byte{
	return uuid.NewV4()
}

