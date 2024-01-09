import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RelatorioViewComponent } from './relatorio-view.component';

const APP_MODULES = [
];

const PRIMENG_MODULES = [
];

@NgModule({
    declarations: [
    ],
    imports: [
        CommonModule,
        RouterModule,
        APP_MODULES,
        PRIMENG_MODULES
    ],
    providers: [
        RelatorioViewComponent],
    bootstrap: []
})
export class RelatorioViewModule {
}
