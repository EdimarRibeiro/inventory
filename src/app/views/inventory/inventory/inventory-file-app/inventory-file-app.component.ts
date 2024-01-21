import { Component, OnInit, Output, Input, EventEmitter } from '@angular/core';
import { Validators, FormControl, FormGroup, FormBuilder } from '@angular/forms';
import { InventoryFileService } from '@services/inventory/inventory-file.service';
import { MessageService } from 'primeng/api';

@Component({
  templateUrl: './inventory-file-app.component.html',
  selector: 'app-inventory-file',
})

export class InventoryFileApp implements OnInit {
  @Output() showChange = new EventEmitter();
  @Input() show = false;
  @Input() data;

  public dataSourceTypeFile = [{id:"xml", name:"xml"},{id:"txt", name:"txt"}]
  public formModel: FormGroup;

  public submitted = false;
  public salvando = false;

  constructor(private service: InventoryFileService, private fb: FormBuilder, 
    private messageService: MessageService) { }

  async ngOnInit() {   
    let inventoryId = null;
    let id = null;

    if (this.data['data']) {
      inventoryId = this.data['data'].inventoryId;
      id = this.data['data'].id;

      await this.createForm();

      if ((inventoryId > 0) && (id>0)) {
        this.edit(inventoryId, id);
      } else {
        this.formModel.reset({
          edit: false,
          inventoryId: inventoryId,
          processed: 0,
          fileType: 'txt'
        });
      }
    }
  }

  private createForm() {
    this.formModel = this.fb.group({
      edit: false,
      inventoryId: null,
      id: null,
      fileName: new FormControl('', Validators.required),
      fileType: new FormControl('', Validators.required),
      processed: null
    });
  }

  back() {
    this.showChange.emit(false);
  }

  async edit(inventoryId, id) {
    await this.service.getId(inventoryId, id).subscribe(result => {
      this.formModel.reset({
        edit: true,
        inventoryId: result.inventoryId,
        id: result.id,
        fileName: result.fileName,
        fileType: result.fileType,
        processed: result.processed
      });
    });
  }

  async save() {
    this.submitted = true;
    if (this.formModel.valid && !this.salvando) {
      this.salvando = true;
      await this.service.save(this.formModel.value).subscribe((result) => {
        this.back();
      });
    }
  }

  onUpload(event) {
    for (let file of event.files) {
      this.service.setFileOcean(file).subscribe((result) => {
        this.formModel.controls['fileName'].setValue(result["url"]);
        this.messageService.add({ key: '001', severity: 'info', summary: 'Arquivo carregado!', detail: this.formModel.value.fileName });
      });
    }
  }

  onClear() {
    this.formModel.value.fileName = '';
  }
}