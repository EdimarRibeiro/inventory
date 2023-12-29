package main

import (
	"fmt"
	"net/http"

	"github.com/EdimarRibeiro/inventory/api/controllers"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/gorilla/mux"
)

func main() {
	println("Initial")

	database.Initialize(false)
	userRepo := database.CreateUserRepository(database.DB)
	userLogin := controllers.CreateLogin(userRepo)
	user := controllers.CreateUserController(userRepo)

	router := mux.NewRouter()

	// Handle POST requests to /api/login with the LoginHandler function
	router.HandleFunc("/api/login", userLogin.LoginHandler).Methods("POST")

	// Private route to get all users (requires JWT)
	router.HandleFunc("/api/users", user.GetAllHandler).Methods("GET")

	// Private route to get a specific user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.GetByIdlHandler).Methods("GET")

	// Private route to update a user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.UpdateHandler).Methods("PUT")

	// Private route to delete a user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.DeleteHandler).Methods("DELETE")

	port := 8181
	fmt.Printf("Server started on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
