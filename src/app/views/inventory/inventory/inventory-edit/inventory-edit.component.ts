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
import { Participant } from '@interfaces/general/participant';
import { ParticipantService } from '@services/General/participant.service';

@Component({
  templateUrl: './inventory-edit.component.html',
})

export class InventoryEditComponent implements OnInit {
  public dataSourceFile: InventoryFile[];
  public dataSourceProduct: InventoryProduct[];  
  public resultsParticipant: Participant[];
  public formModel: FormGroup;
  public index = 0;
  public dados;
  public nomeBotao = '';
  public submitted = false;
  public salvando = false;

  constructor(private service: InventoryService,
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
    const id = Number.parseInt(this.router.paramsInheritanceStrategy.toString());
    if (id > 0) {
      this.edit(id);
    }
  }

  private createForm() {
    this.formModel = this.fb.group({
      id: null,
      name: null,
      participantId: new FormControl('', Validators.required),
      startDate: null,
      endDate: null,
      processed: new FormControl("0", Validators.required),
      cloused: new FormControl("0", Validators.required),
    });
  }

  async edit(id) {
    this.nomeBotao = 'Salvar';
    await this.service.getId(id).subscribe((result: Inventory) => {
      this.formModel.reset({
        edit: true,
        id: result.id,
        name: result.name,
        startDate: new Date(formatDate(result.startDate, 'yyyy-MM-ddTHH:mm', 'en')),
        endDate: new Date(formatDate(result.endDate, 'yyyy-MM-ddTHH:mm', 'en')),
        participantId: { id: result.id, nome: result.id + ' - ' + result.participant.name + ' ( ' + result.participant.document + ' ) ' },
        processed: result.processed,
        cloused: result.cloused,
      });
      this.formModel.controls['startDate'].disable();
      this.formModel.controls['endDate'].disable();
    });
  }

  async save() {
    this.submitted = true;
    if (this.formModel.valid && !this.salvando) {
      const model = JSON.parse(JSON.stringify(this.formModel.value))
      if (model.participantId)
        model.participantId = model.participantId.id;

      await this.service.save(model).subscribe((result) => {
        if (result != null) {
          model.edit = true;
          this.edit(result["id"]);
        } else {
          this.router.navigate(['inventory']);
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
    this.serviceParticipan.getAllSearch(0, event.query).subscribe(result => {
      result.forEach(element => { element.id = element.id , element.name = element.id + ' - ' + element.name + ' ( ' + element.document + ' ) ' });
      this.resultsParticipant = result;
    });
  }

  addParticipant(){

  }

  AddFile(){

  }

  editFile(row){}
  deleteFile(row){}
}
