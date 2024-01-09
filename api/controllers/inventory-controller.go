package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/gorilla/mux"
)

type inventoryProcessController struct {
	InventoryFile    entitiesinterface.InventoryFileRepositoryInterface
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
	Unit             entitiesinterface.UnitRepositoryInterface
	UnitConvert      entitiesinterface.UnitConvertRepositoryInterface
	Participant      entitiesinterface.ParticipantRepositoryInterface
	Product          entitiesinterface.ProductRepositoryInterface
	Document         entitiesinterface.DocumentRepositoryInterface
	DocumentItem     entitiesinterface.DocumentItemRepositoryInterface
}

type inventoryProcessCalcController struct {
	Inventory        entitiesinterface.InventoryRepositoryInterface
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
	DocumentItem     entitiesinterface.DocumentItemRepositoryInterface
}

type inventoryController struct {
	inventory entitiesinterface.InventoryRepositoryInterface
}

func CreateInventoryController(inventoryRep entitiesinterface.InventoryRepositoryInterface) *inventoryController {
	return &inventoryController{inventory: inventoryRep}
}

func CreateInventoryProcessController(inventoryFile entitiesinterface.InventoryFileRepositoryInterface,
	inventoryProduct entitiesinterface.InventoryProductRepositoryInterface,
	unit entitiesinterface.UnitRepositoryInterface,
	unitConvert entitiesinterface.UnitConvertRepositoryInterface,
	participant entitiesinterface.ParticipantRepositoryInterface,
	product entitiesinterface.ProductRepositoryInterface,
	document entitiesinterface.DocumentRepositoryInterface,
	documentItem entitiesinterface.DocumentItemRepositoryInterface) *inventoryProcessController {
	return &inventoryProcessController{
		InventoryFile:    inventoryFile,
		InventoryProduct: inventoryProduct,
		Unit:             unit,
		UnitConvert:      unitConvert,
		Participant:      participant,
		Product:          product,
		Document:         document,
		DocumentItem:     documentItem,
	}
}

func CreateInventoryProcessCalcController(inventory entitiesinterface.InventoryRepositoryInterface,
	inventoryProduct entitiesinterface.InventoryProductRepositoryInterface,
	documentItem entitiesinterface.DocumentItemRepositoryInterface) *inventoryProcessCalcController {
	return &inventoryProcessCalcController{
		Inventory:        inventory,
		InventoryProduct: inventoryProduct,
		DocumentItem:     documentItem,
	}
}

func (repo *inventoryController) GetAll(tenantId uint64) ([]entities.Inventory, error) {
	inventorys, err := repo.inventory.Search("Inventory.TenantId=" + strconv.FormatUint(tenantId, 10))
	if err != nil {
		return nil, err
	}
	return inventorys, nil
}

func (repo *inventoryController) GetById(tenantId uint64, id uint64) (*entities.Inventory, error) {
	inventorys, err := repo.inventory.Search("Inventory.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Inventory.Id=" + strconv.FormatUint(id, 10))
	if len(inventorys) == 0 || err != nil {
		return nil, err
	}
	return &inventorys[0], nil
}

func (repo *inventoryController) Update(tenantId uint64, id uint64, inventory *entities.Inventory) error {
	inventorys, err := repo.inventory.Search("Inventory.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Inventory.Id=" + strconv.FormatUint(id, 10))
	if len(inventorys) == 0 || err != nil {
		return err
	}

	var inventoryNew = inventorys[0]

	if len(inventory.Name) > 0 {
		inventoryNew.Name = inventory.Name
	}
	_, err = repo.inventory.Save(&inventoryNew)
	return err
}

func (repo *inventoryController) Delete(tenantId uint64, id uint64) error {
	inventorys, err := repo.inventory.Search("Inventory.TenantId=" + strconv.FormatUint(tenantId, 10) + " and Inventory.Id=" + strconv.FormatUint(id, 10))
	if len(inventorys) == 0 || err != nil {
		return err
	}
	var inventoryNew = inventorys[0]
	inventoryNew.Cloused = true

	_, err = repo.inventory.Save(&inventoryNew)
	return err
}

func (repo *inventoryController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	inventorys, err := repo.GetAll(tenantId)
	if err != nil {
		http.Error(w, "Error retrieving inventory "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventorys)
}

func (repo *inventoryController) GetByIdlHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	inventory, err := repo.GetById(tenantId, inventoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventory == nil {
		http.Error(w, "Not Found inventory id ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}
func (repo *inventoryController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory Id "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedInventory entities.Inventory
	if err := json.NewDecoder(r.Body).Decode(&updatedInventory); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	inventory, err := repo.GetById(tenantId, inventoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventory == nil {
		http.Error(w, "Not Found inventory id ", http.StatusNotFound)
		return
	}

	if err := repo.Update(tenantId, inventoryId, &updatedInventory); err != nil {
		http.Error(w, "Error updating inventory "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (repo *inventoryController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory Id "+err.Error(), http.StatusBadRequest)
		return
	}
	inventory, err := repo.GetById(tenantId, inventoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventory == nil {
		http.Error(w, "Not Found inventory id ", http.StatusNotFound)
		return
	}
	if err := repo.Delete(tenantId, inventoryId); err != nil {
		http.Error(w, "Error deleting inventory "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (repo *inventoryProcessController) InventaryProcessFileHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory Id "+err.Error(), http.StatusBadRequest)
		return
	}
	invs, err := repo.InventoryFile.Search("InventoryFile.InventoryId = " + strconv.FormatUint(inventoryId, 10) + " and InventoryFile.FileType ='txt'")
	if err != nil {
		http.Error(w, " inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(invs) == 0 {
		http.Error(w, "Not Found inventory id ", http.StatusBadRequest)
		return
	}

	impt := infrastructure.CreateImportFileText(repo.InventoryFile, repo.InventoryProduct, repo.Unit, repo.UnitConvert, repo.Participant, repo.Product, repo.Document, repo.DocumentItem)

	for i := 0; i < len(invs); i++ {
		item := invs[i]
		err = impt.Execute(&item)

		if err != nil {
			http.Error(w, "Erro in process TXT file", http.StatusInternalServerError)
			return
		}
	}

	/*****XML*****/
	invs, err = repo.InventoryFile.Search("InventoryFile.InventoryId = " + strconv.FormatUint(inventoryId, 10) + " and InventoryFile.FileType ='xml'")
	if err != nil {
		http.Error(w, " XML inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(invs) == 0 {
		http.Error(w, "Not Found XML inventory id ", http.StatusBadRequest)
		return
	}
	impXml := infrastructure.CreateImportFileXml(repo.InventoryFile, repo.InventoryProduct, repo.Unit, repo.UnitConvert, repo.Participant, repo.Product, repo.Document, repo.DocumentItem)

	for i := 0; i < len(invs); i++ {
		item := invs[i]

		err = impXml.Execute(&item)

		if err != nil {
			http.Error(w, "Erro in process XML file", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *inventoryProcessCalcController) InventaryProcessCalcHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	inventoryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventory Id "+err.Error(), http.StatusBadRequest)
		return
	}
	inves, err := repo.Inventory.Search("Inventory.Id = " + strconv.FormatUint(inventoryId, 10) + " and Inventory.Cloused=0")
	if err != nil {
		http.Error(w, " inventory id "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(inves) == 0 {
		http.Error(w, "Not Found inventory id ", http.StatusBadRequest)
		return
	}

	calc := infrastructure.CreateCalculateBalanceQuantityData(repo.InventoryProduct, repo.DocumentItem)

	err = calc.Execute(inves[0].Id)
	if err != nil {
		http.Error(w, " Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
