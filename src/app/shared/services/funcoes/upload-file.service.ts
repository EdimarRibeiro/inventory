import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { PermissoesStorageService } from '@auth/permissoes-storage.service';

@Injectable()
export class UploadFileService {
    private URL = environment.baseServer + 'uploadfile';
    private URLDow = environment.baseServer + 'download';
    constructor(private http: HttpClient) {
    }

    getFile(path): Observable<any> {
        const body = { "url": path }
        return this.http.post<any>(this.URLDow, body);
    }

    setFile(file): Observable<any> {
        const formData: FormData = new FormData();
        formData.append('file', file, file.name);

        const headers = new HttpHeaders();
        headers.set('Content-Type', 'multipart/form-data');

        return this.http.post<any>(this.URL, formData, { headers });
    }
}
