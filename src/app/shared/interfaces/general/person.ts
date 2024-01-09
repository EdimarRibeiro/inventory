import { City } from "./city";
import { Country } from "./country";

export interface Person {
    tenantId: number;
    id: number;
    name: string;
    document: string;
    registration: string;
    countryId: number;
    country: Country;
    cityId: number;
    city: City;
    suframa: string;
    street: string;
    number: string;
    complememt: string;
    neighborhood: string;
    zipCode: string;
}
