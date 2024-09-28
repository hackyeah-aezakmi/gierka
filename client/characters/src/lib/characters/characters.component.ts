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

      .character {
          position: absolute;
          min-width: 150px;
          max-width: 150px;
          min-height: 150px;
          max-height: 150px;
          bottom: 0;

          &__label {
              padding: 3px 6px;
              border-radius: 3px;
              outline: 3px solid #e5e7eb;
              background-color: #C6EDCE;
              left: 50%;
              transform: translateX(-50%);
          }

          img {
              min-width: 150px;
              max-width: 150px;
              min-height: 150px;
              max-height: 150px;
          }
      }
  `,
})
export class CharactersComponent {
  @Input({required: true}) index!: number;
  @Input() nick?: string;
}
