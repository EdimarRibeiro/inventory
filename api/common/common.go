package common

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(r *http.Request) (string, uint64, error) {
	tokenString := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")

	if tokenString == "" {
		return "", 0, fmt.Errorf("authorization token not provided")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return models.SecretKey, nil
	})

	if err != nil {
		return "", 0, err
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || !token.Valid {
		return "", 0, fmt.Errorf("invalid token")
	}

	return claims.Username, claims.ExternalId, nil
}
func ExtractSearch(r *http.Request) (string, int64, int64) {
	search := strings.ReplaceAll(r.FormValue("search"), "undefined", "")
	page, err := strconv.ParseInt(r.FormValue("page"), 10, 64)
	if err != nil || page == 0 {
		page = 1
	}
	rows, err := strconv.ParseInt(r.FormValue("rows"), 10, 64)
	if err != nil {
		rows = 20
	}
	return search, page, rows
}

func PageResult[T any](dataSet []T, page int64, rows int64) *models.ResponsePage {

	if page == 0 {
		page = 1
	}
	totalRecords := int64(len(dataSet))
	pages := int64(math.Ceil(float64(totalRecords) / float64(rows)))

	startIndex := int64((page - 1) * rows)
	endIndex := int64(startIndex + rows)

	if startIndex < 0 {
		startIndex = 0
	}

	if endIndex > totalRecords {
		endIndex = totalRecords
	}

	pageData := dataSet[startIndex:endIndex]

	return &models.ResponsePage{
		Page:      page,
		Rows:      rows,
		TotalRows: totalRecords,
		Records:   pageData,
		Pages:     pages,
	}
}
