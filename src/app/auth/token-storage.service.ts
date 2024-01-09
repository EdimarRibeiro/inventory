import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable()
export class TokenStorageService {

    private nameStorage = 'accessToken';

    public getAccessTokenSync(): string {
        return <string>localStorage.getItem(this.nameStorage);
    }

    public getAccessToken(): Observable<string> {
        const token: string = <string>localStorage.getItem(this.nameStorage);
       return of(token);
    }

    public setAccessToken(token: string): TokenStorageService {        
        localStorage.setItem(this.nameStorage, token);
        return this;
    }

    public clear() {
        localStorage.removeItem(this.nameStorage);
    }
}