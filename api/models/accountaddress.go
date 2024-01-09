package models

type AccountAddress struct {
	CountryId    string `json:"countryId"`
	CityId       string `json:"cityId"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complememt   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	ZipCode      string `json:"zipCode"`
}
