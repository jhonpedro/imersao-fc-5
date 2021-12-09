package services

import (
	"github.com/google/uuid"
)

type UniqueIdentifierService struct {
}

func NewUniqueIdentifierService() *UniqueIdentifierService {
	return &UniqueIdentifierService{}
}

func (u *UniqueIdentifierService) Generate() string {
	uuid, _ := uuid.NewUUID()

	return uuid.String()
}
