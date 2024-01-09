package infrastructure

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/infrastructure/database"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
	"github.com/EdimarRibeiro/inventory/internal/internalfunc"
	"github.com/EdimarRibeiro/inventory/internal/models"
)

type ImportFileXml struct {
	InventoryFile    entitiesinterface.InventoryFileRepositoryInterface
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
	Unit             entitiesinterface.UnitRepositoryInterface
	UnitConvert      entitiesinterface.UnitConvertRepositoryInterface
	Participant      entitiesinterface.ParticipantRepositoryInterface
	Product          entitiesinterface.ProductRepositoryInterface
	Document         entitiesinterface.DocumentRepositoryInterface
	DocumentItem     entitiesinterface.DocumentItemRepositoryInterface
}

func CreateImportFileXml(inventoryFile entitiesinterface.InventoryFileRepositoryInterface,
	inventoryProduct entitiesinterface.InventoryProductRepositoryInterface,
	unit entitiesinterface.UnitRepositoryInterface,
	unitConvert entitiesinterface.UnitConvertRepositoryInterface,
	participant entitiesinterface.ParticipantRepositoryInterface,
	product entitiesinterface.ProductRepositoryInterface,
	document entitiesinterface.DocumentRepositoryInterface,
	documentItem entitiesinterface.DocumentItemRepositoryInterface) *ImportFileXml {
	return &ImportFileXml{
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

func (imp *ImportFileXml) Execute(file *entities.InventoryFile) error {
	err := file.Validate()

	if err != nil {
		return err
	}

	if strings.ToUpper(file.FileType) != "XML" {
		err = errors.New("invalid type file")
	}

	if err != nil {
		return err
	}

	database.Initialize(true)
	byteValue, err := internalfunc.DownloadFile(file.FileName)
	if err != nil {
		return err
	}
	var nFeXml models.NfeProc
	err = xml.Unmarshal(byteValue, &nFeXml)

	if err != nil {
		return err
	}
	inveRep := &database.InventoryRepository{DB: database.DB}
	prodRepo := &database.ProductRepository{DB: database.DB}
	partRepo := &database.ParticipantRepository{DB: database.DB}

	inventoryId := file.InventoryId

	invs, err := inveRep.Search("Inventory.Id=" + strconv.FormatUint(inventoryId, 10))
	if err != nil {
		return err
	}

	if len(invs) == 0 {
		err = errors.New("inventory not found")
		return err
	}
	tenantId := invs[0].TenantId

	pars, err := partRepo.Search("Id=" + strconv.FormatUint(invs[0].ParticipantId, 10))
	if err != nil {
		return err
	}

	if len(pars) == 0 {
		err = errors.New("paticipant not found")
		return err
	}

	participant := pars[0]
	operationId := nFeXml.Nfe.InfNFe.Ide.TpNF
	emitentTypeId := 0

	if nFeXml.Nfe.InfNFe.Emit.CNPJ != participant.Document {
		if (nFeXml.Nfe.InfNFe.Dest.CNPJ + nFeXml.Nfe.InfNFe.Dest.CPF) != participant.Document {
			err = errors.New("xml does not belong to this inventory")
			return err
		} else {
			emitentTypeId = 1
			operationId = 0
		}
	}

	var participantId uint64 = 0

	if emitentTypeId == 1 {
		participantId, err = partRepo.GetDocumentId(nFeXml.Nfe.InfNFe.Dest.CNPJ + nFeXml.Nfe.InfNFe.Dest.CPF)
		if participantId == 0 || err != nil {
			dest := nFeXml.Nfe.InfNFe.Dest
			part, err := entities.NewParticipant(tenantId, dest.CNPJ+dest.CPF, dest.Nome, dest.CNPJ, dest.CPF, dest.IE+dest.RG, strconv.FormatInt(int64(dest.Ender.CPais), 10), strconv.FormatInt(int64(dest.Ender.CMun), 10), dest.Ender.Lgr, dest.Ender.Nro, dest.Ender.Cpl, dest.Ender.Bairro)
			if err != nil {
				return err
			}
			part, err = partRepo.Save(part)

			if err != nil {
				return err
			}
			participantId = part.Id
		}
	} else {
		dest := nFeXml.Nfe.InfNFe.Emit
		participantId, err = partRepo.GetDocumentId(dest.CNPJ)
		if participantId == 0 || err != nil {
			part, err := entities.NewParticipant(tenantId, dest.CNPJ, dest.Nome, dest.CNPJ, "", dest.IE, strconv.FormatInt(int64(dest.Ender.CPais), 10), strconv.FormatInt(int64(dest.Ender.CMun), 10), dest.Ender.Lgr, dest.Ender.Nro, dest.Ender.Cpl, dest.Ender.Bairro)
			if err != nil {
				return err
			}
			part, err = partRepo.Save(part)

			if err != nil {
				return err
			}
			participantId = part.Id
		}
	}
	nfe := nFeXml.Nfe.InfNFe.Ide
	total := nFeXml.Nfe.InfNFe.Total.ICMSTot

	doc, err := entities.NewDocument(file.InventoryId, participantId, strconv.FormatInt(int64(operationId), 10), strconv.FormatInt(int64(emitentTypeId), 10), nfe.Mod, "00", nfe.Serie, nfe.NF, nFeXml.ProtNFe.InfProt.ChNFe, nfe.DhEmi, nfe.DhEmi, total.VNF, "", total.VDesc, 0, total.VProd, "", total.VFrete, total.VSeg, total.VOutro, total.VBC, total.VICMS, total.VBCST, total.VST, total.VIPI, total.VPIS, total.VCOFINS, 0, 0, true, "xml")
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
		doc, err = docRepo.Save(doc)
		if err != nil {
			return err
		}
	} else {
		doc = &docs[0]
	}

	for _, item := range nFeXml.Nfe.InfNFe.Det {
		produtoId, err := prodRepo.GetProductId(item.Prod.CProd, err)
		if produtoId == 0 || err != nil {
			unitRepo := &database.UnitRepository{DB: database.DB}
			unitId := item.Prod.UCom
			units, err := unitRepo.Search("Id='" + unitId + "'")
			if len(units) == 0 || err != nil {
				unit, err := entities.NewUnit(tenantId, unitId, "UNIDADE em "+unitId)
				if err != nil {
					return err
				}
				unit, err = unitRepo.Save(unit)

				if err != nil {
					return err
				}
				unitId = unit.Id
			}
			prod, err := entities.NewProduct(tenantId, item.Prod.Prod, item.Prod.EAN, "", unitId, "00", item.Prod.NCM, "", "", "", 0, item.Prod.CEST, item.Prod.CProd)
			if err != nil {
				return err
			}

			prod, err = prodRepo.Save(prod)
			if err != nil {
				return err
			}
			produtoId = prod.Id
		}

		item, err := entities.NewDocumentItem(doc.Id, inventoryId, item.Item, produtoId, "", item.Prod.QCom, item.Prod.UCom, item.Prod.VProd, 0, "", "", item.Prod.CFOP, "", 0, 0, 0, 0, 0, 0, "", "", "", 0, 0, 0, "", 0, 0, 0, 0, 0, "", 0, 0, 0, 0, 0, "", 0)
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
	}

	file.SetProcessed()
	fileRepo := &database.InventoryFileRepository{DB: database.DB}
	_, err = fileRepo.Save(file)

	if err != nil {
		return err
	}
	return nil
}
