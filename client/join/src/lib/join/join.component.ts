import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { ChipsModule } from 'primeng/chips';
import { ErrorComponent } from '@gierka/error';
import { StateService } from '@gierka/state';
import { Router } from '@angular/router';

@Component({
  selector: 'lib-join',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, ChipsModule, ErrorComponent],
  templateUrl: './join.component.html',
  styles: `
      :host {
          height: 100vh;
          width: 100%;
          display: flex;
          align-items: center;
          justify-content: center;
          background-image: url("/background.png");
          background-color: #adc68e;
          background-repeat: no-repeat;
          background-position: center center;
          background-size: contain;
      }

      input {
          outline: 5px solid #9c5252;
      }
  `
})
export class JoinComponent {
  stateService: StateService = inject(StateService);
  router = inject(Router);

  nameFormControl = new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(10)]);

  join() {
    const value = this.nameFormControl.value;

    this.nameFormControl.setValue((this.nameFormControl.value ?? '').replace(/[^a-zA-Z-0-9]/g, ''));
    this.nameFormControl.markAsTouched();

    if (value !== this.nameFormControl.value) {
      return;
    }

    if (this.nameFormControl.valid) {
      this.stateService.addUser(this.stateService.createNewUser(this.nameFormControl.value!));
      this.router.navigate(['/']);
    }
  }
}
