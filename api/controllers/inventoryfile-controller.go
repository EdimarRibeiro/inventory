package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/gorilla/mux"
)

type inventoryfileController struct {
	inventoryfile entitiesinterface.InventoryFileRepositoryInterface
}

func CreateInventoryFileController(inventoryfileRep entitiesinterface.InventoryFileRepositoryInterface) *inventoryfileController {
	return &inventoryfileController{inventoryfile: inventoryfileRep}
}

func (repo *inventoryfileController) GetAll(inventoryId uint64) ([]entities.InventoryFile, error) {
	inventoryfiles, err := repo.inventoryfile.Search("InventoryFile.InventoryId=" + strconv.FormatUint(inventoryId, 10))
	if err != nil {
		return nil, err
	}
	return inventoryfiles, nil
}

func (repo *inventoryfileController) GetById(inventoryId uint64, id uint64) (*entities.InventoryFile, error) {
	inventoryfiles, err := repo.inventoryfile.Search("InventoryFile.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryFile.Id=" + strconv.FormatUint(id, 10))
	if len(inventoryfiles) == 0 || err != nil {
		return nil, err
	}
	return &inventoryfiles[0], nil
}

func (repo *inventoryfileController) Update(inventoryId uint64, id uint64, inventoryfile *entities.InventoryFile) error {
	inventoryfiles, err := repo.inventoryfile.Search("InventoryFile.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryFile.Id=" + strconv.FormatUint(id, 10))
	if len(inventoryfiles) == 0 || err != nil {
		return err
	}

	var inventoryfileNew = inventoryfiles[0]

	if len(inventoryfile.FileName) > 0 {
		inventoryfileNew.FileName = inventoryfile.FileName
	}
	if len(inventoryfile.FileType) > 0 {
		inventoryfileNew.FileType = inventoryfile.FileType
	}
	_, err = repo.inventoryfile.Save(&inventoryfileNew)
	return err
}

func (repo *inventoryfileController) Delete(inventoryId uint64, id uint64) error {
	inventoryfiles, err := repo.inventoryfile.Search("InventoryFile.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryFile.Id=" + strconv.FormatUint(id, 10))
	if len(inventoryfiles) == 0 || err != nil {
		return err
	}
	var inventoryfileNew = inventoryfiles[0]
	inventoryfileNew.Processed = true

	_, err = repo.inventoryfile.Save(&inventoryfileNew)
	return err
}

func (repo *inventoryfileController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["inventoryId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryfiles, err := repo.GetAll(inventoryId)
	if err != nil {
		http.Error(w, "Error retrieving inventoryfile "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryfiles)
}

func (repo *inventoryfileController) GetByIdlHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["inventoryId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryfile id "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryfile, err := repo.GetById(inventoryId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryfile == nil {
		http.Error(w, "Not Found inventoryfile id ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryfile)
}
func (repo *inventoryfileController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["inventoryId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryfile id "+err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Invalid inventoryfile Id "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedInventoryFile entities.InventoryFile
	if err := json.NewDecoder(r.Body).Decode(&updatedInventoryFile); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryfile, err := repo.GetById(inventoryId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryfile == nil {
		http.Error(w, "Not Found inventoryfile id ", http.StatusNotFound)
		return
	}

	if err := repo.Update(inventoryId, id, &updatedInventoryFile); err != nil {
		http.Error(w, "Error updating inventoryfile "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *inventoryfileController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["inventoryId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryfile id "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryfile, err := repo.GetById(inventoryId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryfile == nil {
		http.Error(w, "Not Found inventoryfile id ", http.StatusNotFound)
		return
	}

	if err := repo.Delete(inventoryId, id); err != nil {
		http.Error(w, "Error deleting inventoryfile "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
