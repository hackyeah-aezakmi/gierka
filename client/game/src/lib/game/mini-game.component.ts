import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MiniGame } from '../../../../state/src/lib/state/model/mini-game.enum';
import { GameQuizComponent } from './mini-games/game-quiz.component';
import { GameColorsComponent } from './mini-games/game-colors.component';

@Component({
  selector: 'lib-mini-game',
  standalone: true,
  imports: [CommonModule, GameQuizComponent, GameColorsComponent],
  templateUrl: './mini-game.component.html',
})
export class MiniGameComponent {
  protected readonly MiniGame = MiniGame;

  @Input({ required: true }) miniGame!: MiniGame;
  @Input() config!: any;

  handleEnd() {

  }
}
