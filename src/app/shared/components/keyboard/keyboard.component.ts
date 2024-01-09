import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
    selector: 'app-keyboard',
    template: `
        <input type="hidden" disabled="true" pInputText [(ngModel)]="displayText" style="text-align: end;">
        <table class="keyboard">
            <tbody>
            <tr>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'7'" (click)="onClickText('7')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'8'" (click)="onClickText('8')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'9'" (click)="onClickText('9')"></button>
                </td>
                <td>
                    <button pButton type="button" icon="pi pi-angle-left" (click)="onClearLast()"></button>
                </td>
            </tr>
            <tr>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'4'" (click)="onClickText('4')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'5'" (click)="onClickText('5')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'6'" (click)="onClickText('6')"></button>
                </td>
                <td rowspan="3" height="100">
                    <button [disabled]="justNumbers" pButton type="button" class="botao-concluir"  icon="pi pi-check" (click)="onCheckPressed()"></button>
                </td>
            </tr>
            <tr>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'1'" (click)="onClickText('1')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'2'" (click)="onClickText('2')"></button>
                </td>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'3'" (click)="onClickText('3')"></button>
                </td>
            </tr>
            <tr>
                <td>
                    <button [disabled]="disabledNumber" pButton type="button" [label]="'0'" (click)="onClickText('0')"></button>
                </td>
                <td>
                    <button [disabled]="display.indexOf(',')>0" pButton type="button" [label]="','" (click)="onClickText(',')"></button>
                </td>
                <td>
                    <button pButton type="button" icon="pi pi-times-circle" (click)="onClearAll()"></button>
                </td>
            </tr>
            </tbody>
        </table>

        <style>
        @supports (-moz-appearance:none) {
            .botao-concluir{
                height: 105px;
            }
        }
        </style>
    `
})
export class KeyboardComponent implements OnInit {
    @Input() disabledFunctions = false;
    @Input() disabledNumber = false;
    @Input() justNumbers = false;

    @Output() checkPressed = new EventEmitter();
    @Output() sendDisplay = new EventEmitter();
    @Output() clearLast = new EventEmitter();
    @Output() clearAll = new EventEmitter();

    public display: string[] = [];
    public displayText: string = '';

    constructor() {
    }

    ngOnInit(): void {
    }

    private setDisplay() {
        this.displayText = this.display.join('');
    }

    onClickText(event) {
        this.display.push(event);
        this.setDisplay();
        this.sendDisplay.emit(event);
    }

    onClearLast() {
        this.display.pop();
        this.setDisplay();
        this.clearLast.emit();
    }

    onClearAll() {
        this.display = [];
        this.setDisplay();
        this.clearAll.emit();
    }

    onCheckPressed() {
        if (!this.justNumbers) {
            this.checkPressed.emit(this.displayText.replace(',', '.'));
            this.display = [];
            this.setDisplay();
        }
    }
}
