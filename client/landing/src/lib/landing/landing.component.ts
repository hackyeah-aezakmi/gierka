import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import {QRCodeModule} from "angularx-qrcode";
import { Game, StateService, User } from '@gierka/state';
import { ButtonDirective } from 'primeng/button';
import { Ripple } from 'primeng/ripple';
import { CharactersComponent } from '@gierka/characters';

@Component({
  selector: 'lib-landing',
  standalone: true,
  imports: [CommonModule, QRCodeModule, ButtonDirective, Ripple, CharactersComponent],
  templateUrl: './landing.component.html',
  styles: `
      :host {
          background-image: url("/background.png");
          background-color: #adc68e;
          background-size: 600px 600px;
          display: block;
          background-repeat: no-repeat;
          background-position: center center;
      }

      .qr-code {
          box-shadow: 0 0 20px #60bc3f;
          display: block;
          border-radius: 14px;
          overflow: hidden;
      }
  `,
})
export class LandingComponent implements OnInit {
  stateService = inject(StateService);

  ngOnInit() {
    this.stateService.loadGame();
  }
}
