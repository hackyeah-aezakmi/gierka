export class Game {
  id!: string;

  static new(id: string) {
    const game = new Game();

    game.id = id;
    return game;
  }
}
