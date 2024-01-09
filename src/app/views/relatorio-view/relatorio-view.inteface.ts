
export interface RelatorioView {
    modelo: string;
    title: RelatorioTitle;
    infoTitle: RelatorioSubTitulo[];
    filter: RelatorioFilter;
    fields: RelatorioFields[];
    fieldsDetail: RelatorioFields[];
    fieldsSubDetail: RelatorioFields[];
    order: RelatorioOrder;
    orientation?: "p" | "portrait" | "l" | "landscape";
    unit?: "pt" | "px" | "in" | "mm" | "cm" | "ex" | "em" | "pc";
    format?: string | number[]; //
    dataSet: any[];
    dataSetDetail: any[];
    dataSetSubDetail: string;
    observation: string;
    localDataSet: string;
}
export interface RelatorioTitle {
    name: string;
    font: string;
    style: string;
    size: number;
    color: number[];
    position: number;
    row: number;
    rowPage: number;
    pageHeight: number;
    margin: number;
}

export interface RelatorioFilter {
    name: string;
    label: string;
    type: string;
    dataSet: any[];
    keyList: string;
    localDataSet: string;
}

export interface RelatorioSubTitulo {
    name: string;
    label: string;
    font: string;
    style: string;
    styleLabel: string;
    size: number;
    color: number[];
    position: number;
    row: number;
    format: string;
    totalize: boolean;
}

export interface RelatorioFields {
    name: string;
    label: string;
    font: string;
    style: string;
    styleLabel: string;
    size: number;
    color: number[];
    position: number;
    row: number;
    format: string;
    totalize: boolean;
}
export interface RelatorioOrder {
    name: string;
    label: string;
}
