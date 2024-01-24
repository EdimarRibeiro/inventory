import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '@environments';
import { Participant } from '@interfaces/general/participant';
import { Observable } from 'rxjs';

@Injectable()
export class ParticipantService {
  private URL = environment.baseServer + 'participant';

  constructor(private http: HttpClient) {
  }

  getAll(): Observable<Participant[]> {
    return this.http.get<Participant[]>(this.URL);
  }

  getId(id: number): Observable<Participant> {
    return this.http.get<Participant>(this.URL + '/' + id);
  }

  getAllSearch(page: number, search: string): Observable<Participant[]> {
    return this.http.get<Participant[]>(`${this.URL}s?page=${page}&search=${search}`);
  }

  save(participant: any) {
    if (participant.edit) {
      return this.http.put(this.URL + '/' + participant.id, participant);
    } else {
      return this.http.post(this.URL, participant);
    }
  }

  delete(participant: any) {
    return this.http.delete(`${this.URL}/${participant.id}`);
  }

}
