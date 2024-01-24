import { Inventory } from "./inventory";

export interface InventoryFile {
    id: number;
    inventoryId: number;
    fileName: string;
    fileType: string;   
    processed: boolean;    
    inventory?: Inventory;
}
 