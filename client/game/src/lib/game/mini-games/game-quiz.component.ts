import { Component, EventEmitter, Input, Output, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GameCommonComponent } from './game-common.component';

export type QuizConfig = {
  question: string,
  answers: [string, string, string, string],
  correctIndex: number,
}

@Component({
  selector: 'lib-game-quiz',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './game-quiz.component.html',
})
export class GameQuizComponent extends GameCommonComponent {
  @Input() config!: QuizConfig;
  public questionsIndexArray: Array<number> = [
    0,
    1,
    2,
    3,
  ];

  handleAnswerClick(i: number) {
    if (i === this.config.correctIndex) {
      this.end.emit();
    }
  }
}
