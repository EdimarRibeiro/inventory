package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/gorilla/mux"
)

type userController struct {
	user entitiesinterface.UserRepositoryInterface
}

func CreateUserController(userRep entitiesinterface.UserRepositoryInterface) *userController {
	return &userController{user: userRep}
}

func (repo *userController) GetAll(tenantId uint64) ([]entities.User, error) {
	users, err := repo.user.Search("TenantId=" + strconv.FormatUint(tenantId, 10) + " and EndDate is null")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userController) GetById(tenantId uint64, id uint64) (*entities.User, error) {
	users, err := repo.user.Search("TenantId=" + strconv.FormatUint(tenantId, 10) + " and Id=" + strconv.FormatUint(id, 10) + " and EndDate is null")
	if len(users) == 0 || err != nil {
		return nil, err
	}
	return &users[0], nil
}

func (repo *userController) Update(tenantId uint64, id uint64, user *entities.User) error {
	users, err := repo.user.Search("TenantId=" + strconv.FormatUint(tenantId, 10) + " and Id=" + strconv.FormatUint(id, 10) + " and EndDate is null")
	if len(users) == 0 || err != nil {
		return err
	}

	var userNew = users[0]

	if len(user.Name) > 0 {
		userNew.Name = user.Name
	}

	if len(user.Password) > 0 {
		userNew.Password = user.Password
	}
	_, err = repo.user.Save(&userNew)
	return err
}

func (repo *userController) Delete(tenantId uint64, id uint64) error {
	users, err := repo.user.Search("TenantId=" + strconv.FormatUint(tenantId, 10) + " and Id=" + strconv.FormatUint(id, 10) + " and EndDate is null")
	if len(users) == 0 || err != nil {
		return err
	}
	var endDate = time.Now()
	var userNew = users[0]
	userNew.EndDate = &endDate

	_, err = repo.user.Save(&userNew)
	return err
}

func (repo *userController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	users, err := repo.GetAll(tenantId)
	if err != nil {
		http.Error(w, "Error retrieving user "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (repo *userController) GetByIdlHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repo.GetById(tenantId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Not Found user id ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func (repo *userController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user Id "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUser entities.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repo.GetById(tenantId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Not Found user id ", http.StatusNotFound)
		return
	}

	if err := repo.Update(tenantId, userId, &updatedUser); err != nil {
		http.Error(w, "Error updating user "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (repo *userController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user Id "+err.Error(), http.StatusBadRequest)
		return
	}
	user, err := repo.GetById(tenantId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Not Found user id ", http.StatusNotFound)
		return
	}
	if err := repo.Delete(tenantId, userId); err != nil {
		http.Error(w, "Error deleting user "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
