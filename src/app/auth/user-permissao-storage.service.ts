import { Injectable } from '@angular/core';
import { User } from '@interfaces/user/user';
import { UserStorageService } from './user-storage.service';

@Injectable()
export class UserAuthStorageService {

    private nameStorage = 'user-auth';

    constructor(private userStorageService: UserStorageService) {
    }

    public setUserAuth(): UserAuthStorageService {
        let user: User = this.userStorageService.getUserLogado();
        return this;
    }

    public clear() {
        localStorage.removeItem(this.nameStorage);
    }
}
