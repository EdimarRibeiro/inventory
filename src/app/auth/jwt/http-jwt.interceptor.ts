import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable } from 'rxjs';
import { TokenStorageService } from '@auth/token-storage.service';

@Injectable()
export class HttpJwtInterceptor implements HttpInterceptor {
    constructor(private tokenStorageService: TokenStorageService) {
    }

    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        const token: string = this.tokenStorageService.getAccessTokenSync();
        if (token && request.url.toString().indexOf("login") === -1) {
           request = request.clone({
                setHeaders: {
                    Authorization: `${token}`
                }
            });
        }
        return next.handle(request);
    }
}
