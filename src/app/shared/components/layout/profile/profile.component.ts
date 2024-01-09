import {Component} from '@angular/core';
import {trigger, state, transition, style, animate} from '@angular/animations';

@Component({
    selector: 'app-inline-profile',
    templateUrl: './profile.component.html',
    animations: [
        trigger('menu', [
            state('hidden', style({
                height: '0px'
            })),
            state('visible', style({
                height: '*'
            })),
            transition('visible => hidden', animate('400ms cubic-bezier(0.86, 0, 0.07, 1)')),
            transition('hidden => visible', animate('400ms cubic-bezier(0.86, 0, 0.07, 1)'))
        ])
    ]
})
export class ProfileComponent {

    active: boolean;

    onClick(event) {
        this.active = !this.active;
        event.preventDefault();
    }
}
