import { Component } from '@angular/core';
import { UserStorageService } from '@auth/user-storage.service';
import { ViewsComponent } from '@views/views.component';

@Component({
    selector: 'app-topbar',
    templateUrl: './topbar.component.html'
})
export class TopbarComponent {
    public userLogado;
    public sistemaVersao = '';
    darkDemoStyle: HTMLStyleElement;

    constructor(public app: ViewsComponent,
        private userStorageService: UserStorageService,
    ) {
        this.userLogado = userStorageService.getUserLogado();
        this.versaoSistema();
    }

    changeTheme(event: Event, theme: string, dark: boolean) {
        let themeLink: HTMLLinkElement = <HTMLLinkElement>document.getElementById('theme-css');
        themeLink.href = 'assets/themes/' + theme + '/theme-accent.css';

        if (dark) {
            if (!this.darkDemoStyle) {
                this.darkDemoStyle = document.createElement('style');
                this.darkDemoStyle.type = 'text/css';
                this.darkDemoStyle.innerHTML = '.implementation { background-color: #3f3f3f; color: #dedede} .implementation > h3, .implementation > h4{ color: #dedede}';
                document.body.appendChild(this.darkDemoStyle);
            }
        } else if (this.darkDemoStyle) {
            document.body.removeChild(this.darkDemoStyle);
            this.darkDemoStyle = null;
        }

        event.preventDefault();
    }


    limparCache() {
        window.history.forward();
        window.location.reload();
    }

    versaoSistema() {
        this.sistemaVersao ="1.00.0"
    }
}
