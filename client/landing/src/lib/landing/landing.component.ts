import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {QRCodeModule} from "angularx-qrcode";

@Component({
  selector: 'lib-landing',
  standalone: true,
  imports: [CommonModule, QRCodeModule],
  templateUrl: './landing.component.html',
})
export class LandingComponent {
  url: string = 'http://localhost:8080';
}
