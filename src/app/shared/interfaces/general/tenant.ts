export interface Tenant {
    id: number;
    name: string;
    document: string;
    startDate: Date;
    personId: number;
    canceled: boolean;
}
