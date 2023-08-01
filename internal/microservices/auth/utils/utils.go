package utils

import (
	auth_proto "store-project/internal/microservices/auth/proto"
	"store-project/internal/models"
)

func ProtoBySessionModel(s *models.Session) *auth_proto.Session {
	return &auth_proto.Session{
		UID:   s.UID,
		Value: s.SessionID,
	}
}
