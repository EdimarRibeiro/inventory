package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/dgrijalva/jwt-go"
)

type loginControler struct {
	user entitiesinterface.UserRepositoryInterface
}

func CreateLogin(userRep entitiesinterface.UserRepositoryInterface) *loginControler {
	return &loginControler{user: userRep}
}

func (login *loginControler) GetUserLogin(credentials models.Credentials) (*entities.User, error) {
	users, err := login.user.Search("Login ='" + credentials.Username + "' and EndDate is null")
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user login not autorized")
	}

	return &users[0], nil
}

func AuthenticateUser(login *loginControler, credentials models.Credentials) (*bool, *uint64, error) {
	user, err := login.GetUserLogin(credentials)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, errors.New("user login not autorized")
	}

	if user.Password != credentials.Password {
		return nil, nil, errors.New("invalid password")
	}

	auth := user.Password == credentials.Password
	return &auth, &user.TenantId, nil
}

func CreateToken(username string, tenantId uint64) (string, error) {
	claims := models.CustomClaims{
		Username:   username,
		ExternalId: tenantId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(models.SecretKey)
}

func (login *loginControler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	auth, tenantId, err := AuthenticateUser(login, credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if *auth {
		token, err := CreateToken(credentials.Username, *tenantId)
		if err != nil {
			http.Error(w, "Error generating token "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Send the token in the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		// Authentication failed
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
