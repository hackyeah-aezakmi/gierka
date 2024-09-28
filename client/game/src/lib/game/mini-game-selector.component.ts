import { Component, inject, OnDestroy, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StateService } from '@gierka/state';
import { MiniGame } from '../../../../state/src/lib/state/model/mini-game.enum';

@Component({
  selector: 'lib-mini-game-selector',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './mini-game-selector.component.html'
})
export class MiniGameSelectorComponent implements OnInit, OnDestroy {
  stateService = inject(StateService);

  selectMiniGame() {
    this.stateService.voteForMiniGame(MiniGame.QUIZ);
  }

  ngOnInit(): void {
    console.log(this.stateService.user);

    if (this.stateService.user.host) {
      setTimeout(() => {
        this.stateService.endVoting();
      }, 4000);
    }
  }

  ngOnDestroy() {
    this.stateService.resetVoteForMiniGame();
  }
}
