import { NO_ERRORS_SCHEMA, NgModule } from '@angular/core';
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

/////////////  Sistema //////////////
import { InventoryService } from '@services/inventory/inventory.service';
import { InventoryEditComponent } from './inventory/inventory-edit/inventory-edit.component';
import { InventoryComponent } from './inventory/inventory.component';
import { InventoryRoutingModule } from './inventory-routing.module';
import { SearchDynamicService } from './search-dynamic/search-dynamic.service';
import { SearchDynamicComponent } from './search-dynamic/search-dynamic.component';

const APP_COMPONENT = [
    InventoryComponent,
    InventoryEditComponent,
    SearchDynamicComponent,
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
    //schemas: [NO_ERRORS_SCHEMA],
    declarations: [
        ...APP_COMPONENT
    ],
    imports: [
        FormsModule,
        ReactiveFormsModule,
        CommonModule,
        DialogModule,
        ...APP_MODULES,
        ...PRIMENG_MODULES,
        InventoryRoutingModule,
    ],
    providers: [
        MessageService,
        UploadFileService,
        ConfirmationService,     
        InventoryService,
        SearchDynamicService,
    ]
})
export class InventoryModule {
}