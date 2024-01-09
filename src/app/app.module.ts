import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LOCALE_ID, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpJwtModule } from './auth/jwt/http-jwt.module';
import { AuthenticationModule } from '@auth/authentication.module';
import { registerLocaleData } from '@angular/common';
import { PoliticasComponent } from './politicas/politicas.component';
import { RippleModule } from 'primeng/ripple';
import ptBr from '@angular/common/locales/pt';

const APP_MODULES = [
  HttpJwtModule,
  AuthenticationModule
];

const PRIMENG_MODULES = [];

registerLocaleData(ptBr);

@NgModule({
  declarations: [
    AppComponent,
    PoliticasComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    RouterModule,
    APP_MODULES,
    PRIMENG_MODULES,
    RippleModule,
  ],
  providers: [{ provide: LOCALE_ID, useValue: 'pt-BR' }],
  bootstrap: [AppComponent]
})

export class AppModule {
}
