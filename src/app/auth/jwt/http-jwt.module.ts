import { NgModule } from '@angular/core';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { HttpJwtInterceptor } from './http-jwt.interceptor';
import { JwtHelperService } from '@auth0/angular-jwt';

@NgModule({
    imports: [HttpClientModule],
    providers: [
        JwtHelperService,
        {
            provide: HTTP_INTERCEPTORS,
            useClass: HttpJwtInterceptor,
            multi: true
        }
    ]
})
export class HttpJwtModule {
}
