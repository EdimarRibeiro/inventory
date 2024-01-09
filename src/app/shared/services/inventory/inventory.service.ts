import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { Inventory } from '@interfaces/inventory/inventory';

@Injectable({providedIn: 'root'})
export class InventoryService {
    private URL = environment.baseServer + 'inventory';
    private URLAll = environment.baseServer + 'inventories';

    constructor(private http: HttpClient) {
    }

    getAll(): Observable<Inventory[]> {
        return this.http.get<Inventory[]>(this.URLAll);
    }
    
    getAllSearch(page: number, search: string): Observable<Inventory[]> {
        return this.http.get<Inventory[]>(`${this.URLAll}/?page=${page}&&search=${search}`);
    }    

    getId(id: number): Observable<Inventory> {
        return this.http.get<Inventory>(this.URL + '/' + id);
    }

    getProcess(id: number): Observable<Inventory> {
        return this.http.get<Inventory>(this.URL + '/process/' + id);
    }

    getCalc(id: number): Observable<Inventory> {
        return this.http.get<Inventory>(this.URL + '/calc/' + id);
    }

    save(inventory) {

        if (inventory.edit) {
            return this.http.put(this.URL + '/' + inventory.id, inventory);
        } else {
            return this.http.post(this.URL, inventory);
        }
    }

    delete(inventory) {
        return this.http.delete(this.URL + '/' + inventory.id , inventory);
    }
}