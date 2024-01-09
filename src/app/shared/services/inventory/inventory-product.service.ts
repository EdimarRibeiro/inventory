import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { InventoryProduct } from '@interfaces/inventory/inventory-product';

@Injectable()
export class InventoryProductService {
    private URL = environment.baseServer + 'inventoryproduct';
    private URLs = environment.baseServer + 'inventoryproducts';

    constructor(private http: HttpClient) {
    }

    get(inventoryId: Number): Observable<InventoryProduct[]> {
        return this.http.get<InventoryProduct[]>(`${this.URLs}/${inventoryId}`);
    }

    getId(inventoryId: Number, productId: Number): Observable<InventoryProduct> {
        return this.http.get<InventoryProduct>(`${this.URL}/${inventoryId}/${productId}`);
    }

    getAllSearch(inventoryId: number, page: number, search: string): Observable<InventoryProduct[]> {
        return this.http.get<InventoryProduct[]>(`${this.URLs}/${inventoryId}/?page=${page}&&search=${search}`);
    }

    save(inventoryProduct: any) {
        if (inventoryProduct.edit) {
            return this.http.put(this.URL + '/' + inventoryProduct.inventoryId + '/' + inventoryProduct.productId, inventoryProduct);
        } else {
            return this.http.post(this.URL, inventoryProduct);
        }
    }

    delete(inventoryProduct: any) {
        return this.http.delete(this.URL + '/' + inventoryProduct.inventoryId + '/' + inventoryProduct.productId);
    }
}
