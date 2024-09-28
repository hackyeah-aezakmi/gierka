import { AfterViewInit, Component, DestroyRef, EventEmitter, inject, Output, signal } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-game-common',
  standalone: true,
  template: '',
})
export class GameCommonComponent implements AfterViewInit {
  destroyRef = inject(DestroyRef);

  @Output() allowClick = new EventEmitter<boolean>();
  @Output() end = new EventEmitter<number>();
  buttonsDisabled = signal(false);
  startDate!: number;

  ngAfterViewInit() {
    this.startDate = Date.now();
  }
}
