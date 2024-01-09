import { Injectable } from '@angular/core';
import { TokenStorageService } from '@auth/token-storage.service';

@Injectable()
export class HelperService {

    constructor(private tokenStorageService: TokenStorageService) {
    }

    public urlBase64Decode(str: string): string {
        let output = str.replace(/-/g, '+').replace(/_/g, '/');
        switch (output.length % 4) {
            case 0: {
                break;
            }
            case 2: {
                output += '==';
                break;
            }
            case 3: {
                output += '=';
                break;
            }
            default: {
                throw 'Illegal base64url string!';
            }
        }
        return this.b64DecodeUnicode(output);
    }

    private b64decode(str: string): string {
        let chars =
            'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
        let output: string = '';

        str = String(str).replace(/=+$/, '');

        if (str.length % 4 === 1) {
            throw new Error(
                '\'atob\' failed: The string to be decoded is not correctly encoded.'
            );
        }

        for (
            // initialize result and counters
            let bc: number = 0, bs: any, buffer: any, idx: number = 0;
            // get next character
            (buffer = str.charAt(idx++));
            // character found in table? initialize bit storage and add its ascii value;
            ~buffer &&
            (
                (bs = bc % 4 ? bs * 64 + buffer : buffer),
                    // and if not first of each 4 characters,
                    // convert the first 8 bits to one ascii character
                bc++ % 4
            )
                ? (output += String.fromCharCode(255 & (bs >> ((-2 * bc) & 6))))
                : 0
        ) {
            // try to find character in table (0-63, not found => -1)
            buffer = chars.indexOf(buffer);
        }
        return output;
    }

    private b64DecodeUnicode(str: any) {
        return decodeURIComponent(
            Array.prototype.map
                .call(this.b64decode(str), (c: any) => {
                    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
                })
                .join('')
        );
    }

    public decodeToken(token: string = this.tokenStorageService.getAccessTokenSync()): any {
        if ((token === null)||(token === undefined)||(token === 'undefined')) {
            return null;
        }

        let parts = token.split('.');

        if (parts.length !== 3) {
            throw new Error('The inspected token doesn\'t appear to be a JWT. Check to make sure it has three parts and see https://jwt.io for more.');
        }

        let decoded = this.urlBase64Decode(parts[1]);
        if (!decoded) {
            throw new Error('Cannot decode the token.');
        }

        return JSON.parse(decoded);
    }

    public getTokenExpirationDate(token: string = this.tokenStorageService.getAccessTokenSync()): Date | null {
        let decoded: any;
        decoded = this.decodeToken(token);

        if (!decoded.hasOwnProperty('exp')) {
            return null;
        }

        const date = new Date(0);
        date.setUTCSeconds(decoded.exp);
        return date;
    }

    public isTokenExpired(token: string = this.tokenStorageService.getAccessTokenSync(), offsetSeconds?: number): boolean {        
        token = String(token);
        if ((token == null) || (token == '') || (token == undefined) || (token == 'undefined')) {
            return true;
        }
        let date = this.getTokenExpirationDate(token);
        offsetSeconds = offsetSeconds || 0;

        if (date === null) {
            return true;
        }

        return !(date.valueOf() > new Date().valueOf() + offsetSeconds * 1000);
    }
}
