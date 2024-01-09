import { Injectable } from '@angular/core';
import { User } from '@interfaces/user/user';

@Injectable()
export class PermissoesStorageService {
  private nameUser = 'user';
  
  public clear() {
    localStorage.removeItem(this.nameUser);
  }

  public getSequencial() {
    return new Promise(resolve => {

      let sequencial = Number.parseInt(localStorage.getItem('sequancial_001'), 0);
      if ((!sequencial) || (sequencial === 500) || (sequencial === 0)) {
        sequencial = 1;
      } else {
        ++sequencial;
      }
      localStorage.setItem('sequancial_001', sequencial.toString());
      resolve(sequencial);
    });
  }

  public getUser(): User {
    const user = localStorage.getItem(this.nameUser);
    if (user != 'undefined') {
      return <User>JSON.parse(user);
    } else {
      return null;
    }
  }
}
