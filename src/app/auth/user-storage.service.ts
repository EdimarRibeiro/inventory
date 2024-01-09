import { Injectable } from '@angular/core';
import { User } from '@interfaces/user/user';

@Injectable()
export class UserStorageService {

    private nameStorage = 'user';

    public getUserLogado(): User {
        return <User>JSON.parse(localStorage.getItem(this.nameStorage));
    }

    public setUserLogado(user: User): UserStorageService {
        localStorage.setItem(this.nameStorage, JSON.stringify(user));
        return this;
    }

    public clear() {
        localStorage.removeItem(this.nameStorage);
    }
}
