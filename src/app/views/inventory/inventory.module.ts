import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ComponentsModule } from '@components/components.module';
import { UploadFileService } from '@services/funcoes/upload-file.service';


//////////////////  PrimeNG  ///////////////
import { DataViewModule } from 'primeng/dataview';
import { RadioButtonModule } from 'primeng/radiobutton';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { ToastModule } from 'primeng/toast';
import { InputTextModule } from 'primeng/inputtext';
import { TableModule } from 'primeng/table';
import { TooltipModule } from 'primeng/tooltip';
import { CheckboxModule } from 'primeng/checkbox';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { ScrollPanelModule } from 'primeng/scrollpanel';
import { SelectButtonModule } from 'primeng/selectbutton';
import { CalendarModule } from 'primeng/calendar';
import { ChartModule } from 'primeng/chart';
import { TabViewModule } from 'primeng/tabview';
import { DialogModule } from 'primeng/dialog';
import { MessageModule } from 'primeng/message';
import { ConfirmationService, MessageService } from 'primeng/api';
import { DropdownModule } from 'primeng/dropdown';
import { FileUploadModule } from 'primeng/fileupload';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { InputNumberModule } from 'primeng/inputnumber';
import { SkeletonModule } from 'primeng/skeleton';
import { FiltroDinamicoService } from '@components/filtro-dinamico/filtro-dinamico.service';

/////////////  Sistema //////////////
import { InventoryService } from '@services/inventory/inventory.service';
import { InventoryEditComponent } from './inventory/inventory-edit/inventory-edit.component';
import { InventoryComponent } from './inventory/inventory.component';
import { InventoryRoutingModule } from './inventory-routing.module';

const APP_COMPONENT = [
    InventoryComponent,
    InventoryEditComponent,
];

const APP_MODULES = [
    ComponentsModule
];

const PRIMENG_MODULES = [
    ButtonModule,
    TableModule,
    InputTextModule,
    SelectButtonModule,
    ScrollPanelModule,
    DataViewModule,
    CardModule,
    CalendarModule,
    ChartModule,
    TabViewModule,
    RadioButtonModule,
    AutoCompleteModule,
    MessageModule,
    ToastModule,
    DropdownModule,
    CheckboxModule,
    TooltipModule,
    FileUploadModule,
    ConfirmDialogModule,
    InputNumberModule,
    SkeletonModule
];

@NgModule({
    declarations: [
        APP_COMPONENT
    ],
    imports: [
        FormsModule,
        ReactiveFormsModule,
        CommonModule,
        InventoryRoutingModule,
        APP_MODULES,
        PRIMENG_MODULES,
        DialogModule
    ],
    providers: [
        MessageService,
        UploadFileService,
        ConfirmationService,
        FiltroDinamicoService,        
        InventoryService,
    ]
})
export class InventoryModule {
}