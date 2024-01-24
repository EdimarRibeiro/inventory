import { Component, OnInit } from '@angular/core';
import { BreadcrumbService } from '@services/layout/breadcrumb/breadcrumb.service';
import { Router } from '@angular/router';
import { formatDate } from '@angular/common';
import { Validators, FormControl, FormGroup, FormBuilder } from '@angular/forms';
import { ConfirmationService, MessageService } from 'primeng/api';
import { InventoryService } from '@services/inventory/inventory.service';
import { InventoryFile } from '@interfaces/inventory/inventory-file';
import { InventoryProduct } from '@interfaces/inventory/inventory-product';
import { Inventory } from '@interfaces/inventory/inventory';
import { ParticipantService } from '@services/general/participant.service';
import { InventoryProductService } from '@services/inventory/inventory-product.service';
import { InventoryFileService } from '@services/inventory/inventory-file.service';

@Component({
  templateUrl: './inventory-edit.component.html',
})

export class InventoryEditComponent implements OnInit {
  public dataSourceFile: InventoryFile[];
  public dataSourceProduct: InventoryProduct[];
  public resultsParticipant = [];
  public formModel: FormGroup;
  public index = 0;
  public dados;
  public nomeBotao = '';
  public submitted = false;
  public salvando = false;
  public inventoryFile;
  public isInventoryFile;

  constructor(private service: InventoryService,
    private serviceFile: InventoryFileService,
    private serviceProd: InventoryProductService,
    private serviceParticipan: ParticipantService,
    private breadcrumbService: BreadcrumbService,
    private fb: FormBuilder,
    private router: Router,
    private messageService: MessageService,
    private confirmationService: ConfirmationService
  ) {

    this.breadcrumbService.setItems([
      { label: 'Inventory' },
      { label: 'Inventory', routerLink: '/inventory/cadastro' },
      { label: 'Edição' }
    ]);
  }

  ngOnInit() {
    this.nomeBotao = 'Salvar e continuar'
    this.createForm();
    const customData = window.history.state.customData;

    if (customData) {
      const id = Number.parseInt(customData.inventoryId);
      if (id > 0) {
        this.edit(id);
      }
    }
  }

  private createForm() {
    this.formModel = this.fb.group({
      id: null,
      name: null,
      participantId: new FormControl('', Validators.required),
      startDate: null,
      endDate: null,
      processed: new FormControl(false, Validators.required),
      cloused: new FormControl(false, Validators.required),
      edit: false,
    });
  }

  enableField(enabled) {
    if (enabled) {
      this.formModel.controls['startDate'].enable();
      this.formModel.controls['endDate'].enable();
    } else {
      this.formModel.controls['startDate'].disable();
      this.formModel.controls['endDate'].disable();
    }
  }

  async edit(id) {
    this.nomeBotao = 'Salvar';
    await this.service.getId(id).subscribe((result: Inventory) => {
      this.resultsParticipant = [{ id: result.id, name: result.id + ' - ' + result.participant?.name + ' ( ' + result.participant?.document + ' ) ' }];
      this.formModel.reset({
        edit: true,
        id: result.id,
        name: result.name,
        startDate: new Date(formatDate(result.startDate, 'yyyy-MM-ddTHH:mm', 'en')),
        endDate: result.endDate.getDate > new Date("0001-01-01T00:00:00Z").getDate ? new Date(formatDate(result.endDate, 'yyyy-MM-ddTHH:mm', 'en')) : null,
        participantId: { id: result.id, name: result.id + ' - ' + result.participant?.name + ' ( ' + result.participant?.document + ' ) ' },
        processed: result.processed,
        cloused: result.cloused,
      });
      this.enableField(false);
      this.LoadGrid(result.id);

    });
  }

  LoadGrid(id) {
    this.serviceFile.get(id).subscribe((result: InventoryFile[]) => {
      this.dataSourceFile = result
    });

    this.serviceProd.get(id).subscribe((result: InventoryProduct[]) => {
      this.dataSourceProduct = result
    });
  }

  async save() {
    this.submitted = true;
    if (this.formModel.valid && !this.salvando) {
      this.enableField(true);
      const model = JSON.parse(JSON.stringify(this.formModel.value));
      this.enableField(false);

      if (model.participantId)
        model.participantId = model.participantId.id;

      await this.service.save(model).subscribe((result) => {
        if (result != null) {
          model.edit = true;
          this.edit(result["id"]);
        } else {
          this.router.navigate(['/inventory/inventory']);
        }
        this.salvando = false;
      }, error => {
        this.salvando = false;
        var msg = error.error ? error.error.split(':')[1].split('.')[0] : error.statusText;
        this.messageService.add({ key: '001', severity: 'info', summary: 'Falha ao save dados!', detail: msg });
      });
    }
  }

  searchParticipant(event) {
    const search = `(lower(name) like '%25${event.query}%25' or lower(document) like '%25${event.query}%25')`; 
    this.serviceParticipan.getAllSearch(0, search).subscribe(result => {
      result["records"].forEach(element => { element.id = element.id, element.name = element.id + ' - ' + element.name + ' ( ' + element.document + ' ) ' });
      this.resultsParticipant = result["records"];
    });
  }

  addParticipant() {

  }

  AddFile(id) {
    this.inventoryFile = {
      data: { inventoryId: this.formModel.value.id, id: id },
      title: "Cadastro",
      filtro: null
    };
    this.isInventoryFile = true;
  }

  editFile(row) {
    this.AddFile(row.id)
  }

  deleteFile(row) {
    this.confirmationService.confirm({
      header: "Deseja realmente delete ?",
      message: row.fileName,
      icon: "pi pi-info-circle",

      accept: () => {
        this.serviceFile.delete(row).subscribe(() => {
          this.messageService.add({ key: "001", severity: "success", summary: "Excluido!", detail: row.fileName });
          this.LoadGrid(row.inventoryId);
        },
          (err) => {
            this.messageService.add({ key: "001", life: 5000, severity: "error", summary: "Não foi possivel Excluir!", detail: "Verifique se os itens dessa tabela já foram excluídos " + err.status, });
          }
        );
      },
      reject: () => {
        this.messageService.add({
          key: "001",
          severity: "error",
          summary: "Exclusão Cancelada!",
          detail: "",
        });
      },
    });
  }

  closeInventoryFile() {
    this.isInventoryFile = false;
    this.LoadGrid(this.formModel.value.id);
  }
}
