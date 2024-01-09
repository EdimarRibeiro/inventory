import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { LoginRoutingModule } from './login-routing.module';
import { LoginComponent } from './login.component';
import { AuthenticationService } from '@auth/authentication.service';
import { StepsModule } from 'primeng/steps';
import { AccountService } from '@services/account/account.service';

//////////////////  PrimeNG  ///////////////

import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { MessageService } from 'primeng/api';
import { DropdownModule } from 'primeng/dropdown';
import { SidebarModule } from 'primeng/sidebar';
import { PasswordModule } from 'primeng/password';
import { MessagesModule } from 'primeng/messages';
import { MessageModule } from 'primeng/message';

const APP_COMPONENT = [
  LoginComponent,
];

const PRIMENG_MODULES = [
  ButtonModule,
  PasswordModule,
  InputTextModule,
  MessagesModule,
  MessageModule,
  StepsModule,
  DialogModule,
  DropdownModule,
  SidebarModule,
];

@NgModule({
  declarations: [
    APP_COMPONENT
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    CommonModule,
    LoginRoutingModule,
    PRIMENG_MODULES
  ],
  providers: [
    AuthenticationService,
    MessageService,
    AccountService
  ]
})

export class LoginModule {
}
