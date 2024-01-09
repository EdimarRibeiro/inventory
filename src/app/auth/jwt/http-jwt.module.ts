import { NgModule } from '@angular/core';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { HttpJwtInterceptor } from './http-jwt.interceptor';

@NgModule({
    imports: [HttpClientModule],
    providers: [
        {
            provide: HTTP_INTERCEPTORS,
            useClass: HttpJwtInterceptor,
            multi: true
        }
    ]
})
export class HttpJwtModule {
}
