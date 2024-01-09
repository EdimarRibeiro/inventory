export interface User {
    tenantId?: number;
    id: number;
    name: string;
    login: string;
    password: string;
    image: string;
    startDate: Date;
    endDate?: Date;
}
