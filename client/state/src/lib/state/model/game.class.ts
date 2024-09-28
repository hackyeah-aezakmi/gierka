import { MiniGame } from './mini-game.enum';

export class Game {
  id!: string;
  started: boolean = false;
  loading: boolean = false;
  miniGame?: MiniGame;
  config?: any;

  static new(id: string) {
    const game = new Game();

    game.id = id;
    return game;
  }
}
