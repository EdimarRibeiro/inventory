package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
)

type ParticipantRepositoryInterface interface {
	Save(Participant *entities.Participant) (*entities.Participant, error)
	GetParticipantId(value string, err error) (uint64, error)
	Search(where string) ([]entities.Participant, error)
}
