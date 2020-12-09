package boundaries

import "github.com/google/uuid"

type CommonInfrastructure interface {
	NewUUID() uuid.UUID
	DecodeBase64String(s string) ([]byte, error)
}
