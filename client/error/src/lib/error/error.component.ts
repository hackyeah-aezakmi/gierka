import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'lib-error',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './error.component.html',
})
export class ErrorComponent {
  @Input() control!: FormControl<any>;
}
