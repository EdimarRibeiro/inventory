package infrastructure

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/EdimarRibeiro/inventory/internal/internalfunc"
	"github.com/EdimarRibeiro/inventory/internal/utils"
)

type ImportFileText struct {
	InventoryFile    entitiesinterface.InventoryFileRepositoryInterface
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
	Unit             entitiesinterface.UnitRepositoryInterface
	UnitConvert      entitiesinterface.UnitConvertRepositoryInterface
	Participant      entitiesinterface.ParticipantRepositoryInterface
	Product          entitiesinterface.ProductRepositoryInterface
	Document         entitiesinterface.DocumentRepositoryInterface
	DocumentItem     entitiesinterface.DocumentItemRepositoryInterface
}

func CreateImportFileText(inventoryFile entitiesinterface.InventoryFileRepositoryInterface,
	inventoryProduct entitiesinterface.InventoryProductRepositoryInterface,
	unit entitiesinterface.UnitRepositoryInterface,
	unitConvert entitiesinterface.UnitConvertRepositoryInterface,
	participant entitiesinterface.ParticipantRepositoryInterface,
	product entitiesinterface.ProductRepositoryInterface,
	document entitiesinterface.DocumentRepositoryInterface,
	documentItem entitiesinterface.DocumentItemRepositoryInterface) *ImportFileText {
	return &ImportFileText{
		InventoryFile:    inventoryFile,
		InventoryProduct: inventoryProduct,
		Unit:             unit,
		UnitConvert:      unitConvert,
		Participant:      participant,
		Product:          product,
		Document:         document,
		DocumentItem:     documentItem,
	}
}

func (imp *ImportFileText) Execute(file *entities.InventoryFile) error {
	err := file.Validate()

	if err != nil {
		return err
	}

	if strings.ToUpper(file.FileType) != "TXT" {
		err = errors.New("invalid type file")
	}

	if err != nil {
		return err
	}

	database.Initialize(true)

	conteudo, err := internalfunc.DownloadFile(file.FileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(conteudo), "\n")
	key := ""
	doc := &entities.Document{}
	dateInventory := time.Now()

	inveRep := &database.InventoryRepository{DB: database.DB}
	prodRepo := &database.ProductRepository{DB: database.DB}
	partRepo := &database.ParticipantRepository{DB: database.DB}

	invs, err := inveRep.Search("Inventory.Id =" + strconv.FormatUint(file.InventoryId, 10))
	if err != nil {
		return err
	}
	inventario := invs[len(invs)-1]
	keys := []string{"0150", "0190", "0200", "0220", "C100", "C170", "H005", "H010"}

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			key, err = utils.CopyText(line, 1)

			if strings.LastIndex(strings.Join(keys, "|"), key) >= 0 {
				if err != nil {
					return err
				}
				if key == "0150" {
					participantId, err := partRepo.GetParticipantId(utils.CopyText(line, 2))
					if participantId == 0 || err != nil {
						part, err := entities.CreateParticipant(inventario.TenantId, line)
						if err != nil {
							return err
						}
						_, err = partRepo.Save(part)

						if err != nil {
							return err
						}
					}
				} else if key == "0190" {
					unit, err := entities.CreateUnit(inventario.TenantId, line)
					if err != nil {
						return err
					}

					unitRepo := &database.UnitRepository{DB: database.DB}
					units, err := unitRepo.Search("Id='" + unit.Id + "'")
					if len(units) == 0 || err != nil {
						_, err = unitRepo.Save(unit)

						if err != nil {
							return err
						}
					}
				} else if key == "0200" {
					prod, err := entities.CreateProduct(inventario.TenantId, line)
					if err != nil {
						return err
					}
					produtoId, err := prodRepo.GetProductId(utils.CopyText(line, 2))
					if produtoId == 0 || err != nil {
						_, err = prodRepo.Save(prod)

						if err != nil {
							return err
						}
					}
				} else if key == "0220" {
					produtoId, err := prodRepo.GetProductIdBarCode(utils.CopyText(line, 4))
					if err != nil {
						return err
					}
					conv, err := entities.CreateUnitConvert(inventario.TenantId, produtoId, line)
					if err != nil {
						return err
					}
					convRepo := &database.UnitConvertRepository{DB: database.DB}

					convs, err := convRepo.Search("ProductId=" + strconv.FormatUint(produtoId, 10) +
						" and TenantId=" + strconv.FormatUint(inventario.TenantId, 10) + " and BarCode='" + conv.BarCode + "'")
					if len(convs) > 0 || err != nil {
						_, err = convRepo.Save(conv)

						if err != nil {
							return err
						}
					}
				} else if key == "C100" {
					participantId, err := partRepo.GetParticipantId(utils.CopyText(line, 4))
					if err != nil {
						return err
					}
					doc, err = entities.CreateDocument(file.InventoryId, participantId, line)
					if err != nil {
						return err
					}
					docRepo := &database.DocumentRepository{DB: database.DB}

					docs, err := docRepo.Search("InventoryId=" + strconv.FormatUint(doc.InventoryId, 10) + " and " +
						"ParticipantId=" + strconv.FormatUint(doc.ParticipantId, 10) + " and " +
						"Number='" + doc.Number + "' and " +
						"ModelId='" + doc.ModelId + "' and " +
						"Serie='" + doc.Serie + "'")

					if len(docs) == 0 || err != nil {
						doc, err := docRepo.Save(doc)
						if err != nil {
							return err
						}
						if doc == nil {
							return errors.New("invalid save document")
						}
					} else {
						doc = &docs[0]
					}
				} else if key == "C170" {
					produtoId, err := prodRepo.GetProductId(utils.CopyText(line, 3))
					if err != nil {
						return err
					}
					item, err := entities.CreateDocumentItem(doc.Id, file.InventoryId, produtoId, line)
					if err != nil {
						return err
					}
					itemRepo := &database.DocumentItemRepository{DB: database.DB}

					itens, err := itemRepo.Search(
						"DocumentId=" + strconv.FormatUint(item.DocumentId, 10) + " and " +
							"InventoryId=" + strconv.FormatUint(item.InventoryId, 10) + " and " +
							"ProductId=" + strconv.FormatUint(item.ProductId, 10) + " and " +
							"Sequency='" + item.Sequency + "'")

					if len(itens) == 0 || err != nil {
						_, err = itemRepo.Save(item)
						if err != nil {
							return err
						}
					}
				} else if key == "H005" {
					dateInventory, err = utils.CopyTextDate(line, 2, "##-##-####")
					if err != nil {
						return err
					}
				} else if key == "H010" {
					produtoId, err := prodRepo.GetProductId(utils.CopyText(line, 2))
					if err != nil {
						return err
					}
					participantId, err := partRepo.GetParticipantId(utils.CopyText(line, 8))
					if err != nil {
						return err
					}
					inv, err := entities.CreateInventoryProduct(file.InventoryId, produtoId, participantId, dateInventory, line)
					if err != nil {
						return err
					}
					invRepo := &database.InventoryProductRepository{DB: database.DB}

					invs, err := invRepo.Search(
						"InventoryId=" + strconv.FormatUint(inv.InventoryId, 10) + " and " +
							"ProductId=" + strconv.FormatUint(inv.ProductId, 10))

					if len(invs) == 0 || err != nil {
						_, err = invRepo.Save(inv)

						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	file.SetProcessed()
	fileRepo := &database.InventoryFileRepository{DB: database.DB}
	_, err = fileRepo.Save(file)

	if err != nil {
		return err
	}
	return nil
}
