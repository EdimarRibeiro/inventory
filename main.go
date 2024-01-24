package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EdimarRibeiro/inventory/api/controllers"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-File-Name")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)

		//	http.Error(w, "Origem não permitida", http.StatusForbidden)
	})
}

func main() {
	println("Initial")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading file .env:", err)
	}

	database.Initialize(false)
	tenRepo := database.CreateTenantRepository(database.DB)
	cityRepo := database.CreateCityRepository(database.DB)
	personRepo := database.CreatePersonRepository(database.DB)
	participantRepo := database.CreateParticipantRepository(database.DB)
	userRepo := database.CreateUserRepository(database.DB)
	inventoryRepo := database.CreateInventoryRepository(database.DB)
	inventoryFileRepo := database.CreateInventoryFileRepository(database.DB)
	inventoryProductRepo := database.CreateInventoryProductRepository(database.DB)
	unitRep := database.CreateUnitRepository(database.DB)
	unitConvertRep := database.CreateUnitConvertRepository(database.DB)
	partRep := database.CreateParticipantRepository(database.DB)
	prodRep := database.CreateProductRepository(database.DB)
	docRep := database.CreateDocumentRepository(database.DB)
	docItemRep := database.CreateDocumentItemRepository(database.DB)

	user := controllers.CreateUserController(userRepo)
	city := controllers.CreateCityController(cityRepo)
	participant := controllers.CreateParticipantController(participantRepo)
	inventory := controllers.CreateInventoryController(inventoryRepo)
	inventoryFile := controllers.CreateInventoryFileController(inventoryFileRepo)
	inventoryProduct := controllers.CreateInventoryProductController(inventoryProductRepo)
	inventoryProcess := controllers.CreateInventoryProcessController(inventoryFileRepo, inventoryProductRepo, unitRep, unitConvertRep, partRep, prodRep, docRep, docItemRep)
	inventoryCalc := controllers.CreateInventoryProcessCalcController(inventoryRepo, inventoryProductRepo, docItemRep)

	fileUpload := controllers.CreateFileUploadController(tenRepo)
	createAccount := controllers.CreateAccountController(tenRepo, personRepo, userRepo, cityRepo)
	userLogin := controllers.CreateLogin(userRepo)

	router := mux.NewRouter()

	// Aplicar el middleware CORS a todas las rutas
	router.Use(corsMiddleware)

	// Handle POST requests to /api/createaccount with the CreateAccountHandler function
	router.HandleFunc("/api/createaccount", createAccount.CreateAccountHandler).Methods("OPTIONS")
	router.HandleFunc("/api/createaccount", createAccount.CreateAccountHandler).Methods("POST")

	// Handle GET requests to /api/cep with the GetCepHandler function
	router.HandleFunc("/api/cep/{cep}", createAccount.GetCepHandler).Methods("OPTIONS")
	router.HandleFunc("/api/cep/{cep}", createAccount.GetCepHandler).Methods("GET")

	// Handle GET requests to /api/document with the GetDocumentHandler function
	router.HandleFunc("/api/document/{document}", createAccount.GetDocumentHandler).Methods("OPTIONS")
	router.HandleFunc("/api/document/{document}", createAccount.GetDocumentHandler).Methods("GET")

	// Auth validate
	// Private route Handle POST requests to /api/download with the HandleFileDownload function
	router.HandleFunc("/api/download", fileUpload.HandleFileDownload).Methods("OPTIONS")
	router.HandleFunc("/api/download", fileUpload.HandleFileDownload).Methods("POST")
	// Private route Handle POST requests to /api/upload with the HandleFileUpload function
	router.HandleFunc("/api/upload", fileUpload.HandleFileUpload).Methods("OPTIONS")
	router.HandleFunc("/api/upload", fileUpload.HandleFileUpload).Methods("POST")

	/*************************************/

	// Handle POST requests to /api/login with the LoginHandler function
	router.HandleFunc("/api/login", userLogin.LoginHandler).Methods("OPTIONS")
	router.HandleFunc("/api/login", userLogin.LoginHandler).Methods("POST")

	// Private route to get all users (requires JWT)
	router.HandleFunc("/api/cities", city.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/cities", city.GetAllHandler).Methods("GET")

	// Private route to get all users (requires JWT)
	router.HandleFunc("/api/users", user.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/users", user.GetAllHandler).Methods("GET")

	// Private route to get a specific user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.GetByIdlHandler).Methods("OPTIONS")
	router.HandleFunc("/api/user/{id}", user.GetByIdlHandler).Methods("GET")

	// Private route to update a user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.UpdateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/user/{id}", user.UpdateHandler).Methods("PUT")

	// Private route to delete a user by ID (requires JWT)
	router.HandleFunc("/api/user/{id}", user.DeleteHandler).Methods("OPTIONS")
	router.HandleFunc("/api/user/{id}", user.DeleteHandler).Methods("DELETE")

	// Private route to get all participant (requires JWT)
	router.HandleFunc("/api/participants", participant.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/participants", participant.GetAllHandler).Methods("GET")

	// Private route to get a specific inventory by ID (requires JWT)
	router.HandleFunc("/api/participant/{id}", participant.GetByIdlHandler).Methods("OPTIONS")
	router.HandleFunc("/api/participant/{id}", participant.GetByIdlHandler).Methods("GET")

	// Private route to create a participant by ID (requires JWT)
	router.HandleFunc("/api/participant", participant.CreateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/participant", participant.CreateHandler).Methods("POST")

	// Private route to update a participant by ID (requires JWT)
	router.HandleFunc("/api/participant/{id}", participant.UpdateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/participant/{id}", participant.UpdateHandler).Methods("PUT")

	// Private route to delete a participant by ID (requires JWT)
	router.HandleFunc("/api/participant/{id}", participant.DeleteHandler).Methods("OPTIONS")
	router.HandleFunc("/api/participant/{id}", participant.DeleteHandler).Methods("DELETE")

	// Private route to get all inventories (requires JWT)
	router.HandleFunc("/api/inventories", inventory.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventories", inventory.GetAllHandler).Methods("GET")

	// Private route to get process inventory (requires JWT)
	router.HandleFunc("/api/inventory/process/{id}", inventoryProcess.InventaryProcessFileHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory/process/{id}", inventoryProcess.InventaryProcessFileHandler).Methods("GET")

	// Private route to get calc inventory (requires JWT)
	router.HandleFunc("/api/inventory/calc/{id}", inventoryCalc.InventaryProcessCalcHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory/calc/{id}", inventoryCalc.InventaryProcessCalcHandler).Methods("GET")

	// Private route to get a specific inventory by ID (requires JWT)
	router.HandleFunc("/api/inventory/{id}", inventory.GetByIdlHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory/{id}", inventory.GetByIdlHandler).Methods("GET")

	// Private route to create a inventory by ID (requires JWT)
	router.HandleFunc("/api/inventory", inventory.CreateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory", inventory.CreateHandler).Methods("POST")

	// Private route to update a inventory by ID (requires JWT)
	router.HandleFunc("/api/inventory/{id}", inventory.UpdateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory/{id}", inventory.UpdateHandler).Methods("PUT")

	// Private route to delete a inventory by ID (requires JWT)
	router.HandleFunc("/api/inventory/{id}", inventory.DeleteHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventory/{id}", inventory.DeleteHandler).Methods("DELETE")

	// Private route to get all inventory files (requires JWT)
	router.HandleFunc("/api/inventoryfiles/{inventoryId}", inventoryFile.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryfiles/{inventoryId}", inventoryFile.GetAllHandler).Methods("GET")

	// Private route to get a specific inventoryfile by ID (requires JWT)
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.GetByIdlHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.GetByIdlHandler).Methods("GET")

	// Private route to create a inventoryfile by ID (requires JWT)
	router.HandleFunc("/api/inventoryfile", inventoryFile.CreateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryfile", inventoryFile.CreateHandler).Methods("POST")

	// Private route to update a inventoryfile by ID (requires JWT)
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.UpdateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.UpdateHandler).Methods("PUT")

	// Private route to delete a inventoryfile by ID (requires JWT)
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.DeleteHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryfile/{inventoryId}/{id}", inventoryFile.DeleteHandler).Methods("DELETE")

	// Private route to get all inventory products (requires JWT)
	router.HandleFunc("/api/inventoryproducts/{inventoryId}", inventoryProduct.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryproducts/{inventoryId}", inventoryProduct.GetAllHandler).Methods("GET")

	// Private route to get a specific inventoryproduct by ID (requires JWT)
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.GetByIdlHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.GetByIdlHandler).Methods("GET")

	// Private route to update a inventoryproduct by ID (requires JWT)
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.UpdateHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.UpdateHandler).Methods("PUT")

	// Private route to delete a inventoryproduct by ID (requires JWT)
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.DeleteHandler).Methods("OPTIONS")
	router.HandleFunc("/api/inventoryproduct/{inventoryId}/{productId}", inventoryProduct.DeleteHandler).Methods("DELETE")

	port := 8181
	fmt.Printf("Server started on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
