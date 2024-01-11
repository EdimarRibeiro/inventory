import { formatDate } from "@angular/common";
import { Component, Output, Input, EventEmitter } from "@angular/core";
import { FormBuilder, FormGroup, Validators, FormControl } from "@angular/forms";
import { MessageService } from "primeng/api";
import { NgxIndexedDB } from "ngx-indexed-db";
import { FiltroDinamicoService } from "./filtro-dinamico.service";

@Component({
  selector: "app-filtro-dinamico",
  styleUrls: ["./filtro-dinamico.component.scss"],
  templateUrl: "./filtro-dinamico.component.html",
})
export class FiltroDinamicoComponent {
  @Output() showChange = new EventEmitter();
  @Input() configuracao: any = {};

  public formModel: FormGroup;
  public submitted = false;
  public dataSourceCondicao = [{ id: "and", descricao: "E (and)"}, { id: "or", descricao: "Ou (Or)"}];
  public dataSourceOperacao = [
    { id: "==", descricao: "Igual(=)", inativo: false },
    { id: "!=", descricao: "Diferente(<>)", inativo: false },
    { id: "<", descricao: "Menor(<)", inativo: false },
    { id: "<=", descricao: "Menor&Igual(<=)", inativo: false },
    { id: ">", descricao: "Maior(>)", inativo: false },
    { id: ">=", descricao: "Maior&Igual(>=)", inativo: false },
    { id: "contains", descricao: "Contém", inativo: false },
    { id: "startswith", descricao: "Inicia com", inativo: false },
    { id: "endswith", descricao: "Termina com", inativo: false },
    { id: "in", descricao: "Entre", inativo: false },
  ];

  public dataSourceListaPesquisas: any = [];
  public tabelaLista = [
    { table: "Inventory", id: "id" },
    { table: "Participant", id: "id" },
    { table: "User", id: "id" },
  ];

  public userId = 1;
  public userTemConfiguracao = false;
  public dataSourceLista = [];

  public db = new NgxIndexedDB("coneplusdbFiltro", 1);

  constructor(
    private fb: FormBuilder,
    private messageService: MessageService,
    private service: FiltroDinamicoService
  ) { }

  ngOnInit() {
    this.createForm();
    this.criarBancoOffline(1, this.tabelaLista).then(() => {
      this.getConfiguracaoUser();
    });
  }

  private createForm() {
    this.formModel = this.fb.group({
      campo: new FormControl('', Validators.required),
      operacao: new FormControl('', Validators.required),
      valor: new FormControl('', Validators.required),
      tipo: "string",
      condicao: new FormControl('', Validators.required),
      dataInicio: new Date(),
      dataFim: new Date(),
      paginas: 20,
    });
  }

  configuracaoPadrao(config) {
    this.configurarOperacaoAtivas(config.tipoCampo);
    this.formModel.controls["tipo"].setValue(config.filtroPadrao.tipoCampo);
    this.formModel.controls["campo"].setValue(config?.dataSourceCampos.find((a) => a?.id === config.filtroPadrao.campo));
    this.formModel.controls["operacao"].setValue(this.dataSourceOperacao.find((a) => a.id === config.filtroPadrao.operacao));
    this.formModel.controls["paginas"].setValue(this.service.events.rows);
  }

  limparFiltro() {
    this.createForm();
    this.filtrar(this.formModel);
    this.dataSourceListaPesquisas = [];
    this.excluirFiltroUser();
  }

  filtrar(formModel) {
    if (this.dataSourceListaPesquisas.length) {
      let filtro = this.dataSourceListaPesquisas[0];
      this.adicionarFiltroBanco(this.dataSourceListaPesquisas, "listaPesquisa");
      this.service.filtrar(this.filtrarPorPesquisa(), filtro.paginas);
    } else {
      this.adicionarFiltroBanco(formModel, "formulario");
      this.service.filtrar(this.formatarTipo(formModel), formModel.paginas);
    }
    this.UpdatePage(formModel);
  }

  UpdatePage(formModel){
    this.showChange.emit(formModel);
  }

  filtrarPorPesquisa() {
    let search = "";
    this.dataSourceListaPesquisas.forEach((element) => {
      search += element.id + " ";
    });
    return search;
  }

  formatarTipo(formModel: any) {
    if (!formModel.operacao?.id) formModel.operacao.id = { id: "==", descricao: "Igual(=)", inativo: false, };

    if (formModel.tipo === "string")
      return `(${formModel.campo?.id} ${formModel.operacao?.id} '${formModel.valor}')`;
    if (formModel.tipo === "number")
      return `(${formModel.campo?.id}${formModel.operacao?.id}${formModel.valor})`;

    if (formModel.tipo === "date")
      if (formModel.operacao?.id === "in") {
        return `(${formModel.campo?.id} >= '${formatDate(formModel.dataInicio, "yyyy-MM-dd", "en")} 00:00:00' and ${formModel.campo?.id} <= '${formatDate(formModel.dataFim, "yyyy-MM-dd", "en")} 23:59:59')`;
      } else {
        return `(${formModel.campo?.id}${formModel.operacao?.id}${formModel.valor})`;
      }
    if (formModel.tipo === "dateTime")
      if (formModel.operacao?.id === "in") {
        return `(${formModel.campo?.id} >= '${formatDate(formModel.dataInicio, "yyyy-MM-dd", "en")} 00:00:00' and ${formModel.campo?.id} <= '${formatDate(formModel.dataFim, "yyyy-MM-dd", "en")} 23:59:59')`;
      } else {
        return `(${formModel.campo?.id}${formModel.operacao?.id}${formModel.valor})`;
      }
    if (formModel.tipo === "lista") {
      let search = `(${formModel.campo?.id} ${formModel.operacao?.id} `;
      let valor = formModel?.campo?.idTipo === "number" ? `${formModel.valor?.id})` : `'${formModel.valor?.id}')`;
      return search + valor;
    }
  }

  campoSelecionado(event: any) {
    if (this.formModel.value.tipo !== "lista") {
      this.formModel.controls["valor"].setValue(null);
    }
    this.configurarTipo(event?.tipo);
    this.configurarOperacaoAtivas(event?.tipo);
    this.configurarOperacaoPorTipo(event?.tipo, event?.dataSource, event?.idTipo);
    if (this.formModel.value.tipo === "lista") {
      this.formModel.controls["valor"].setValue(this.dataSourceLista[0]);
    }
  }

  configurarOperacaoPorTipo(tipo, lista?, tipoLista?) {
    if (tipo === "number")
      this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[0]);
    else if (tipo === "string")
      this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[6]);
    else if (tipo === "lista") {
      if (tipoLista === "number")
        this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[0]);
      else
        this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[6]);
      this.dataSourceLista = lista;
    }
    else if (tipo === "date" || tipo === "dateTime")
      this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[9]);
    else
      this.formModel.controls["operacao"].setValue(this.dataSourceOperacao[6]);
  }

  configurarTipo(tipo) {
    if (tipo) this.formModel.controls["tipo"].setValue(tipo);
    else this.formModel.controls["tipo"].setValue(null);
  }

  configurarOperacaoAtivas(tipo) {
    if (tipo === "date" || tipo === "dateTime") {
      this.dataSourceOperacao.forEach((operacao) => {
        if (operacao.descricao !== "Entre") operacao.inativo = true;
        else this.formModel.controls["operacao"].setValue(operacao);
      });
    } else {
      this.dataSourceOperacao.forEach((operacao) => {
        operacao.inativo = false;
      });
      this.formModel.controls["operacao"].setValue(null);
    }
  }

  //////////////////////////////////////  Lista Pesquisa  //////////////////////////
  adicionarListaPesquisa() {
    if (this.validarAdicionarListaPesquisa()) {
      this.adicionarDataSourcePesquisa(Math.random().toString(36).substr(2, 9));
      this.createForm();
    }
  }

  adicionarDataSourcePesquisa(idExcluir) {
    (this.formModel.value.tipo === "data" || this.formModel.value.tipo === "dateTime")
      ? this.adicionarDataSourcePesquisaData(idExcluir)
      : this.adicionarDataSourcePesquisaValor(idExcluir);
  }

  adicionarDataSourcePesquisaData(idExcluir) {
    if (this.formModel.value.condicao) {
      this.dataSourceListaPesquisas.push({
        descricao: this.formModel.value.condicao.descricao,
        id: this.formModel.value.condicao ? this.formModel.value.condicao.id : null,
        tipo: "condicao",
        paginas: this.formModel.value.paginas,
        operacao: this.formModel.value.operacao,
        idExcluir: idExcluir,
      });
    }

    this.dataSourceListaPesquisas.push({
      descricao: `De ${formatDate(this.formModel.value.dataInicio, "dd/MM/yyyy", "en")} até ${formatDate(this.formModel.value.dataFim, "dd/MM/yyyy ", "en")} `,
      id: this.formatarTipo(this.formModel.value),
      tipo: "operacao",
      paginas: this.formModel.value.paginas,
      operacao: this.formModel.value.operacao,
      idExcluir: idExcluir,
    });
  }

  adicionarDataSourcePesquisaValor(idExcluir) {
    if (this.formModel.value.condicao) {
      this.dataSourceListaPesquisas.push({
        descricao: this.formModel.value.condicao.descricao,
        id: this.formModel.value.condicao ? this.formModel.value.condicao.id : null,
        tipo: "condicao",
        paginas: this.formModel.value.paginas,
        operacao: this.formModel.value.operacao,
        idExcluir: idExcluir,
      });
    }

    this.dataSourceListaPesquisas.push({
      descricao: this.formModel.value.campo.descricao + " " + this.formModel.value.operacao.descricao + " " + (this.formModel.value.valor?.id ? this.formModel.value.valor.descricao : this.formModel.value.valor),
      id: this.formatarTipo(this.formModel.value),
      tipo: "operacao",
      paginas: this.formModel.value.paginas,
      operacao: this.formModel.value.operacao,
      idExcluir: idExcluir,
    });
  }

  ////////////// Validações //////////////
  validarAdicionarListaPesquisa() {
    if (this.validarFormModel()) return true;
    else
      this.messageService.add({
        key: "001",
        severity: "error",
        summary: "Preencha os campos corretamente!",
        detail: "",
      });
  }

  validarFormModel() {
    return (this.formModel.value.campo && this.formModel.value.operacao && (this.formModel.value.valor || (this.formModel.value.dataInicio && this.formModel.value.dataFim))
    );
  }

  ////////////// Remover //////////////
  removerDaListaPesquisa(item) {
    if (this.dataSourceListaPesquisas.length > 1 && this.dataSourceListaPesquisas[0] === item) {
      this.dataSourceListaPesquisas = this.dataSourceListaPesquisas.filter((a) => a !== this.dataSourceListaPesquisas[1]);
    }
    this.dataSourceListaPesquisas = this.dataSourceListaPesquisas.filter((a) => item.idExcluir !== a.idExcluir);
    if (this.dataSourceListaPesquisas.length === 0) this.limparFiltro();
  }

  //////////////////////////////////////  Banco de dados  //////////////////////////
  criarBancoOffline(versao: number, tabelas: any[]) {
    return new Promise((resolve) => {
      this.db.openDatabase(versao, (evt: {
        currentTarget: {
          result: {
            createObjectStore: (
              arg0: string,
              arg1: { keyPath: string; autoIncrement: boolean }
            ) => any;
          };
        };
      }) => {
        this.criarTabelasNaEstrutura(tabelas, evt);
      }
      ).then(() => {
        resolve(true);
      });
    });
  }

  criarTabelasNaEstrutura(tabelas: any, evt: any) {
    tabelas.forEach((tabela: any) => {
      evt.currentTarget.result.createObjectStore(tabela.table, {
        keyPath: tabela.id,
        autoIncrement: false,
      });
    });
  }

  getConfiguracaoUser() {
    this.configuracaoPadrao(this.configuracao);
    this.db.getByKey(this.configuracao?.tabela, this.userId).then((filtroUser) => {
      if (filtroUser) {

        filtroUser.dataSourceLista ? this.loadListaPesquisa(filtroUser.dataSourceLista) : this.loadFiltroUser(filtroUser.formulario);
      }
    }), error => {
      this.configuracaoPadrao(this.configuracao);
    };
  }

  loadFiltroUser(filtroUser) {
    this.userTemConfiguracao = true;
    this.configurarOperacaoAtivas(filtroUser.tipo);
    this.definirFiltroUser(filtroUser);
  }

  loadListaPesquisa(listPesquisa) {
    this.dataSourceListaPesquisas = listPesquisa;
  }

  adicionarFiltroBanco(filtro, tipoFiltro) {
    let salvarFiltro: any = { id: this.userId };
    if (tipoFiltro === "listaPesquisa") salvarFiltro.dataSourceLista = filtro;
    if (tipoFiltro === "formulario") salvarFiltro.formulario = filtro;
    this.excluirFiltroUser();
    this.db.add(this.configuracao?.tabela, salvarFiltro);
    this.submitted = true;
  }

  definirFiltroUser(filtroSalvo) {
    this.configurarOperacaoAtivas(filtroSalvo.tipo);
    this.formModel.controls["tipo"].setValue(filtroSalvo.tipo);
    this.formModel.controls["campo"].setValue(filtroSalvo.campo);
    this.formModel.controls["operacao"].setValue(filtroSalvo.operacao);
    this.formModel.controls["condicao"].setValue(filtroSalvo.condicao);
    this.formModel.controls["dataInicio"].setValue(filtroSalvo.dataInicio);
    this.formModel.controls["dataFim"].setValue(filtroSalvo.dataFim);
    this.formModel.controls["paginas"].setValue(filtroSalvo.paginas ?? this.formModel.value.paginas);

    if (filtroSalvo?.campo?.dataSource) {
      this.dataSourceLista = filtroSalvo?.campo?.dataSource;
      this.formModel.controls["valor"].setValue(filtroSalvo.valor);
    }
  }

  excluirFiltroUser() {
    this.db.delete(this.configuracao?.tabela, this.userId);
  }
}
