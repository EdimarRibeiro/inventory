import { Component, OnInit, Output, Input, EventEmitter } from '@angular/core';
import { Validators, FormControl, FormGroup, FormBuilder } from '@angular/forms';
import { InventoryFile } from '@interfaces/inventory/inventory-file';
import { InventoryFileService } from '@services/inventory/inventory-file.service';
import { MessageService } from 'primeng/api';

interface UploadEvent { originalEvent: Event; files: File[]; }

@Component({
  templateUrl: './inventory-file-app.component.html',
  selector: 'app-inventory-file',
})

export class InventoryFileApp implements OnInit {
  @Output() showChange = new EventEmitter();
  @Input() show = false;
  @Input() data;

  uploadedFiles: any[] = [];

  public dataSourceTypeFile = [{ id: "xml", name: "xml" }, { id: "txt", name: "txt" }]
  public formModel: FormGroup;

  public submitted = false;
  public salvando = false;
  public allFilesFolder = false;

  constructor(private service: InventoryFileService, private fb: FormBuilder,
    private messageService: MessageService) { }

  async ngOnInit() {
    let inventoryId = null;
    let id = null;

    if (this.data['data']) {
      inventoryId = this.data['data'].inventoryId;
      id = this.data['data'].id;

      await this.createForm();

      if ((inventoryId > 0) && (id > 0)) {
        this.edit(inventoryId, id);
      } else {
        this.formModel.reset({
          edit: false,
          inventoryId: inventoryId,
          processed: false,
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
      processed: false
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
        fileType: { id: result.fileType, name: result.fileType },
        processed: result.processed
      });
    });
  }

  async save() {
    this.submitted = true;
    if (this.formModel.valid && !this.salvando) {
      this.salvando = true;
      const model = JSON.parse(JSON.stringify(this.formModel.value));
      if (model.fileType?.id) model.fileType = model.fileType.id;
      await this.service.save(model).subscribe((result) => {
        this.back();
      }, (error) => {
        this.messageService.add({ key: '001', severity: 'error', summary: 'Falha', detail: error.message });
        this.salvando = false;
      });
    }
  }

  onUpload(event) {
    for (let file of event.files) {
      const nameFile: string = file.name;
      this.service.setFileOcean(file).subscribe((result) => {
        this.formModel.controls['fileName'].setValue(result["url"]);
        this.formModel.controls['fileType'].setValue({ id: nameFile.toLowerCase().slice(-3), name: nameFile.toLowerCase().slice(-3) });
        this.messageService.add({ key: '001', severity: 'info', summary: 'Arquivo carregado!', detail: this.formModel.value.fileName });
      }, (error) => {
        this.messageService.add({ key: '001', severity: 'error', summary: 'Falha', detail: error.message });
      });
    }
  }

  onUploadAll(event: UploadEvent) {
    this.salvando = true;
    for (let file of event.files) {
      this.service.setFileOcean(file).subscribe((result) => {
        const nameFile: string = file.name;
        let inventoryFile: InventoryFile = {
          id: 0,
          inventoryId: this.formModel.controls['inventoryId'].value,
          fileName: result["url"],
          fileType: nameFile.toLowerCase().slice(-3),
          processed: false
        }
        this.service.save(inventoryFile).subscribe((result) => {
          this.uploadedFiles.push(file);
        }, error => {
          this.messageService.add({ key: '001', severity: 'error', summary: 'Falha', detail: error.message });
        });

      }, (error) => {
        this.messageService.add({ key: '001', severity: 'error', summary: 'Falha', detail: error.message });
      });
    }

    this.messageService.add({ key: '001', severity: 'info', summary: 'Importação concluída!', detail: '' });
  }

  fileAllFolder() {
    this.salvando = true;
    this.allFilesFolder = true;
  }

  onClear() {
    this.formModel.value.fileName = '';
  }
}