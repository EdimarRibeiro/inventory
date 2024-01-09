import { Inventory } from "./inventory";

export interface InventoryProduct {
    inventoryId: number;
    productId: number;
    originCode: string;
    date: Date;
    unitId: string;
    quantity: number;
    value: number;
    valueTotal: number;
    prossessionCode: string;
    participantId: number;
    complement: string;
    accountCode: string;
    valueIr: number;
    inputQuantity: number;
    outputQuantity: number;
    balanceQuantity: number;
    inventory: Inventory;
}
 