import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProtectedGuard } from 'ngx-auth';
import { InventoryEditComponent } from './inventory/inventory-edit/inventory-edit.component';
import { InventoryComponent } from './inventory/inventory.component';

const routes: Routes = [

    { path: 'inventory', component: InventoryComponent, canActivate: [ProtectedGuard] },
    { path: 'inventory/edit', component: InventoryEditComponent, canActivate: [ProtectedGuard] },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class InventoryRoutingModule {
}