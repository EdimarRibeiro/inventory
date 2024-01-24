import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { InventoryFile } from '@interfaces/inventory/inventory-file';

@Injectable({ providedIn: 'root' })
export class InventoryFileService {

  private URL = environment.baseServer + 'inventoryfile';
  private URLs = environment.baseServer + 'inventoryfiles';
  private URLUp = environment.baseServer + 'upload';

  constructor(private http: HttpClient) {
  }

  get(id): Observable<InventoryFile[]> {
    return this.http.get<InventoryFile[]>(this.URLs + '/' + id);
  }

  getAllSearch(inventoryId: number, page: number, search: string): Observable<InventoryFile[]> {
    return this.http.get<InventoryFile[]>(`${this.URLs}/${inventoryId}/?page=${page}&search=${search}`);
  }

  getId(inventoryId, id): Observable<InventoryFile> {
    return this.http.get<InventoryFile>(this.URL + '/' + inventoryId + '/' + id);
  }

  save(inventoryFile) {
    if (inventoryFile.edit) {
      return this.http.put(this.URL + '/' + inventoryFile.inventoryId+ '/' + inventoryFile.id, inventoryFile);
    } else {
      return this.http.post(this.URL, inventoryFile);
    }
  }

  delete(inventoryFile) {
    return this.http.delete(this.URL + '/' + inventoryFile.inventoryId+ '/' + inventoryFile.id, inventoryFile);
  }

  setFileOcean(file) {
    const headers = new HttpHeaders({
      'content-type': 'multipart/form-data',
      'x-file-name': file.name,
    });
    const formData: FormData = new FormData();
    formData.append('file', file, file.name);
    return this.http.post(this.URLUp, formData, {headers});
  }
}
