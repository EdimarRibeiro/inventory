import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { BreadcrumbComponent } from '@components/layout/breadcrumb/breadcrumb.component';
import { FooterComponent } from '@components/layout/footer/footer.component';
import { MenuComponent, SubMenuComponent } from '@components/layout/menu/menu.component';
import { ProfileComponent } from '@components/layout/profile/profile.component';
import { TopbarComponent } from '@components/layout/topbar/topbar.component';
import { CommonModule } from '@angular/common';

const APP_COMPONENTS = [
    BreadcrumbComponent,
    FooterComponent,
    MenuComponent,
    SubMenuComponent,
    ProfileComponent,
    TopbarComponent
];

@NgModule({
    declarations: [
        APP_COMPONENTS
    ],
    imports: [
        CommonModule,
        RouterModule
    ],
    exports: [APP_COMPONENTS]
})
export class LayoutModule {
}
