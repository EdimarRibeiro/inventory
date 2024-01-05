package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
)

type cityController struct {
	city entitiesinterface.CityRepositoryInterface
}

func CreateCityController(cityRep entitiesinterface.CityRepositoryInterface) *cityController {
	return &cityController{city: cityRep}
}

func (repo *cityController) GetAll(tenantId uint64) ([]entities.City, error) {
	citys, err := repo.city.Search("")
	if err != nil {
		return nil, err
	}
	return citys, nil
}

func (repo *cityController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	citys, err := repo.GetAll(tenantId)
	if err != nil {
		http.Error(w, "Error retrieving city "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(citys)
}
