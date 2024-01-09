import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthService } from 'ngx-auth';
import { TokenStorageService } from '@auth/token-storage.service';
import { environment } from '@environments';
import { User } from '@interfaces/user/user';
import { HelperService } from '@auth/helper/helper.service';
import { PermissoesStorageService } from '@auth/permissoes-storage.service';
import { Router } from '@angular/router';
import { UserStorageService } from './user-storage.service';

interface AccessData {
    token: string;
    authenticated: boolean;
    created: Date;
    expiration: Date;
    message: string;
    user: User;
}

@Injectable()
export class AuthenticationService implements AuthService {

    private URL_LOGIN = environment.baseServer + 'login';
    private URL_REFRESH = environment.baseServer + 'login/refresh';

    constructor(private http: HttpClient,
        private tokenStorageService: TokenStorageService,
        private userStorageService: UserStorageService,
        private permissoesStorageService: PermissoesStorageService,
        private helperService: HelperService,
        private router: Router) {
    }

    public isAuthorized(): Observable<boolean> {
        return this.tokenStorageService.getAccessToken().pipe(map(token => {
            return !!token && !this.helperService.isTokenExpired();
        }));
    }

    public getAccessToken(): Observable<string> {
        return this.tokenStorageService.getAccessToken();
    }

    public login({ username, password }) {
        const headers = new HttpHeaders({
            'Content-Type': 'application/json',
          });
        return this.http.post(this.URL_LOGIN, { username, password }, {headers}).pipe(map((access: AccessData) => {
            if (access.authenticated === true) {
                this.saveAccessData(access);
                this.userStorageService.setUserLogado(access.user);
            }
            return access;
        }));
       
    }

    public cryptLogin(password: string) {
        return password;
    }

    public logout(): void {
        this.tokenStorageService.clear();
        this.userStorageService.clear();
        this.permissoesStorageService.clear();
        this.router.navigate(['/login']);
    }

    private saveAccessData({ token }: AccessData) {
        this.tokenStorageService.setAccessToken(token);
    }

    refreshShouldHappen(response: HttpErrorResponse): boolean {
        return false;
    }

    refreshToken(): Observable<any> {
        return undefined;
    }

    verifyTokenRequest(url: string): boolean {
        return false;
    }

}
