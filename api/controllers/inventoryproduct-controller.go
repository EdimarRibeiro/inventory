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

type inventoryProductController struct {
	inventoryProduct entitiesinterface.InventoryProductRepositoryInterface
}

func CreateInventoryProductController(inventoryProductRep entitiesinterface.InventoryProductRepositoryInterface) *inventoryProductController {
	return &inventoryProductController{inventoryProduct: inventoryProductRep}
}

func (repo *inventoryProductController) GetAll(inventoryId uint64) ([]entities.InventoryProduct, error) {
	inventoryProducts, err := repo.inventoryProduct.Search("InventoryProduct.InventoryId=" + strconv.FormatUint(inventoryId, 10))
	if err != nil {
		return nil, err
	}
	return inventoryProducts, nil
}

func (repo *inventoryProductController) GetById(inventoryId uint64, productId uint64) (*entities.InventoryProduct, error) {
	inventoryProducts, err := repo.inventoryProduct.Search("InventoryProduct.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryProduct.ProductId=" + strconv.FormatUint(productId, 10))
	if len(inventoryProducts) == 0 || err != nil {
		return nil, err
	}
	return &inventoryProducts[0], nil
}

func (repo *inventoryProductController) Update(inventoryId uint64, productId uint64, inventoryProduct *entities.InventoryProduct) error {
	inventoryProducts, err := repo.inventoryProduct.Search("InventoryProduct.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryProduct.ProductId=" + strconv.FormatUint(productId, 10))
	if len(inventoryProducts) == 0 || err != nil {
		return err
	}

	var inventoryProductNew = inventoryProducts[0]

	if len(inventoryProduct.AccountingCode) > 0 {
		inventoryProductNew.AccountingCode = inventoryProduct.AccountingCode
	}
	if len(inventoryProduct.Complement) > 0 {
		inventoryProductNew.Complement = inventoryProduct.Complement
	}
	if len(inventoryProduct.OriginCode) > 0 {
		inventoryProductNew.OriginCode = inventoryProduct.OriginCode
	}
	if len(inventoryProduct.PossessionCode) > 0 {
		inventoryProductNew.PossessionCode = inventoryProduct.PossessionCode
	}
	if len(inventoryProduct.UnitId) > 0 {
		inventoryProductNew.UnitId = inventoryProduct.UnitId
	}
	if inventoryProduct.BalanceQuantity > 0 {
		inventoryProductNew.BalanceQuantity = inventoryProduct.BalanceQuantity
	}
	if inventoryProduct.InputQuantity > 0 {
		inventoryProductNew.InputQuantity = inventoryProduct.InputQuantity
	}
	if inventoryProduct.OutputQuantity > 0 {
		inventoryProductNew.OutputQuantity = inventoryProduct.OutputQuantity
	}
	if inventoryProduct.ParticipantId != nil && *inventoryProduct.ParticipantId > 0 {
		inventoryProductNew.ParticipantId = inventoryProduct.ParticipantId
	}
	if inventoryProduct.Quantity > 0 {
		inventoryProductNew.Quantity = inventoryProduct.Quantity
	}
	if inventoryProduct.Value > 0 {
		inventoryProductNew.Value = inventoryProduct.Value
	}
	if inventoryProduct.ValueIr > 0 {
		inventoryProductNew.ValueIr = inventoryProduct.ValueIr
	}
	if inventoryProduct.ValueTotal > 0 {
		inventoryProductNew.ValueTotal = inventoryProduct.ValueTotal
	}
	_, err = repo.inventoryProduct.Save(&inventoryProductNew)
	return err
}

func (repo *inventoryProductController) Delete(inventoryId uint64, productId uint64) error {
	inventoryProducts, err := repo.inventoryProduct.Search("InventoryProduct.InventoryId=" + strconv.FormatUint(inventoryId, 10) + " and InventoryProduct.ProductId=" + strconv.FormatUint(productId, 10))
	if len(inventoryProducts) == 0 || err != nil {
		return err
	}
	var inventoryProductNew = inventoryProducts[0]

	_, err = repo.inventoryProduct.Save(&inventoryProductNew)
	return err
}

func (repo *inventoryProductController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
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

	inventoryProducts, err := repo.GetAll(inventoryId)
	if err != nil {
		http.Error(w, "Error retrieving inventoryProduct "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryProducts)
}

func (repo *inventoryProductController) GetByIdlHandler(w http.ResponseWriter, r *http.Request) {
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

	productId, err := strconv.ParseUint(params["productId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryProduct productId "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryProduct, err := repo.GetById(inventoryId, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryProduct == nil {
		http.Error(w, "Not Found inventoryProduct productId ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryProduct)
}
func (repo *inventoryProductController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	productId, err := strconv.ParseUint(params["productId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryProduct productId "+err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Invalid inventoryProduct Id "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedInventoryProduct entities.InventoryProduct
	if err := json.NewDecoder(r.Body).Decode(&updatedInventoryProduct); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryProduct, err := repo.GetById(inventoryId, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryProduct == nil {
		http.Error(w, "Not Found inventoryProduct id ", http.StatusNotFound)
		return
	}

	if err := repo.Update(inventoryId, productId, &updatedInventoryProduct); err != nil {
		http.Error(w, "Error updating inventoryProduct "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repo *inventoryProductController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	productId, err := strconv.ParseUint(params["productId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid inventoryProduct productId "+err.Error(), http.StatusBadRequest)
		return
	}

	inventoryProduct, err := repo.GetById(inventoryId, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inventoryProduct == nil {
		http.Error(w, "Not Found inventoryProduct id ", http.StatusNotFound)
		return
	}

	if err := repo.Delete(inventoryId, productId); err != nil {
		http.Error(w, "Error deleting inventoryProduct "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
