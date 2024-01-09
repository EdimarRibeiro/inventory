import { animate, state, style, transition, trigger } from "@angular/animations";
import { AfterViewInit, Component, Input, OnInit } from "@angular/core";
import { ViewsComponent } from "@views/views.component";
import { MenuItem } from "primeng/api";

@Component({
  selector: "app-menu",
  template: `
    <ul
      app-submenu
      [item]="app.grouped ? modelGrouped : modelUngrouped"
      root="true"
      class="layout-menu"
      visible="true"
      [reset]="reset"
      parentActive="true"
    ></ul>
  `,
})
export class MenuComponent implements OnInit, AfterViewInit {
  @Input() reset: boolean;

  modelGrouped: any[];

  modelUngrouped: any[];

  constructor(public app: ViewsComponent) { }

  ngOnInit() {
    this.modelGrouped = [
      { label: "Dashboard", icon: "pi pi-fw pi-home", routerLink: ["/"] },
      {
        label: "Inventário",
        icon: "pi pi-fw pi-circle-on",
        items: [
          {
            label: "Inventário",
            icon: "pi pi-fw pi-circle-on",
            routerLink: ["/inventory/inventory"],
          },
        ],
      },
    ];

    this.modelUngrouped = [
      {
        label: "Menu",
        icon: "pi pi-fw pi-circle-on",
        items: this.modelGrouped,
      },
    ];
  }

  ngAfterViewInit() {
    setTimeout(() => {
      this.app.layoutMenuScrollerViewChild.moveBar();
    }, 100);
  }

  changeTheme(theme: string, scheme: string) {
    const layoutLink: HTMLLinkElement = document.getElementById(
      "layout-css"
    ) as HTMLLinkElement;
    layoutLink.href = "assets/layout/css/layout-" + theme + ".css";

    const themeLink: HTMLLinkElement = document.getElementById(
      "theme-css"
    ) as HTMLLinkElement;
    themeLink.href = "assets/theme/" + theme + "/theme-" + scheme + ".css";

    const topbarLogo: HTMLImageElement = document.getElementById(
      "layout-topbar-logo"
    ) as HTMLImageElement;

    const menuLogo: HTMLImageElement = document.getElementById(
      "layout-menu-logo"
    ) as HTMLImageElement;

    if (theme === "yellow" || theme === "lime") {
      topbarLogo.src = "assets/layout/images/logo-branca.jpg";
      menuLogo.src = "assets/layout/images/logo-branca.jpg";
    } else {
      topbarLogo.src = "assets/layout/images/logo-branca.jpg";
      menuLogo.src = "assets/layout/images/logo-branca.jpg";
    }

    if (scheme === "dark") {
      this.app.darkMenu = true;
    } else if (scheme === "light") {
      this.app.darkMenu = false;
    }
  }
}

@Component({
  /* tslint:disable:component-selector */
  selector: "[app-submenu]",
  /* tslint:enable:component-selector */
  template: `
    <ng-template
      ngFor
      let-child
      let-i="index"
      [ngForOf]="root ? item : item.items"
    >
      <li
        [ngClass]="{
          'layout-root-menuitem': root,
          'active-menuitem': isActive(i)
        }"
        [class]="child.badgeStyleClass"
        *ngIf="child.visible === false ? false : true"
      >
        <div *ngIf="root">
          <span class="layout-menuitem-text">{{ child.label }}</span>
        </div>
        <a
          [href]="child.url || '#'"
          (click)="itemClick($event, child, i)"
          (mouseenter)="onMouseEnter(i)"
          *ngIf="!child.routerLink"
          [class]="child.styleClass"
          [attr.tabindex]="!visible ? '-1' : null"
          [attr.target]="child.target"
        >
          <i class="layout-menuitem-icon" [ngClass]="child.icon"></i>
          <span class="layout-menuitem-text">{{ child.label }}</span>
          <i
            class="pi pi-fw pi-angle-down layout-submenu-toggler"
            *ngIf="child.items"
          ></i>
          <span class="menuitem-badge" *ngIf="child.badge">{{
            child.badge
          }}</span>
        </a>

        <a
          (click)="itemClick($event, child, i)"
          (mouseenter)="onMouseEnter(i)"
          *ngIf="child.routerLink"
          [routerLink]="child.routerLink"
          routerLinkActive="active-route"
          [fragment]="child.fragment"
          [routerLinkActiveOptions]="{ exact: true }"
          [attr.tabindex]="!visible ? '-1' : null"
          [attr.target]="child.target"
        >
          <i class="layout-menuitem-icon" [ngClass]="child.icon"></i>
          <span class="layout-menuitem-text">{{ child.label }}</span>
          <i
            class="pi pi-fw pi-angle-down layout-menuitem-toggler"
            *ngIf="child.items"
          ></i>
          <span class="menuitem-badge" *ngIf="child.badge">{{
            child.badge
          }}</span>
        </a>
        <div class="layout-menu-tooltip">
          <div class="layout-menu-tooltip-arrow"></div>
          <div class="layout-menu-tooltip-text">{{ child.label }}</div>
        </div>
        <ul
          app-submenu
          [item]="child"
          *ngIf="child.items"
          [visible]="isActive(i)"
          [reset]="reset"
          [parentActive]="isActive(i)"
          [@children]="
            (app.isSlim() || app.isHorizontal()) && !app.isMobile() && root
              ? isActive(i)
                ? 'visible'
                : 'hidden'
              : isActive(i)
              ? 'visibleAnimated'
              : app.grouped === true && root
              ? 'visibleAnimated'
              : 'hiddenAnimated'
          "
        ></ul>
      </li>
    </ng-template>
  `,
  animations: [
    trigger("children", [
      state(
        "hiddenAnimated",
        style({
          height: "0px",
        })
      ),
      state(
        "visibleAnimated",
        style({
          height: "*",
        })
      ),
      state(
        "visible",
        style({
          height: "*",
          "z-index": 100,
        })
      ),
      state(
        "hidden",
        style({
          height: "0px",
          "z-index": "*",
        })
      ),
      transition(
        "visibleAnimated => hiddenAnimated",
        animate("400ms cubic-bezier(0.86, 0, 0.07, 1)")
      ),
      transition(
        "hiddenAnimated => visibleAnimated",
        animate("400ms cubic-bezier(0.86, 0, 0.07, 1)")
      ),
    ]),
  ],
})
export class SubMenuComponent {
  @Input() item: MenuItem;

  @Input() root: boolean;

  @Input() visible: boolean;

  _parentActive: boolean;

  _reset: boolean;

  activeIndex: number;

  constructor(public app: ViewsComponent, public appMenu: MenuComponent) { }

  itemClick(event: Event, item: MenuItem, index: number) {
    if (this.root) {
      this.app.menuHoverActive = !this.app.menuHoverActive;
    }
    // avoid processing disabled items
    if (item.disabled) {
      event.preventDefault();
      return true;
    }

    // activate current item and deactivate active sibling if any
    this.activeIndex = this.activeIndex === index ? null : index;

    // execute command
    if (item.command) {
      item.command({ originalEvent: event, item });
    }

    // prevent hash change
    if (item.items || (!item.url && !item.routerLink)) {
      setTimeout(() => {
        this.app.layoutMenuScrollerViewChild.moveBar();
      }, 450);
      event.preventDefault();
    }

    // hide menu
    if (!item.items) {
      if (this.app.isHorizontal() || this.app.isSlim()) {
        this.app.resetMenu = true;
      } else {
        this.app.resetMenu = false;
      }

      this.app.overlayMenuActive = false;
      this.app.staticMenuMobileActive = false;
      this.app.menuHoverActive = !this.app.menuHoverActive;
    }
  }

  onMouseEnter(index: number) {
    if (
      this.root &&
      this.app.menuHoverActive &&
      (this.app.isHorizontal() || this.app.isSlim()) &&
      !this.app.isMobile() &&
      !this.app.isTablet()
    ) {
      this.activeIndex = index;
    }
  }

  isActive(index: number): boolean {
    return this.activeIndex === index;
  }

  @Input() get reset(): boolean {
    return this._reset;
  }

  set reset(val: boolean) {
    this._reset = val;

    if (this._reset && (this.app.isHorizontal() || this.app.isSlim())) {
      this.activeIndex = null;
    }
  }

  @Input() get parentActive(): boolean {
    return this._parentActive;
  }

  set parentActive(val: boolean) {
    this._parentActive = val;

    if (!this._parentActive) {
      this.activeIndex = null;
    }
  }
}
