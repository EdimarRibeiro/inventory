import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { User } from '@interfaces/user/user';

@Injectable()
export class UserService {

    private URL = environment.baseServer + 'user';

    constructor(private http: HttpClient) {
    }

    getAll(): Observable<User[]> {
        return this.http.get<User[]>(this.URL);
    }

    getId(id): Observable<User> {
        return this.http.get<User>(this.URL + '/' + id);
    }

    getAllGrid(page: number, search: string): Observable<User[]> {
        return this.http.get<User[]>(`${this.URL}/grid/?pagina=${page}&&pesquisa=${search}`);
    }
    
    salvar(user) {
        if (user.edit) {
            return this.http.put(this.URL + '/' + user.id, user);
        } else {
            return this.http.post(this.URL, user);
        }
    }

    excluir(user) {
        return this.http.delete(this.URL + '/' + user.id, user);
    }
}
