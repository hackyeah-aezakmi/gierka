import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import {QRCodeModule} from "angularx-qrcode";
import { StateService } from '@gierka/state';
import { ButtonDirective } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
import { CharactersComponent } from '@gierka/characters';

@Component({
  selector: 'lib-landing',
  standalone: true,
  imports: [CommonModule, QRCodeModule, ButtonDirective, Ripple, CharactersComponent],
  templateUrl: './room.component.html',
  styles: `
      :host {
          background-image: url("/background.png");
          background-color: #adc68e;
          display: block;
          background-repeat: no-repeat;
          background-position: center center;
          background-size: contain;
      }
  `,
})
export class RoomComponent {
  stateService = inject(StateService);

  url: string = 'http://localhost:8080';
}
