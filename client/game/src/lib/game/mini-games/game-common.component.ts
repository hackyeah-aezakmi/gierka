import { Component, EventEmitter, Output, signal } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-game-common',
  standalone: true,
  template: '',
})
export class GameCommonComponent {
  @Output() allowClick = new EventEmitter<boolean>();
  @Output() start = new EventEmitter<void>();
  @Output() end = new EventEmitter<void>();
  buttonsDisabled = signal(false);
}
