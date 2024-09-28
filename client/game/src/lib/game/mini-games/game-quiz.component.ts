import { AfterViewInit, Component, Input, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GameCommonComponent } from './game-common.component';
import { delay, of, tap, timeout } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

export type QuizConfig = {
  question: string,
  answers: [string, string, string, string],
  correctIndex: number,
}

@Component({
  selector: 'lib-game-quiz',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './game-quiz.component.html'
})
export class GameQuizComponent extends GameCommonComponent implements AfterViewInit {
  @Input() config!: QuizConfig;
  showAnswers = signal(false);
  public questionsIndexArray: Array<number> = [
    0,
    1,
    2,
    3
  ];

  override ngAfterViewInit() {
    super.ngAfterViewInit();

    of(true).pipe(
      takeUntilDestroyed(this.destroyRef),
      delay(1000),
      tap(() => {
        this.showAnswers.set(true);
      }),
      delay(1000),
    ).subscribe(() => {
      this.end.emit(0);
    });


  }

  handleAnswerClick(i: number) {
    if (i === this.config.correctIndex) {
      this.end.emit(Date.now() - this.startDate);
      this.buttonsDisabled.set(true);
    } else {
      this.buttonsDisabled.set(true);

      of(true).pipe(
        takeUntilDestroyed(this.destroyRef),
        delay(1000)
      ).subscribe(() => {
        this.buttonsDisabled.set(false);
      });
    }
  }
}
