
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class FiltroDinamicoService {

  constructor() { }

  public dataSource = [];
  public loadGrid: any;
  public paginacaoConfig: any;
  public events: any = { first: 0, rows: 20 };
  public searchSalvo = '';
  public numeroPagina;
  public dataSourceCarregado = false;


  filtrar(search?, page = null) {
    this.searchSalvo = search;
    this.dataSource = [];
    if (page) this.events.rows = page;
    this.paginar(this.events, search);
  }

  iniciarLoadGrid() {
    this.paginar(this.events, null);
  }
  paginar(event, search?) {
    this.paginacaoConfig.numeroPagina = 0;
    const indexAtual = event.first;
    const numeroRows = event.rows;
    const numeroPagina = this.calcularNumeroPaginacao(indexAtual, numeroRows);
    const ultimoIndex = indexAtual + numeroRows;
    if (this.searchSalvo) search = this.searchSalvo;
    this.validarJaAdicionado(indexAtual, search).then(() => {
      this.loadGrid(numeroPagina, search+`&&rows=${numeroRows}`).then((resultado: any) => {
        this.configurarPaginacao(resultado.quantidadeDadosTotais, numeroRows).then(() => {
          this.adicionarDataSource(ultimoIndex, resultado.dados, indexAtual).then(() => {
            this.dataSourceCarregado = true;
          });
        });
      });
    });
  }

  calcularNumeroPaginacao(rowAtual: number, numeroRows: number) {
    return (rowAtual + numeroRows) / numeroRows;
  }

  validarJaAdicionado(indexAtual, search) {
    return new Promise(resolve => {
      if (!this.dataSource[indexAtual] || search) resolve(true);
    })
  }

  async configurarPaginacao(quantidadeDadosTotais: number, numeroRows: number) {
    if (this.dataSource.length === 0) this.dataSource = new Array(quantidadeDadosTotais);
    if (!this.paginacaoConfig)
      this.paginacaoConfig.numeroRows = numeroRows ?? this.events.rows;
    else
      this.paginacaoConfig.numeroRows = numeroRows;
    this.paginacaoConfig.totalPagina = quantidadeDadosTotais;
  }

  async adicionarDataSource(ultimoIndex: number, dados: any, indexAtual: number) {
    let indexDadosBanco = 0;
    for (let index = indexAtual; index < ultimoIndex; index++) {
      if (dados[indexDadosBanco]) this.dataSource[index] = JSON.parse(JSON.stringify(dados[indexDadosBanco]));
      indexDadosBanco++;
    }
  }

}
