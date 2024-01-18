import { Component, OnInit } from "@angular/core";
import { InventoryService } from "@services/inventory/inventory.service";
import { BreadcrumbService } from "@services/layout/breadcrumb/breadcrumb.service";
import { Router } from "@angular/router";
import { FormGroup, FormBuilder } from "@angular/forms";
import { ConfirmationService, MessageService } from "primeng/api";
import { SearchDynamicService } from "../search-dynamic/search-dynamic.service";

@Component({
  templateUrl: "./inventory.component.html",
})
export class InventoryComponent implements OnInit {
  public paginar: any;
  public paginacaoConfig = {
    numeroRows: 0,
    totalPagina: 0,
    numeroPagina: 0,
  };
  public configuracaoFiltro: any = {
    tabela: "Inventory",
    filtroPadrao: { tipoCampo: "date", campo: "Inventory.StartDate", operacao: "in" },
    dataSourceCampos: [
      { id: "Inventory.StartDate", descricao: "Data", tipo: "Date" },
      { id: "Inventory.Name", descricao: "nome", tipo: "string" },
      { id: "Inventory.Id", descricao: "Código", tipo: "number" },
      { id: "Inventory.Cloused", descricao: "Fechado?", tipo: "number" },
    ],
  };
  constructor(
    private service: InventoryService,
    private breadcrumbService: BreadcrumbService,
    private router: Router,
    private fb: FormBuilder,
    private messageService: MessageService,
    private confirmationService: ConfirmationService,
    private searchDynamicDinamico: SearchDynamicService,
  ) {
    this.breadcrumbService.setItems([
      { label: "Inventários", routerLink: "/cadastro" },
    ]);
    this.router.paramsInheritanceStrategy = "emptyOnly";
  }

  ngOnInit(): void {
    this.iniciarConfiguracaoFiltro();
  }

  edit(row) {
    if (row) {
      this.router.paramsInheritanceStrategy = row.pessoaId;
      this.router.navigate(["/inventory/edit"], row.pessoaId);
    }
  }

  delete(row) {
    this.confirmationService.confirm({
      header: "Deseja realmente delete ?",
      message: row.pessoa.nome,
      icon: "pi pi-info-circle",

      accept: () => {
        this.service.delete(row).subscribe(() => {
          this.messageService.add({ key: "001", severity: "success", summary: "Excluido!", detail: row.pessoa.nome });
          this.iniciarConfiguracaoFiltro();
        },
          (err) => {
            this.messageService.add({ key: "001", life: 5000, severity: "error", summary: "Não foi possivel Excluir!", detail: "Verifique se os itens dessa tabela já foram excluídos " + err.status, });
          }
        );
      },
      reject: () => {
        this.messageService.add({
          key: "001",
          severity: "error",
          summary: "Exclusão Cancelada!",
          detail: "",
        });
      },
    });
  }

  ///////////////////////////////// Paginação //////////////////////////////////////

  iniciarConfiguracaoFiltro() {
    this.searchDynamicDinamico.loadGrid = (numeroGrid: any, pesquisa?: any) =>
      this.loadGrid(numeroGrid, pesquisa);
    this.searchDynamicDinamico.paginacaoConfig = this.paginacaoConfig;
    this.paginar = (evento, search) => this.searchDynamicDinamico.paginar(evento, search);
    setTimeout(() => {
      this.searchDynamicDinamico.filtrar();
    }, 600);
  }

  showChange(event) {
    if (event) this.paginacaoConfig.numeroRows = event.linhas;
  }

  dataSource() {
    return this.searchDynamicDinamico?.dataSource;
  }

  loadGrid(numeroGrid, pesquisa?) {
    return new Promise((resolve) => {
      this.service.getAllSearch(numeroGrid, pesquisa).subscribe((result) => {
        resolve({
          quantidadePorPagina: result["rows"]?.length ?? 0,
          dados: result["records"] ?? [],
          quantidadeDadosTotais: result["totalRows"] ?? 0,
        });
      });
    });
  }
}
