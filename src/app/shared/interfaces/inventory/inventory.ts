import { Participant } from "@interfaces/general/participant";

export interface Inventory {
    tenantId: number;
    id: number;
    name: string;
    participantId: number;    
    startDate: Date;
    endDate?: Date;
    processed: boolean;
    cloused: boolean;
    participant: Participant;
}
 