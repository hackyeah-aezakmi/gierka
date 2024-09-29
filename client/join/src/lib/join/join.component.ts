import { Component, DestroyRef, inject, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { ChipsModule } from 'primeng/chips';
import { ErrorComponent } from '@gierka/error';
import { StateService } from '@gierka/state';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

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
export class JoinComponent implements OnInit {
  stateService: StateService = inject(StateService);
  router = inject(Router);
  route = inject(ActivatedRoute);
  destroyRef = inject(DestroyRef);
  joinDisabled = signal(false);

  nameFormControl = new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(10)]);

  ngOnInit(): void {
    this.route.params.pipe(
      takeUntilDestroyed(this.destroyRef),
    ).subscribe((params) => {
      if (!(params['id'] && this.stateService.loadGame(params['id']))) {
        this.router.navigate(['/']);
      }
    });
  }

  async join() {
    const value = this.nameFormControl.value;

    this.nameFormControl.setValue((this.nameFormControl.value ?? '').replace(/[^a-zA-Z-0-9]/g, ''));
    this.nameFormControl.markAsTouched();

    if (value !== this.nameFormControl.value) {
      return;
    }

    if (this.nameFormControl.valid) {
      try {
        this.joinDisabled.set(true);
        await this.stateService.addUser(this.stateService.createNewUser(this.nameFormControl.value!));
        await this.router.navigate(['/']);
      } catch (e) {
        console.error(e);
        this.joinDisabled.set(false);
      }
    }
  }
}
