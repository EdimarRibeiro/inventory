package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/gorilla/mux"
)

type participantController struct {
	participant entitiesinterface.ParticipantRepositoryInterface
}

func CreateParticipantController(participantRep entitiesinterface.ParticipantRepositoryInterface) *participantController {
	return &participantController{participant: participantRep}
}

func (repo *participantController) getAll(tenantId uint64, search string, page int64, rows int64) (*models.ResponsePage, error) {
	condition := " and " + search
	if search == "" {
		condition = ""
	}
	participants, err := repo.participant.Search("Participant.TenantId=" + strconv.FormatUint(tenantId, 10) + condition)
	if err != nil {
		return nil, err
	}
	return common.PageResult(participants, page, rows), nil
}

func (repo *participantController) getById(tenantId uint64, id uint64) (*entities.Participant, error) {
	participants, err := repo.participant.Search("Participant.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Participant.Id=" + strconv.FormatUint(id, 10))
	if len(participants) == 0 || err != nil {
		return nil, err
	}
	return &participants[0], nil
}

func (repo *participantController) update(tenantId uint64, id uint64, participant *entities.Participant) error {
	participants, err := repo.participant.Search("Participant.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Participant.Id=" + strconv.FormatUint(id, 10))
	if len(participants) == 0 || err != nil {
		return err
	}

	var participantNew = participants[0]

	if len(participant.Name) > 0 {
		participantNew.Name = participant.Name
	}
	_, err = repo.participant.Save(&participantNew)
	return err
}

func (repo *participantController) save(participant *entities.Participant) error {
	_, err := repo.participant.Save(participant)
	return err
}

func (repo *participantController) delete(tenantId uint64, id uint64) error {
	participants, err := repo.participant.Search("Participant.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Participant.Id=" + strconv.FormatUint(id, 10))
	if len(participants) == 0 || err != nil {
		return err
	}
	//var participantNew = participants[0]
	//_, err = repo.participant.Delete(&participantNew)
	return err
}

func (repo *participantController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	search, page, rows := common.ExtractSearch(r)
	participants, err := repo.getAll(tenantId, search, page, rows)
	if err != nil {
		http.Error(w, "Error retrieving participant "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participants)
}

func (repo *participantController) GetByIdlHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	participantId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid participant id "+err.Error(), http.StatusBadRequest)
		return
	}

	participant, err := repo.getById(tenantId, participantId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if participant == nil {
		http.Error(w, "Not Found participant id ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participant)
}

func (repo *participantController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var createParticipant entities.Participant
	if err := json.NewDecoder(r.Body).Decode(&createParticipant); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	createParticipant.Id = 0
	createParticipant.TenantId = tenantId

	if err := repo.save(&createParticipant); err != nil {
		http.Error(w, "Error create participant "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *participantController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	participantId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid participant Id "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedParticipant entities.Participant
	if err := json.NewDecoder(r.Body).Decode(&updatedParticipant); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	participant, err := repo.getById(tenantId, participantId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if participant == nil {
		http.Error(w, "Not Found participant id ", http.StatusNotFound)
		return
	}

	if err := repo.update(tenantId, participantId, &updatedParticipant); err != nil {
		http.Error(w, "Error updating participant "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (repo *participantController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	participantId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid participant Id "+err.Error(), http.StatusBadRequest)
		return
	}
	participant, err := repo.getById(tenantId, participantId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if participant == nil {
		http.Error(w, "Not Found participant id ", http.StatusNotFound)
		return
	}
	if err := repo.delete(tenantId, participantId); err != nil {
		http.Error(w, "Error deleting participant "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
