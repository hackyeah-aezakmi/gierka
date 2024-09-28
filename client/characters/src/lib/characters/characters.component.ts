import { Component, Input } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';

@Component({
  selector: 'lib-character',
  standalone: true,
  imports: [CommonModule, NgOptimizedImage],
  templateUrl: './characters.component.html',
  styles: `
      :host {
          position: relative;
          width: 10px;
          height: 0;
          display: flex;
          justify-content: center;
      }

      img {
          position: absolute;
          min-width: 150px;
          max-width: 150px;
          min-height: 150px;
          max-height: 150px;
          bottom: 0;
      }
  `,
})
export class CharactersComponent {
  @Input({required: true}) index!: number;
}
