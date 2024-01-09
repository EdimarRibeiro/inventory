import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthenticationService } from '../auth/authentication.service';
import { Router } from '@angular/router';
import { MenuItem, Message, MessageService } from 'primeng/api';
import { AccountService } from '@services/account/account.service';

@Component({
  templateUrl: './login.component.html'
})

export class LoginComponent implements OnInit {
  @Output() showChange = new EventEmitter();

  public modelForm: FormGroup;
  public modelFormAccount: FormGroup;
  public errors: Message[];
  public dataSourceCity = [];
  public registroAccount = false;
  items: MenuItem[];
  public activeIndex = 0;

  constructor(private fb: FormBuilder, private authService: AuthenticationService, private router: Router,
    private accountService: AccountService, private messageService: MessageService) {
    this.createForm();
  }

  ngOnInit() {
    this.modelForm.reset({
      username: '',
      password: ''
    });

    this.items = [{
      label: 'Informe um CNPJ',
      command: (event: any) => {
      }
    },
    {
      label: 'Crie Login e Senha',
      command: (event: any) => {
      }
    },
    {
      label: 'Complente seu Cadastro',
      command: (event: any) => {
      }
    },
    {
      label: 'Confirme!',
      command: (event: any) => {
      }
    }
    ];
  }

  private createForm() {
    this.modelForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });

    this.modelFormAccount = this.fb.group({
      email: ['', Validators.required],
      password: ['', Validators.required],
      fantasia: ['', Validators.required],
      document: ['', Validators.required],
      name: ['', Validators.required],
      registration: null,
      AccountAddress: {
        zipCode: null,
        street: null,
        number: null,
        complememt: null,
        neighborhood: null,
        countryId: ['', Validators.required],
        cityId: ['', Validators.required]
      }
    });
  }

  public login() {
    this.errors = [];
    this.authService.login(this.modelForm.value).subscribe((result) => {
      if (result.authenticated === true) {
        this.router.navigate(['']);
      } else {
        this.errors = [{ severity: 'error', summary: 'Falha na autenticação', detail: result.message }];
      }
    }, error => {
      this.errors = [{ severity: 'error', summary: 'Falha na autenticação', detail: error.message }];
    });
  }

  criarAccount() {
    this.activeIndex = 0;
    this.registroAccount = true;
  }

  back() {
    this.registroAccount = false;
    this.showChange.emit(false);
    this.errors = [];
  }

  verificarDocumento(documento: string) {
    this.errors = [];
    documento = documento.trim().replace('.', '').replace('.', '').replace('-', '').replace('/', '');
    this.accountService.getDocument(documento).subscribe(cad => {
      if (documento.length <= 11) {
        this.errors = [{ severity: 'error', summary: 'Verifique o n° do documento -', detail: cad.fantasia + (documento.length === 11 ? ', no momento só estamos cadastrando empresas!' : '') }];
      } else {
        this.activeIndex += 1;
        this.accountService.getCityAll().subscribe(result => {
          const estado = result.reduce((current, next) => {
            current.push({ id: next.id, description: next.description });
            return current;
          }, []);

          this.modelFormAccount.reset({
            username: this.modelFormAccount.value.email,
            password: this.modelFormAccount.value.senha,
            fantasia: cad.fantasia,
            document: cad.document,
            name: cad.name,
            registration: cad.registration,
            AccountAddress: {
              zipCode: cad.accountAddress.zipCode,
              street: cad.accountAddress.street,
              number: cad.accountAddress.number,
              complememt: cad.accountAddress.component,
              neighborhood: cad.accountAddress.neighborhood,
              countryId: cad.accountAddress.countryId,
              cityId: cad.accountAddress.cityId
            }
          });
        });
      }
    });
  }

  isEmail(email: string): boolean {
    var serchfind: boolean;
    var regexp = new RegExp(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/);
    serchfind = regexp.test(email);
    return serchfind
  }

  save() {
    this.errors = [];
    const account = JSON.parse(JSON.stringify(this.modelFormAccount.value));

    this.accountService.save(account).subscribe(dados => {
      if (dados['create'] === false) {
        this.errors = [{ severity: 'error', summary: 'Ops ocorreu uma falha -', detail: dados['message'] }];
      } else {
        this.activeIndex += 1;
        this.modelForm.reset({
          username: this.modelFormAccount.value.email,
          password: this.modelFormAccount.value.senha
        });
      }
    }, error => {
      this.errors = [{ severity: 'error', summary: 'Ops ocorreu uma falha -', detail: error.message }];
    });
  }

  confirmarDados() {
    this.save();
  }
}
