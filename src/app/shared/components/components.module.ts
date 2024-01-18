import { NgModule } from "@angular/core";
import { CommonModule, DatePipe } from "@angular/common";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { LoadingComponent } from "./loading/loading.component";

//////////////////  PrimeNG  ///////////////

import { DataViewModule } from "primeng/dataview";
import { ToastModule } from "primeng/toast";
import { InputTextModule } from "primeng/inputtext";
import { TooltipModule } from "primeng/tooltip";
import { ButtonModule } from "primeng/button";
import { DialogModule } from "primeng/dialog";
import { MessageService } from "primeng/api";
import { AccordionModule } from "primeng/accordion";
import { InputSwitchModule } from "primeng/inputswitch";
import { DropdownModule } from "primeng/dropdown";
import { InputNumberModule } from "primeng/inputnumber";
import { ChipModule } from "primeng/chip";
import { CalendarModule } from "primeng/calendar";
import { StepsModule } from "primeng/steps";
import { AutoCompleteModule } from "primeng/autocomplete";
import { ProgressSpinnerModule } from "primeng/progressspinner";
import { InputMaskModule } from "primeng/inputmask";
import { BadgeModule } from "primeng/badge";

//////////////////////// Sistema /////////////////

import { CarouselButtonComponent } from "@components/carousel-button/carousel-button.component";
import { KeyboardComponent } from "@components/keyboard/keyboard.component";
import { SearchDynamicComponent } from "../../views/inventory/search-dynamic/search-dynamic.component";
import { PermissoesStorageService } from "@auth/permissoes-storage.service";
import { SearchDynamicService } from "../../views/inventory/search-dynamic/search-dynamic.service";
import { UserStorageService } from "@auth/user-storage.service";
import { AccountService } from "@services/account/account.service";
import { AlertComponent } from "./alert/alert.component";

const APP_COMPONENT = [
  CarouselButtonComponent,
  KeyboardComponent,
  AlertComponent,
  LoadingComponent,
  //SearchDynamicComponent,
];

const APP_MODULES = [];

const PRIMENG_MODULES = [
  ButtonModule,
  DialogModule,
  AccordionModule,
  DataViewModule,
  InputSwitchModule,
  ToastModule,
  TooltipModule,
  DropdownModule,
  InputNumberModule,
  ChipModule,
  CalendarModule,
  StepsModule,
  InputMaskModule,
  BadgeModule
];

@NgModule({
  declarations: [APP_COMPONENT],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    CommonModule,
    APP_MODULES,
    PRIMENG_MODULES,
    InputTextModule,
    AutoCompleteModule,
    ProgressSpinnerModule,
  ],
  providers: [
    MessageService,
    DatePipe,
    UserStorageService,
    PermissoesStorageService,
    AccountService,
   // SearchDynamicService,
  ],
})
export class ComponentsModule { }
