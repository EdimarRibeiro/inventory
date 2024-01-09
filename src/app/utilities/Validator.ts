import { AbstractControl, ValidationErrors } from "@angular/forms";
import { validaCPF, validaCNPJ } from "app/utilities/validacao";

export class DateValidator {
  static LessThanToday(control: AbstractControl): ValidationErrors | null {
    let today: Date = new Date();
    today.setHours(0, 0, 0, 0);
    if (control.value && new Date(control.value) < today)
      return { LessThanToday: true };

    return null;
  }
}

export class CPFValidator {
  static IsValid(control: AbstractControl): ValidationErrors | null {
    let value: boolean = validaCPF(control.value);
    if (control.value == '')
      value = true;

    if (!value)
      return { IsValid: true };

    return null;
  }
}

export class CNPJValidator {
  static IsValid(control: AbstractControl): ValidationErrors | null {
    let value: boolean = validaCNPJ(control.value);
    if (!value)
      return { IsValid: true };

    return null;
  }
}
