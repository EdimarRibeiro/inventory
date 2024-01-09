import { NgModule } from '@angular/core';
import { AUTH_SERVICE, AuthModule, PROTECTED_FALLBACK_PAGE_URI, PUBLIC_FALLBACK_PAGE_URI } from 'ngx-auth';
import { TokenStorageService } from '@auth/token-storage.service';
import { AuthenticationService } from '@auth/authentication.service';
import { HelperService } from '@auth/helper/helper.service';
import { PermissoesStorageService } from '@auth/permissoes-storage.service';
import { UserStorageService } from './user-storage.service';
import { UserAuthStorageService } from './user-permissao-storage.service';

export function factory(authenticationService: AuthenticationService) {
    return authenticationService;
}

@NgModule({
    imports: [AuthModule],
    declarations: [],
    providers: [
        HelperService,
        UserStorageService,
        PermissoesStorageService,
        UserAuthStorageService,
        TokenStorageService,
        AuthenticationService,    
        {provide: PROTECTED_FALLBACK_PAGE_URI, useValue: '/'},
        {provide: PUBLIC_FALLBACK_PAGE_URI, useValue: '/login'},
        {
            provide: AUTH_SERVICE,
            deps: [AuthenticationService],
            useFactory: factory
        }
    ]
})
export class AuthenticationModule {
}
