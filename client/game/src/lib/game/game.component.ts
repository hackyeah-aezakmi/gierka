import { Component, DestroyRef, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntil } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { StateService } from '@gierka/state';
import { MiniGameSelectorComponent } from './mini-game-selector.component';
import { MiniGameComponent } from './mini-game.component';

@Component({
  selector: 'lib-game',
  standalone: true,
  imports: [CommonModule, MiniGameSelectorComponent, MiniGameComponent],
  templateUrl: './game.component.html',
})
export class GameComponent implements OnInit {
  stateService: StateService = inject(StateService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private destroyRef = inject(DestroyRef);

  public ngOnInit(): void {
    this.route.params.pipe(
      takeUntilDestroyed(this.destroyRef),
    ).subscribe((params) => {
      if (!(params['id'] && this.stateService.loadGame(params['id']))) {
        this.router.navigate(['/']);
      }
    });
  }

  private connectToWS() {
    // update state
  }

  handleMiniGameEnd($event: number) {
    this.stateService.user.points = $event;
  }
}
