package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/EdimarRibeiro/inventory/api/common"
	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/gorilla/mux"
)

type accountControler struct {
	user   entitiesinterface.UserRepositoryInterface
	tenant entitiesinterface.TenantRepositoryInterface
	person entitiesinterface.PersonRepositoryInterface
	city   entitiesinterface.CityRepositoryInterface
}

func CreateAccountController(tenantRep entitiesinterface.TenantRepositoryInterface, personRep entitiesinterface.PersonRepositoryInterface, userRep entitiesinterface.UserRepositoryInterface, cityRep entitiesinterface.CityRepositoryInterface) *accountControler {
	return &accountControler{user: userRep, tenant: tenantRep, person: personRep, city: cityRep}
}

func (account *accountControler) CreateNewAccount(newAccount models.Account) error {

	pers, err := account.person.Search("document='" + newAccount.Document + "'")

	if err != nil {
		return err
	}
	if len(pers) > 1 {
		return errors.New("unable to create this account, document already registered")
	}
	if newAccount.Email == "" {
		return errors.New("email is invalid")
	}
	users, err := account.user.Search("Login ='" + newAccount.Email + "' and EndDate is null")
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errors.New("email already registered")
	}
	ten, err := entities.NewTenant(newAccount.Name, newAccount.Document, 0)
	if err != nil {
		return err
	}
	ten, err = account.tenant.Save(ten)
	if err != nil {
		return err
	}
	countryId, err := strconv.ParseUint(newAccount.Address.CountryId, 10, 64)
	if err != nil {
		return err
	}
	cityId, err := strconv.ParseUint(newAccount.Address.CityId, 10, 64)
	if err != nil {
		return err
	}
	per, err := entities.NewPerson(ten.Id, newAccount.Name, newAccount.Document, newAccount.Registration, countryId, cityId, newAccount.Address.Street, newAccount.Address.Number, newAccount.Address.Complememt, newAccount.Address.Neighborhood, newAccount.Address.ZipCode)
	if err != nil {
		return err
	}
	per, err = account.person.Save(per)
	if err != nil {
		return err
	}
	ten.PersonId = per.Id
	_, err = account.tenant.Save(ten)
	if err != nil {
		return err
	}

	user, err := entities.NewUser(ten.Id, newAccount.NameUser, newAccount.Email, newAccount.Password)
	if err != nil {
		return err
	}
	_, err = account.user.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func replaceAllDoc(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(value, "-", ""), ".", ""), "/", "")
}

func (account *accountControler) GetZipCode(zipCode string) (*models.AccountAddress, error) {
	url := "https://viacep.com.br/ws/" + zipCode + "/json/"

	headers := map[string]string{}

	resp, err := common.Get(url, headers)

	if err != nil {
		return nil, err
	}

	var cepRes models.ResponseZipCode
	err = json.Unmarshal(resp, &cepRes)
	if err != nil {
		return nil, fmt.Errorf("error deserializing JSON response: %v", err)
	}

	cepRes.CEP = replaceAllDoc(cepRes.CEP)

	if cepRes.CEP != zipCode {
		return nil, errors.New("cep Notfound resp.:" + cepRes.CEP + " - req.:" + zipCode)
	}

	response := &models.AccountAddress{
		CountryId:    "1",
		CityId:       "0",
		Street:       cepRes.Logradouro,
		Number:       "",
		Complememt:   cepRes.Complemento,
		Neighborhood: cepRes.Bairro,
		ZipCode:      cepRes.CEP,
	}

	cityId, err := account.city.GetCityId(cepRes.IBGE)
	if err != nil {
		return nil, err
	}
	response.CityId = strconv.FormatUint(cityId, 10)

	return response, nil
}

func (account *accountControler) GetDocument(document string) (*models.Account, error) {
	url := "https://www.receitaws.com.br/v1/cnpj/" + document

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := common.Get(url, headers)

	if err != nil {
		return nil, err
	}

	var cnpjRes models.ResponseCnpj
	err = json.Unmarshal(resp, &cnpjRes)
	if err != nil {
		return nil, fmt.Errorf("error deserializing JSON response: %v", err)
	}

	cnpjRes.CNPJ = replaceAllDoc(cnpjRes.CNPJ)
	cnpjRes.CEP = replaceAllDoc(cnpjRes.CEP)

	if cnpjRes.CNPJ != document {
		return nil, errors.New("document Notfound " + document)
	}

	if cnpjRes.Situacao != "ATIVA" {
		return nil, errors.New("inactive document at SEFAZ, rejected ")
	}

	if cnpjRes.Status != "OK" {
		return nil, errors.New("document with status restriction at SEFAZ, rejected ")
	}

	response := &models.Account{
		Address: models.AccountAddress{
			CountryId:    "1",
			CityId:       "0",
			Street:       cnpjRes.Logradouro,
			Number:       cnpjRes.Numero,
			Complememt:   "",
			Neighborhood: cnpjRes.Bairro,
			ZipCode:      cnpjRes.CEP},
		Name:         cnpjRes.Nome,
		Document:     cnpjRes.CNPJ,
		Registration: "",
		NameUser:     "",
		Email:        "",
		Password:     "",
		Suframa:      "",
	}

	resCep, err := account.GetZipCode(cnpjRes.CEP)
	if err != nil {
		return nil, err
	}

	response.Address.CityId = resCep.CityId
	return response, nil
}

func (account *accountControler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var accountNew models.Account
	if err := json.NewDecoder(r.Body).Decode(&accountNew); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	err := account.CreateNewAccount(accountNew)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (account *accountControler) GetCepHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cep := params["cep"]
	if len(cep) != 8 {
		http.Error(w, "Invalid CEP "+cep, http.StatusBadRequest)
		return
	}
	resp, err := account.GetZipCode(cep)
	if err != nil {
		http.Error(w, "Error retrieving "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (account *accountControler) GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	document := params["document"]
	if len(document) != 14 {
		http.Error(w, "Invalid document "+document, http.StatusBadRequest)
		return
	}
	resp, err := account.GetDocument(document)
	if err != nil {
		http.Error(w, "Error retrieving "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
