package infrastructure

import (
	"encoding/base64"
	"github.com/google/uuid"
	boundaries "github.com/makushenk/gimage/boundaries/infrastructure"
)

type commonInfrastructure struct {}

func NewCommonInfrastructure() boundaries.CommonInfrastructure {
	return &commonInfrastructure{}
}

func (i *commonInfrastructure) NewUUID() uuid.UUID {
	return uuid.New()
}

func (i *commonInfrastructure) DecodeBase64String(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}