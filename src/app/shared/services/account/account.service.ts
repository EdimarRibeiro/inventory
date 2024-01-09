import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments';
import { Observable } from 'rxjs';
import { City } from '@interfaces/general/city';

@Injectable({ providedIn: 'root' })
export class AccountService {

  private URL = environment.baseServer;

  constructor(private http: HttpClient) {
  }

  getCityAll(): Observable<City[]> {
    return this.http.get<City[]>(this.URL + '/cities');
  }

  getDocument(documento): Observable<any> {
    return this.http.get<any>(this.URL + '/document/' + documento);
  }

  getCityIbge(ibge): Observable<City> {
    return this.http.get<City>(this.URL + '/city/ibge/' + ibge);
  }

  getCep(cep): Observable<any> {
    return this.http.get<any>(this.URL + '/cep/' + cep);
  }

  save(conta) {
    return this.http.post(this.URL + '/createaccount', conta);
  }


}
