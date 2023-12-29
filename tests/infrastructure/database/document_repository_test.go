package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestValidDocumentCreateDataBase(t *testing.T) {
	database.Initialize(false)
	partRepo := &database.ParticipantRepository{DB: database.DB}
	invRepo := &database.InventoryRepository{DB: database.DB}
	tenRepo := &database.TenantRepository{DB: database.DB}
	perRepo := &database.PersonRepository{DB: database.DB}
	docRepo := &database.DocumentRepository{DB: database.DB}

	var tenantId uint64 = 0

	tens, err := tenRepo.Search("Id >= 1")
	if err != nil || len(tens) == 0 {
		ten, err := entities.NewTenant("Teste", "09066936754", 0)
		assert.NoError(t, err, nil)

		ten, err = tenRepo.Save(ten)
		assert.NoError(t, err, nil)
		if ten.Id == 0 {
			err = errors.New("invalid value ID")
			assert.NoError(t, err, nil)
		}

		ten.PersonId, err = GetPersonData(perRepo, ten.Id)
		assert.NoError(t, err, nil)

		ten, err = tenRepo.Save(ten)
		assert.NoError(t, err, nil)

		tens = append(tens, *ten)
	}

	if tens[0].Id == 0 {
		err = errors.New("invalid value ID")
		assert.NoError(t, err, nil)
	}
	tenantId = tens[len(tens)-1].Id

	invs, err := invRepo.Search("Id >= 1")
	assert.NoError(t, err, nil)
	if err != nil || len(invs) == 0 {
		inv, err := entities.NewInventory(tenantId, "Teste")
		assert.NoError(t, err, nil)

		inv, err = invRepo.Save(inv)
		assert.NoError(t, err, nil)
		invs = append(invs, *inv)
	}
	var inventoryId uint64 = invs[len(invs)-1].Id

	docs, err := docRepo.Search("Id >= 1")
	assert.NoError(t, err, nil)
	if len(docs) == 0 {
		participantId, err := GetParticipantData(partRepo, tenantId)
		assert.NoError(t, err, nil)

		model, err := entities.NewDocument(inventoryId, participantId, "", "", "", "", "", "", "", time.Now(), time.Now(), 0, "", 0, 0, 0, "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, false, "xml")
		assert.NoError(t, err, nil)

		doc, err := docRepo.Save(model)
		assert.NoError(t, err, nil)

		if doc.Id == 0 {
			err = errors.New("invalid id of document")
		}
		assert.NoError(t, err, nil)
		//docs = append(docs, *doc)
	}
}

func GetPersonData(tenPer *database.PersonRepository, tenantId uint64) (uint64, error) {
	pers, err := tenPer.Search("Id>=1")
	if err != nil || len(pers) == 0 {
		per, err := tenPer.Save(CreatePersonData(tenantId))
		if err != nil {
			return 0, err
		}
		pers = append(pers, *per)
	}
	return pers[len(pers)-1].Id, nil
}

func GetParticipantData(part *database.ParticipantRepository, tenantId uint64) (uint64, error) {
	pars, err := part.Search("Id>=1")
	if err != nil || len(pars) == 0 {
		par, err := part.Save(CreateParticipantData(tenantId))
		if err != nil {
			return 0, err
		}
		pars = append(pars, *par)
	}
	return pars[len(pars)-1].Id, nil
}

func CreatePersonData(tenantId uint64) *entities.Person {
	return &entities.Person{
		TenantId:     tenantId,
		Name:         "Teste",
		Document:     "09066936754",
		Registration: "",
		CountryId:    1,
		CityId:       1,
		Street:       "Rua Teste",
		Number:       "sem",
		Suframa:      "Teste",
		Complememt:   "",
		Neighborhood: "teste",
		ZipCode:      "29000000",
	}
}

func CreateParticipantData(tenantId uint64) *entities.Participant {
	return &entities.Participant{
		TenantId:     tenantId,
		Name:         "Teste",
		OriginCode:   "1",
		Document:     "09066936754",
		Registration: "",
		CountryCode:  "1058",
		CityCode:     "1",
		Street:       "Rua Teste",
		Number:       "sem",
		Suframa:      "Teste",
		Complememt:   "",
		Neighborhood: "teste",
	}
}
