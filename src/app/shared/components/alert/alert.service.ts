import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({providedIn: 'root'})
export class AlertService {
    public displayChange: Subject<Object> = new Subject<Object>();

    show(text: string, header: string) {
        this.displayChange.next({text: text, header: header});
    }

    hide() {
        this.displayChange.next(null);
    }
}
