import { Component, Input, OnDestroy, OnInit } from "@angular/core";
@Component({
  templateUrl: "./loading.component.html",
  styleUrls: ["./loading.component.scss"],
  selector: "app-loading",
})
export class LoadingComponent implements OnInit {
  constructor() {}

  @Input()
  mensagem: string = "";
  @Input()
  carregando: boolean = false;
  @Input()
  modal: boolean = false;

  ngOnInit() {}
}
