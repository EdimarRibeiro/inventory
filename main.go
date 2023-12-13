package main

import (
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
)

func main() {
	println("Initial")

	database.Initialize()

}
