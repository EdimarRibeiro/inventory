import { Component, AfterViewInit, ElementRef, Renderer2, ViewChild, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { AuthenticationService } from '@auth/authentication.service';
import { ScrollPanel } from 'primeng/scrollpanel';

@Component({
    templateUrl: './views.component.html'
})
export class ViewsComponent implements AfterViewInit {

    layoutMode = 'static';
    darkMenu = true;
    profileMode = 'popup';
    rotateMenuButton: boolean;
    topbarMenuActive: boolean;
    overlayMenuActive: boolean;
    staticMenuDesktopInactive: boolean;
    staticMenuMobileActive: boolean;
    layoutMenuScroller: HTMLDivElement;
    menuClick: boolean;
    topbarItemClick: boolean;
    activeTopbarItem: any;
    resetMenu: boolean;
    menuHoverActive: boolean;
    grouped = false;

    @ViewChild('layoutMenuScroller', {static: true}) layoutMenuScrollerViewChild: ScrollPanel;

    constructor(public renderer: Renderer2, private router: Router, private authService: AuthenticationService) {
    }

    ngAfterViewInit() {
        setTimeout(() => {
            this.layoutMenuScrollerViewChild.moveBar();
        }, 100);
    }

    onLayoutClick() {
        if (!this.topbarItemClick) {
            this.activeTopbarItem = null;
            this.topbarMenuActive = false;
        }

        if (!this.menuClick) {
            if (this.isHorizontal() || this.isSlim()) {
                this.resetMenu = true;
            }

            if (this.overlayMenuActive || this.staticMenuMobileActive) {
                this.hideOverlayMenu();
            }

            this.menuHoverActive = false;
        }

        this.topbarItemClick = false;
        this.menuClick = false;
    }

    onMenuButtonClick(event) {
        this.menuClick = true;
        this.rotateMenuButton = !this.rotateMenuButton;
        this.topbarMenuActive = false;

        if (this.layoutMode === 'overlay' && !this.isMobile()) {
            this.overlayMenuActive = !this.overlayMenuActive;
        } else {
            if (this.isDesktop()) {
                this.staticMenuDesktopInactive = !this.staticMenuDesktopInactive;
            } else {
                this.staticMenuMobileActive = !this.staticMenuMobileActive;
            }
        }

        event.preventDefault();
    }

    onMenuClick($event) {
        this.menuClick = true;
        this.resetMenu = false;

        if (!this.isHorizontal()) {
            setTimeout(() => {
                this.layoutMenuScrollerViewChild.moveBar();
            }, 450);
        }
    }

    onTopbarMenuButtonClick(event) {
        this.topbarItemClick = true;
        this.topbarMenuActive = !this.topbarMenuActive;

        this.hideOverlayMenu();

        event.preventDefault();
    }

    onTopbarItemClick(event, item) {
        this.topbarItemClick = true;

        if (this.activeTopbarItem === item) {
            this.activeTopbarItem = null;
        } else {
            this.activeTopbarItem = item;
        }

        event.preventDefault();
    }

    onTopbarSubItemClick(event) {
        event.preventDefault();
    }

    hideOverlayMenu() {
        this.rotateMenuButton = false;
        this.overlayMenuActive = false;
        this.staticMenuMobileActive = false;
    }

    isTablet() {
        const width = window.innerWidth;
        return width < 1024 && width > 640;
    }

    isDesktop() {
        return window.innerWidth >= 897;

    }

    isMobile() {
        return window.innerWidth <= 640;
    }

    isOverlay() {
        return this.layoutMode === 'overlay';
    }

    isHorizontal() {
        return this.layoutMode === 'horizontal';
    }

    isSlim() {
        return this.layoutMode === 'slim';
    }

    isStatic() {
        return this.layoutMode === 'static';
    }

    onLogout(event) {
        this.authService.logout();
        this.router.navigateByUrl('/');
    }
}
