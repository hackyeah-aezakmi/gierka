import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-game-colors',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './game-colors.component.html',
})
export class GameColorsComponent {
  @Input() config!: any;
}
