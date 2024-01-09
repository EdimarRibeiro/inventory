import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { ViewsRoutingModule } from './views-routing.module';
import { ViewsComponent } from './views.component';

import { LayoutModule } from '@components/layout/layout.module';
import { CommonModule } from '@angular/common';
import { ComponentsModule } from '@components/components.module';


//////////////////  PrimeNG  ///////////////

import { ScrollPanelModule } from 'primeng/scrollpanel';


const APP_MODULES = [
    LayoutModule,
    ComponentsModule
];

const PRIMENG_MODULES = [
    ScrollPanelModule
];

@NgModule({
    declarations: [
        ViewsComponent
    ],
    imports: [
        CommonModule,
        ViewsRoutingModule,
        RouterModule,
        APP_MODULES,
        PRIMENG_MODULES
    ],
    providers: [],
    bootstrap: [ViewsComponent]
})
export class ViewsModule {
}
