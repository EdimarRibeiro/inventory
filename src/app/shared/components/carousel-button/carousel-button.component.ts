import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-carousel-button',
  template: `
        <div class="p-grid">
            <div class="p-col-fixed" style="width: 40px">
                <button class="p-button-secondary" pButton type="button" icon="pi pi-angle-left" (click)="previous()" style="width: 100%"></button>
            </div>
            <div class="p-col" *ngFor="let row of optionAux; let i=index">
                <button class="font-size p-button-secondary" pButton type="button" [label]="row[optionLabel]" (click)="onClick(row)"></button>
            </div>
            <div class="p-col-fixed" style="width: 40px">
                <button class="p-button-secondary" pButton type="button" icon="pi pi-angle-right" (click)="next()" style="width: 100%"></button>
            </div>
        </div>
    `,
  styleUrls: ['carousel-button.component.scss']
})
export class CarouselButtonComponent implements OnInit {

  @Input() optionLabel: string = 'label';
  @Input() pageSize: number = 3;
  @Output() selectedChange = new EventEmitter();

  private _options = [];
  public optionAux = [];

  private page = 0;

  constructor() {
  }

  ngOnInit(): void {
  }

  @Input()
  set options(options) {
    if (options && options.length > 0) {
      this._options = options;
      this.setLimitShow();
    }
  }

  setLimitShow() {
    this.optionAux = this._options.slice(this.page * this.pageSize, (this.page + 1) * this.pageSize);
  }

  next() {
    this.page++;
    this.setLimitShow();
    if (this.optionAux.length == 0) {
      this.previous();
    }
  }

  previous() {
    this.page--;
    if (this.page < 0) {
      this.page = 0;
    }
    this.setLimitShow();
  }

  onClick(row) {
    this.selectedChange.emit(row);
  }
}
