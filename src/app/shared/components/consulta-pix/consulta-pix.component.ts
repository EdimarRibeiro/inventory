import { Component, EventEmitter, Input, OnInit, Output } from "@angular/core";
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from "@angular/forms";
import { Impressao, Relatorio } from "@interfaces/comercial/impressao-direta";
import { QRCodePix } from "@interfaces/diversos/qrcode-pix";
import { MessageService } from "primeng/api";
import { jsPDF } from "jspdf";
import { delay } from "rxjs/operators";
import { PedidoService } from "@services/comercial/pedido.service";
import { PixService } from "@services/diversos/pix.service";
import { Pedido } from "@interfaces/comercial/pedido";
import { EmpresaFilialConfiguracaoService } from "@services/inventory/empresa-filial-configuracao.service";
import { FormaPagService } from "@services/diversos/forma-pag.service";
import { TransacaoTefService } from "@services/funcoes/transacao-tef.service";
import { TransacaoCartao } from "@interfaces/transacaocartao/transacaocartao";
import { DateValidator } from "app/utilities/Validator";
import { formatDate } from "@angular/common";

@Component({
  selector: "app-consulta-pix",
  templateUrl: "./consulta-pix.component.html",
  styleUrls: ["./consulta-pix.component.scss"],
  preserveWhitespaces: true,
})
export class ConsultaPixComponent implements OnInit {

  @Input() txid: string;
  @Input() tipo: string;
  @Output() voltar = new EventEmitter();
  public objetoConsulta = null;
  public visivel = false;

  constructor(
    private fb: FormBuilder,
    private messageService: MessageService,
    private servicePix: PixService,
  ) {
  }

  ngOnInit() {
    this.visivel = true;
    if (this.tipo == 'Imediato') {
      this.consulta(this.txid, 4);
    }
    else if (this.tipo == 'Prazo') {
      this.consultaVencimento(this.txid, 4);
    }
    else {
      this.messageService.add({
        key: "001",
        severity: "error",
        summary: "Erro ao consultar transação do PIX",
        life: 10000,
        detail: "Verifique os campos necessários para consultar o PIX.",
      });
    }
  }

  fechar() {
    this.visivel = false;
    this.voltar.emit(false);
  }

  consulta(txid: string, banco: number) {
    this.servicePix.getPixConsultaTransacao(txid, banco).subscribe(
      (result) => {
        this.objetoConsulta = result;
      },
      (err) => {
        this.messageService.add({
          key: "001",
          severity: "error",
          summary: "Erro ao consultar transação do PIX",
          life: 10000,
          detail: "Verifique os campos necessários para consultar o PIX.",
        });
        this.fechar();
      }
    );
  }

  consultaVencimento(txid: string, banco: number) {
    this.servicePix.getPixConsultaTransacaoVencimento(txid, banco).subscribe(
      (result) => {
        this.objetoConsulta = result;
      },
      (err) => {
        this.messageService.add({
          key: "001",
          severity: "error",
          summary: "Erro ao consultar transação do PIX",
          life: 10000,
          detail: "Verifique os campos necessários para consultar o PIX.",
        });
        this.fechar();
      }
    );
  }
}
