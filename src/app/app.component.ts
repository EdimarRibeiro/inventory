import { Component } from "@angular/core";
import { PrimeNGConfig } from "primeng/api";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
})
export class AppComponent {
  constructor(private primengConfig: PrimeNGConfig) { }

  ngOnInit() {
    this.primengConfig.ripple = true;
    this.primengConfig.setTranslation({
      dayNames: [
        "Domingo",
        "Segunda-feira",
        "Terça-feira",
        "Quarta-feira",
        "Quinta-feira",
        "Sexta-feira",
        "Sábado",
      ],
      dayNamesShort: ["Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"],
      dayNamesMin: ["Dom", "2ª", "3ª", "4ª", "5ª", "6ª", "Sáb"],
      monthNames: [
        "Janeiro",
        "Fevereiro",
        "Março",
        "Abril",
        "Maio",
        "Junho",
        "Julho",
        "Agosto",
        "Setembro",
        "Outubro",
        "Novembro",
        "Dezembro",
      ],
      monthNamesShort: [
        "Jan",
        "Fev",
        "Mar",
        "Abr",
        "Mai",
        "Jun",
        "Jul",
        "Ago",
        "Set",
        "Out",
        "Nov",
        "Dez",
      ],
      "clear": "Limpar",
      "apply": "Aplicar",
      "matchAll": "Corresponder a todos",
      "matchAny": "Corresponder a qualquer",
      "addRule": "Adicionar regra",
      "removeRule": "Remover regra",
      "startsWith": "Começa com",
      "contains": "Contém",
      "notContains": "Não contém",
      "endsWith": "Termina com",
      "equals": "É igual a",
      "notEquals": "Não é igual",
      "noFilter": "Sem filtro",
      "is": "é",
      "isNot": "não é",
      "before": "Antes de",
      "after": "Depois de",
      "dateIs": "Data é",
      "dateIsNot": "Data não é",
      "dateBefore": "Data é anterior",
      "dateAfter": "Data é depois",
      "lt": "Menor que",
      "lte": "Menor ou igual a",
      "gt": "Maior que",
      "gte": "Maior ou igual a",
    });
  }
}
