import { Component, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AlertService } from '@components/alert/alert.service';

@Component({
    selector: 'app-alert',
    template: `
        <p-dialog [header]="header" [modal]="true" [(visible)]="display" [resizable]="false" [draggable]="false" [closable]="false" [style]="{'width': '400px'}">
            <div class="p-grid">
                <div class="p-col-12">
                    {{text}}
                </div>
            </div>
            <p-footer>
                <button pButton type="button" [label]="'Ok'" (click)="display=false"></button>
            </p-footer>
        </p-dialog>`
})
export class AlertComponent implements OnInit {
    public header = '';
    public display = false;
    public text = '';
    private subscription: Subscription;

    constructor(private service: AlertService) {
    }

    ngOnInit() {
        this.subscribeToNotifications();
    }

    subscribeToNotifications() {
        this.subscription = this.service.displayChange
            .subscribe((res: any) => {
                if (res != null) {
                    this.display = true;
                    this.text = res.text;
                    this.header = res.header;
                } else {
                    this.display = false;
                }
            });
    }
}
